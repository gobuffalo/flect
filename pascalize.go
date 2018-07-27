package flect

import (
	"unicode"
)

func Pascalize(s string) string {
	return New(s).Pascalize()
}

func (i Ident) Pascalize() string {
	c := i.Camelize()
	if len(c) == 0 {
		return c
	}
	return string(unicode.ToUpper(rune(c[0]))) + c[1:]
}
