package flect

import "strings"

// File creates a suitable file name
// admin/widget = admin/widget
// foo_bar = foo_bar
// U$ser = u_ser
func File(s string, exts ...string) string {
	return New(s).File(exts...)
}

// File creates a suitable file name
// admin/widget = admin/widget
// foo_bar = foo_bar
// U$ser = u_ser
func (i Ident) File(exts ...string) string {
	var parts []string

	for _, part := range strings.Split(i.original, "/") {
		parts = append(parts, Underscore(part))
	}
	return strings.Join(parts, "/") + strings.Join(exts, "")
}
