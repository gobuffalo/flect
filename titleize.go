package flect

import (
	"strings"
	"unicode"
)

func Titleize(s string) string {
	return New(s).Titleize()
}

func (i Ident) Titleize() string {
	var parts []string
	for _, part := range i.parts {
		var x string
		x = string(unicode.ToTitle(rune(part[0])))
		if len(part) > 1 {
			x += part[1:]
		}
		parts = append(parts, x)
	}
	return strings.Join(parts, " ")
}
