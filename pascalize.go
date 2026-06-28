package flect

import "github.com/gobuffalo/flect/internal/core"

// Pascalize returns a string with each segment capitalized
//
//	user = User
//	bob dylan = BobDylan
//	widget_id = WidgetID
func Pascalize(s string) string {
	return core.Pascalize(s)
}
