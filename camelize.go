package flect

import "github.com/gobuffalo/flect/internal/core"

// Camelize returns a camelize version of a string
//
//	bob dylan = bobDylan
//	widget_id = widgetID
//	WidgetID = widgetID
func Camelize(s string) string {
	return core.Camelize(s)
}
