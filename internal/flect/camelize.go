package flect

import (
	"strings"
	"unicode"
)

// Camelize returns a camelize version of a string
//
//	bob dylan = bobDylan
//	widget_id = widgetID
//	WidgetID = widgetID
func Camelize(s string) string {
	return New(s).Camelize().String()
}

// Camelize returns a camelize version of a string
//
//	bob dylan = bobDylan
//	widget_id = widgetID
//	WidgetID = widgetID
func (i Ident) Camelize() Ident {
	var out []string
	for i, part := range i.Parts {
		var x strings.Builder
		var capped bool
		for _, c := range part {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				if i == 0 {
					x.WriteRune(unicode.ToLower(c))
					continue
				}
				if !capped {
					capped = true
					x.WriteRune(unicode.ToUpper(c))
					continue
				}
				x.WriteRune(c)
			}
		}
		if x.Len() > 0 {
			out = append(out, x.String())
		}
	}
	return Ident{Original: strings.Join(out, "")}
}
