package flect

import (
	"unicode"
)

// Pascalize returns a string with each segment capitalized
//	user = User
//	bob dylan = BobDylan
//	widget_id = WidgetID
func Pascalize(s string) string {
	return New(s).Pascalize()
}

// Pascalize returns a string with each segment capitalized
//	user = User
//	bob dylan = BobDylan
//	widget_id = WidgetID
func (i Ident) Pascalize() string {
	c := i.Camelize()
	if len(c) == 0 {
		return c
	}
	return string(unicode.ToUpper(rune(c[0]))) + c[1:]
}
