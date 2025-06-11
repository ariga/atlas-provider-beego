package main

import (
	_ "embed"
	"fmt"
	"go/token"
	"go/types"
	"io"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"ariga.io/atlas-go-sdk/tmplrun"
	"github.com/alecthomas/kong"
	"golang.org/x/tools/go/callgraph"
	"golang.org/x/tools/go/callgraph/static"
	"golang.org/x/tools/go/packages"
	"golang.org/x/tools/go/ssa"
	"golang.org/x/tools/go/ssa/ssautil"
)

// LoadCmd is a command to load models
type LoadCmd struct {
	Path    string `help:"path to schema package" required:""`
	Dialect string `help:"dialect to use" enum:"mysql,sqlite3,postgres" required:""`
	out     io.Writer
}

var (
	//go:embed loader.tmpl
	loader     string
	loaderTmpl = template.Must(template.New("loader").Parse(loader))
	beegoPkg   = "github.com/beego/beego/v2/client/orm"
)

func main() {
	var cli struct {
		Load LoadCmd `cmd:""`
	}
	ctx := kong.Parse(&cli)
	if err := ctx.Run(); err != nil {
		fmt.Fprintln(os.Stderr, err) // nolint: errcheck
		os.Exit(1)
	}
}

func (c *LoadCmd) Run() error {
	cfg := &packages.Config{Mode: packages.NeedName |
		packages.NeedTypes |
		packages.NeedTypesInfo |
		packages.NeedModule |
		packages.NeedDeps |
		packages.NeedImports |
		packages.NeedSyntax,
	}
	pkgs, err := packages.Load(cfg, c.Path)
	if err != nil {
		return err
	}
	models := gatherModels(pkgs)
	p := Payload{
		Dialect: c.Dialect,
		Models:  models,
	}
	s, err := tmplrun.New("beegoschema", loaderTmpl).Run(p)
	if err != nil {
		return err
	}
	if c.out == nil {
		c.out = os.Stdout
	}
	_, err = fmt.Fprintln(c.out, s)
	return err
}

// Payload is the data passed to the loader template.
type Payload struct {
	Dialect string
	Models  []model
}

func (p Payload) Imports() []string {
	imports := make(map[string]struct{})
	for _, m := range p.Models {
		imports[m.ImportPath] = struct{}{}
	}
	var result []string
	for k := range imports {
		result = append(result, k)
	}
	return result
}

type model struct {
	ImportPath string
	PkgName    string
	Name       string
	Pos        string
}

func (m model) String() string {
	return fmt.Sprintf("%s.%s", m.PkgName, m.Name)
}

func gatherModels(pkgs []*packages.Package) []model {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error getting current working directory: %v\n", err)
		return nil
	}
	models := gatherRegisterModels(pkgs, dir)
	slices.SortFunc(models, func(i, j model) int {
		return strings.Compare(i.Name, j.Name)
	})
	return models
}

// gartherRegisterModels gathers models registered with register functions.
func gatherRegisterModels(pkgs []*packages.Package, wd string) []model {
	// Find all functions callable from the beego package that register models.
	prog, _ := ssautil.AllPackages(pkgs, 0)
	prog.Build()
	g := static.CallGraph(prog)
	var models []model
	err := callgraph.GraphVisitEdges(g, func(edge *callgraph.Edge) error {
		caller, callee := edge.Caller.Func, edge.Callee.Func
		if caller == nil || callee == nil {
			return nil
		}
		callerFPkg, calleeFPkg := caller.Pkg, callee.Pkg
		if callerFPkg == nil || calleeFPkg == nil {
			return nil
		}
		// Only consider edges where the callee is a function in the beego package.
		// and the caller is a function in the current package.
		if !isBeegoRegisterCall(callee) || !isFunctionInPackages(caller, pkgs) {
			return nil
		}
		common := edge.Site.Common()
		for _, arg := range common.Args {
			sliceInstr, ok := arg.(*ssa.Slice)
			if !ok {
				continue
			}
			alloc, ok := sliceInstr.X.(*ssa.Alloc)
			if !ok {
				continue
			}
			positions := extractStructsFromSlice(alloc, prog.Fset)
			for modelName, pos := range positions {
				// Resolve path to a relative path if possible.
				relPath, err := filepath.Rel(wd, pos.Filename)
				if err != nil {
					fmt.Fprintf(os.Stderr, "error resolving relative path: %v\n", err)
					relPath = pos.Filename // fallback to absolute path
				}
				models = append(models, model{
					ImportPath: callerFPkg.Pkg.Path(),
					PkgName:    callerFPkg.Pkg.Name(),
					Name:       modelName,
					Pos:        fmt.Sprintf("%s:%d:%d", relPath, pos.Line, pos.Column),
				})
			}
		}
		return nil
	})
	if err != nil {
		fmt.Fprintf(os.Stderr, "error traversing call graph: %v\n", err)
		return nil
	}
	return models
}

func isBeegoRegisterCall(f *ssa.Function) bool {
	if f == nil || f.Pkg == nil || f.Pkg.Pkg == nil {
		return false
	}
	return f.Pkg.Pkg.Path() == beegoPkg &&
		(f.Name() == "RegisterModel" ||
			f.Name() == "RegisterModelWithPrefix" ||
			f.Name() == "RegisterModelWithSuffix")
}

func isFunctionInPackages(f *ssa.Function, pkgs []*packages.Package) bool {
	if f == nil || f.Pkg == nil || f.Pkg.Pkg == nil {
		return false
	}
	for _, pkg := range pkgs {
		if f.Pkg.Pkg.Path() == pkg.PkgPath {
			return true
		}
	}
	return false
}

func extractStructsFromSlice(slice ssa.Value, fset *token.FileSet) map[string]*token.Position {
	var positions = make(map[string]*token.Position)
	if s, ok := slice.(*ssa.Slice); ok {
		slice = s.X
	}
	if u, ok := slice.(*ssa.UnOp); ok {
		slice = u.X
	}
	// Now search for stores into this slice
	parentFn := slice.Parent()
	for _, block := range parentFn.Blocks {
		for _, instr := range block.Instrs {
			store, ok := instr.(*ssa.Store)
			if !ok {
				continue
			}
			indexAddr, ok := store.Addr.(*ssa.IndexAddr)
			if !ok {
				continue
			}
			if indexAddr.X != slice {
				continue
			}
			val := store.Val
			if iface, ok := val.(*ssa.MakeInterface); ok {
				val = iface.X
			}
			alloc, ok := val.(*ssa.Alloc)
			if !ok {
				continue
			}
			ptrType, ok := alloc.Type().(*types.Pointer)
			if !ok {
				continue
			}
			named, ok := ptrType.Elem().(*types.Named)
			if !ok {
				continue
			}
			pos := fset.Position(named.Obj().Pos())
			positions[named.Obj().Name()] = &pos
		}
	}
	return positions
}
