package beegoschema

import (
	"database/sql"
	"fmt"
	"strings"

	"ariga.io/atlas-go-sdk/recordriver"

	"github.com/beego/beego/v2/client/orm"
)

// New returns a new Loader.
func New(dialect string) *Loader {
	return &Loader{dialect: dialect}
}

// Loader is a Loader for beego schema.
type Loader struct {
	dialect string
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
	for _, stmt := range ss.Statements {
		if _, err := fmt.Fprint(&buf, stmt); err != nil {
			return "", err
		}
		buf.WriteString("\n")
	}
	return buf.String(), nil
}
