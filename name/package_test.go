package name

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Package(t *testing.T) {
	gp := goPath()
	fmt.Println("### gp ->", gp)
	table := []tt{
		{filepath.Join(gp, "src", "admin/widget"), "admin/widget"},
		{filepath.Join(gp, "admin/widget"), "admin/widget"},
		{filepath.Join(gp, "admin\\widget"), "admin/widget"},
		{"admin/widget", "admin/widget"},
		{"admin\\widget", "admin/widget"},
	}

	for _, tt := range table {
		t.Run(tt.act, func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, Package(tt.act))
			r.Equal(tt.exp, Package(tt.exp))
		})
	}
}
