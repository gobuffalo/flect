package flect

// source for grammar rules: https://www.grammarly.com/blog/plural-nouns/
import (
	"strings"
)

func Singularize(s string) string {
	return New(s).Singularize()
}

func (i Ident) Singularize() string {
	if len(i.parts) == 0 {
		return ""
	}
	li := len(i.parts) - 1
	last := i.parts[li]
	if _, ok := singleToPlural[strings.ToLower(last)]; ok {
		return i.original
	}
	parts := i.parts[:li]
	if p, ok := pluralToSingle[strings.ToLower(last)]; ok {
		parts = append(parts, p)
		return strings.Join(parts, " ")
	}

	for _, r := range singularRules {
		if strings.HasSuffix(last, r.suffix) {
			parts = append(parts, r.fn(last))
			return strings.Join(parts, " ")
		}
	}

	if len(last) > 2 {
		if isVowel(rune(last[len(last)-2])) {
			parts = append(parts, last)
			return strings.Join(parts, " ")
		}
	}
	//To make regular nouns plural, add ‑s to the end.
	parts = append(parts, strings.TrimSuffix(last, "s"))
	return strings.Join(parts, " ")
}

//  If the singular noun ends in ‑s, -ss, -sh, -ch, -x, or -z, add ‑es to the end to make it plural.
func esSingular(s string) string {
	return strings.TrimSuffix(s, "es")
}

//  If the noun ends with ‑f or ‑fe, the f is often changed to ‑ve before adding the -s to form the plural version.
func vesSingular(s string) string {
	s = strings.TrimSuffix(s, "ves")
	return s + "f"
}

// If a singular noun ends in ‑y and the letter before the -y is a consonant, change the ending to ‑ies to make the noun plural.
func iesSingular(s string) string {
	s = strings.TrimSuffix(s, "ies")
	return s + "y"
}

// If the singular noun ends in ‑us, the plural ending is frequently ‑i.
func iSingular(s string) string {
	s = strings.TrimSuffix(s, "i")
	return s + "us"
}

// If the singular noun ends in ‑on, the plural ending is ‑a.
func aSingular(s string) string {
	s = strings.TrimSuffix(s, "a")
	return s + "on"
}

func exSingular(s string) string {
	s = strings.TrimSuffix(s, "ex")
	return s + "ix"
}

func ssSingular(s string) string {
	if len(s) > 3 {
		if isVowel(rune(s[len(s)-3])) {
			return s
		}
	}
	return strings.TrimSuffix(s, "ss")
}

var singularRules = []rule{
	{"ix", noop},
	{"us", noop},
	{"y", noop},
	{"on", noop},
	{"f", noop},
	{"i", iSingular},
	{"ss", ssSingular},
	{"ies", iesSingular},
	{"ves", vesSingular},
	{"ex", exSingular},
	{"a", aSingular},
	{"es", esSingular},
}
