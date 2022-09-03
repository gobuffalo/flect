package name

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Folder(t *testing.T) {
	table := []tt{
		{"", ""},
		{"admin/widget", "admin/widget"},
		{"admin/widgets", "admin/widgets"},
		{"widget", "widget"},
		{"widgets", "widgets"},
		{"User", "user"},
		{"U$er", "u_er"},
		{"adminuser", "adminuser"},
		{"Adminuser", "adminuser"},
		{"AdminUser", "admin_user"},
		{"adminUser", "admin_user"},
		{"admin-user", "admin_user"},
		{"admin_user", "admin_user"},
	}

	for _, tt := range table {
		t.Run(tt.act, func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, Folder(tt.act))
			r.Equal(tt.exp, Folder(tt.exp))
			r.Equal(tt.exp+".a.b", Folder(tt.act, ".a", ".b"))
		})
	}
}
