package flect

import "github.com/gobuffalo/flect/internal/core"

// Capitalize will cap the first letter of string
//
//	user = User
//	bob dylan = Bob dylan
//	widget_id = Widget_id
func Capitalize(s string) string {
	return core.Capitalize(s)
}
