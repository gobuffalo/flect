package flect

import (
	"strings"
	"unicode"
)

func Camelize(s string) string {
	return New(s).Camelize()
}

func (i Ident) Camelize() string {
	var out []string
	for i, part := range i.parts {
		var x string
		var capped bool
		for _, c := range part {
			if unicode.IsLetter(c) || unicode.IsDigit(c) {
				if i == 0 {
					x += string(unicode.ToLower(c))
					continue
				}
				if !capped {
					capped = true
					x += string(unicode.ToUpper(c))
					continue
				}
				x += string(c)
			}
		}
		if x != "" {
			out = append(out, x)
		}
	}
	return strings.Join(out, "")
}
