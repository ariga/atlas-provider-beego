package main

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLoad(t *testing.T) {
	for _, dialect := range []string{"mysql", "sqlite3", "postgres"} {
		t.Run(dialect, func(t *testing.T) {
			var buf bytes.Buffer
			cmd := &LoadCmd{
				Path:    "./internal/testdata/models",
				Dialect: dialect,
				out:     &buf,
			}
			err := cmd.Run()
			require.NoError(t, err)
			sql := buf.String()
			file, err := os.ReadFile("./internal/testdata/models/" + dialect + ".sql")
			require.NoError(t, err)
			expected := string(file)
			require.Equal(t, expected, sql, "generated SQL does not match expected SQL for dialect %s", dialect)
		})
	}
}
