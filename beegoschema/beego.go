package beegoschema

import (
	"database/sql"
	"fmt"
	"io"
	"maps"
	"reflect"
	"slices"
	"strings"

	"ariga.io/atlas-go-sdk/recordriver"

	"github.com/beego/beego/v2/client/orm"
)

// New returns a new Loader.
func New(dialect string, opts ...Option) *Loader {
	l := &Loader{
		dialect: dialect,
	}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

type (
	// Loader is a Loader for beego schema.
	Loader struct {
		dialect  string
		modelPos map[any]string
	}
	// Option configures the Loader.
	Option func(*Loader)
)

// WithModelPosition sets the model position in the output.
// The position is used to generate the `-- atlas:pos` directive in the output.
func WithModelPosition(pos map[any]string) Option {
	return func(l *Loader) {
		l.modelPos = pos
	}
}

func (l *Loader) Load() (string, error) {
	db, err := sql.Open("recordriver", "beego")
	if err != nil {
		return "", err
	}
	defer db.Close()
	if err := orm.AddAliasWthDB("default", l.dialect, db); err != nil {
		return "", err
	}
	if err := orm.RunSyncdb("default", false, false); err != nil {
		return "", err
	}
	ss, ok := recordriver.Session("beego")
	if !ok {
		return "", fmt.Errorf("could not retrieve recordriver session")
	}
	var buf strings.Builder
	if err = l.directives(&buf); err != nil {
		return "", err
	}
	for _, stmt := range ss.Statements {
		if _, err := fmt.Fprint(&buf, stmt); err != nil {
			return "", err
		}
		buf.WriteString("\n")
	}
	return buf.String(), nil
}

func (l *Loader) directives(w io.Writer) error {
	if len(l.modelPos) > 0 {
		pos := map[string]string{}
		for m, p := range l.modelPos {
			pos[fmt.Sprintf("%s[type=table]", GetTableName(m))] = p
		}
		for _, r := range slices.Sorted(maps.Keys(pos)) {
			if _, err := fmt.Fprintln(w, "-- atlas:pos", r, pos[r]); err != nil {
				return err
			}
		}
		// Add another new line to separate the file directives from the statements.
		if _, err := fmt.Fprintln(w); err != nil {
			return err
		}
	}
	return nil
}

// https://github.com/beego/beego/blob/master/client/orm/internal/models/models_utils.go#L74
// Cloned from beego source code since it is not exported.
func GetTableName(model any) string {
	val := reflect.ValueOf(model)
	if fn := val.MethodByName("TableName"); fn.IsValid() {
		vals := fn.Call([]reflect.Value{})
		// has return and the first val is string
		if len(vals) > 0 && vals[0].Kind() == reflect.String {
			return vals[0].String()
		}
	}
	return SnakeString(reflect.Indirect(val).Type().Name())
}

// https://github.com/beego/beego/blob/master/client/orm/internal/models/models_utils.go#L280
// Cloned from beego source code since it is not exported.
func SnakeString(s string) string {
	data := make([]byte, 0, len(s)*2)
	j := false
	num := len(s)
	for i := 0; i < num; i++ {
		d := s[i]
		if i > 0 && d >= 'A' && d <= 'Z' && j {
			data = append(data, '_')
		}
		if d != '_' {
			j = true
		}
		data = append(data, d)
	}
	return strings.ToLower(string(data))
}
