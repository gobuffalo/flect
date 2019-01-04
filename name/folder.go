package name

import (
	"path/filepath"
	"strings"

	"github.com/gobuffalo/flect"
)

// Folder creates a suitable folder name
//	admin/widget = admin/widget
//	foo_bar = foo_bar
//	U$ser = u_ser
func Folder(s string, exts ...string) string {
	return New(s).Folder(exts...).String()
}

// Folder creates a suitable folder name
//	admin/widget = admin/widget
//	foo_bar = foo/bar
//	U$ser = u/ser
func (i Ident) Folder(exts ...string) Ident {
	var parts []string

	for _, part := range strings.Split(i.Original, "/") {
		part = flect.Underscore(part)
		part = strings.Replace(part, "_", string(filepath.Separator), -1)
		parts = append(parts, part)
	}
	return New(strings.Join(parts, "/") + strings.Join(exts, ""))
}
