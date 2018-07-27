package flect

import (
	"strings"
	"unicode"
)

func Underscore(s string) string {
	return New(s).Underscore()
}

func (i Ident) Underscore() string {
	var out []string
	for _, part := range i.parts {
		var x string
		for _, c := range part {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				x += string(c)
			}
		}
		if x != "" {
			out = append(out, x)
		}
	}
	return strings.ToLower(strings.Join(out, "_"))
}
