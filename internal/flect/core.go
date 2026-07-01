package flect

import (
	"strings"
	"unicode"
)

func isSpace(c rune) bool {
	switch c {
	case '_', ' ', ':', '-', '/':
		return true
	}
	return unicode.IsSpace(c)
}

func xappend(a []string, ss ...string) []string {
	for _, s := range ss {
		s = strings.TrimFunc(s, isSpace)
		up := strings.ToUpper(s)
		if _, ok := BaseAcronyms[up]; ok {
			s = up
		}
		if s != "" {
			a = append(a, s)
		}
	}
	return a
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
