package flect

import "github.com/gobuffalo/flect/internal/core"

// Underscore a string
//
//	bob dylan --> bob_dylan
//	Nice to see you! --> nice_to_see_you
//	widgetID --> widget_id
func Underscore(s string) string {
	return core.Underscore(s)
}
