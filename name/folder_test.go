package name

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Folder(t *testing.T) {
	table := []tt{
		{"", ""},
		{"foo_bar", "foo/bar"},
		{"admin/widget", "admin/widget"},
		{"admin/widgets", "admin/widgets"},
		{"widget", "widget"},
		{"widgets", "widgets"},
		{"User", "user"},
		{"U$er", "u/er"},
	}

	for _, tt := range table {
		t.Run(tt.act, func(st *testing.T) {
			r := require.New(st)
			r.Equal(osify(tt.exp), Folder(tt.act))
			r.Equal(osify(tt.exp), Folder(tt.exp))
			r.Equal(osify(tt.exp)+".a.b", Folder(tt.act, ".a", ".b"))
		})
	}
}
