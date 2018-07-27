package flect

import "unicode"

func Capitalize(s string) string {
	return New(s).Capitalize()
}

func (i Ident) Capitalize() string {
	var x string
	if len(i.parts) == 0 {
		return ""
	}
	x = string(unicode.ToTitle(rune(i.original[0])))
	if len(i.original) > 1 {
		x += i.original[1:]
	}
	return x
}
