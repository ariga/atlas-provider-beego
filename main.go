package main

import (
	_ "embed"
	"fmt"
	"io"
	"os"
	"text/template"

	"ariga.io/atlas-go-sdk/tmplrun"
	"github.com/alecthomas/kong"
	"golang.org/x/tools/go/packages"
)

// LoadCmd is a command to load models
type LoadCmd struct {
	Path    string `help:"path to schema package" required:""`
	Dialect string `help:"dialect to use" enum:"mysql,sqlite,postgres" required:""`
	out     io.Writer
}

var (
	//go:embed loader.tmpl
	loader     string
	loaderTmpl = template.Must(template.New("loader").Parse(loader))
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
	cfg := &packages.Config{Mode: packages.NeedName | packages.NeedTypes | packages.NeedTypesInfo | packages.NeedModule | packages.NeedDeps}
	pkgs, err := packages.Load(cfg, c.Path)
	if err != nil {
		return err
	}
	p := Payload{
		Dialect: c.Dialect,
		Imports: []string{pkgs[0].PkgPath},
	}
	s, err := tmplrun.New("beegoschema", loaderTmpl).Run(p)
	if c.out == nil {
		c.out = os.Stdout
	}
	_, err = fmt.Fprintln(c.out, s)
	return err
}

// Payload is the data passed to the loader template.
type Payload struct {
	Dialect string
	Imports []string
}
