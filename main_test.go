package main

import (
	"bytes"
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
			require.Contains(t, buf.String(), "CREATE TABLE")
			require.Contains(t, buf.String(), "hotdog_type")
			require.Contains(t, buf.String(), "stand")
			require.Contains(t, buf.String(), "hotdog_stock")
		})
	}
}
