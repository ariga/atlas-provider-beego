package beegoschema

import (
	"database/sql"
	"os"

	_ "ariga.io/atlas-go-sdk/recordriver"
	"github.com/beego/beego/v2/client/orm"
)

// New returns a new Loader.
func New(dialect string) *Loader {
	return &Loader{dialect: dialect}
}

// Loader is a Loader for gorm schema.
type Loader struct {
	dialect string
}

func (l *Loader) Load() (string, error) {
	db, err := sql.Open("recordriver", "beego")
	if err != nil {
		return "", err
	}
	if err := orm.AddAliasWthDB("default", l.dialect, db); err != nil {
		return "", err
	}
	before := os.Args
	defer func() {
		os.Args = before
	}()
	// https://github.com/beego/beedoc/blob/master/en-US/mvc/model/cmd.md#print-sql-statements
	os.Args = []string{"beego", "orm", "sqlall"}
	orm.RunCommand()
	return "", nil
}
