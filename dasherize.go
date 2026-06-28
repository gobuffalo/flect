package flect

import "github.com/gobuffalo/flect/internal/core"

// Dasherize returns an alphanumeric, lowercased, dashed string
//
//	Donald E. Knuth = donald-e-knuth
//	Test with + sign = test-with-sign
//	admin/WidgetID = admin-widget-id
func Dasherize(s string) string {
	return core.Dasherize(s)
}
