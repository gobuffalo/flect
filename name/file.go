package name

import (
	"strings"

	core "github.com/gobuffalo/flect/internal/flect"
)

// File creates a suitable file name
//
//	admin/widget = admin/widget
//	foo_bar = foo_bar
//	U$ser = u_ser
func File(s string, exts ...string) string {
	return New(s).File(exts...).String()
}

// File creates a suitable file name
//
//	admin/widget = admin/widget
//	foo_bar = foo_bar
//	U$ser = u_ser
func (i Ident) File(exts ...string) Ident {
	var parts []string

	for _, part := range strings.Split(i.Original, "/") {
		parts = append(parts, core.Underscore(part))
	}
	return New(strings.Join(parts, "/") + strings.Join(exts, ""))
}
