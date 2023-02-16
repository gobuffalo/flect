package flect

import (
	"strings"
	"sync"
)

var singularMoot = &sync.RWMutex{}

// Singularize returns a singular version of the string
//
//	users = user
//	data = datum
//	people = person
func Singularize(s string) string {
	return New(s).Singularize().String()
}

// SingularizeWithSize will singular a string taking a number number into account.
//
//	SingularizeWithSize("user", 1) = user
//	SingularizeWithSize("user", 2) = users
func SingularizeWithSize(s string, i int) string {
	return PluralizeWithSize(s, i)
}

// Singularize returns a singular version of the string
//
//	users = user
//	data = datum
//	people = person
func (i Ident) Singularize() Ident {
	s := i.LastPart()
	if len(s) == 0 {
		return i
	}

	// Special rules for Acronyms
	// As per https://editorsmanual.com/articles/plurals-of-abbreviations/:
	if baseAcronyms[strings.ToUpper(s)] {
		// It's already singular - it's one of our acronyms
		return i
	} else if strings.HasSuffix(strings.ToLower(s), "es") {
		// String ends with `es` - check to see if what's left is an acronym
		a := s[:len(s)-2]
		if baseAcronyms[strings.ToUpper(a)] {
			return New(a)
		}
	} else if strings.HasSuffix(strings.ToLower(s), "s") {
		// String ends with `s` - check to see if what's left is an acronym
		a := s[:len(s)-1]
		if baseAcronyms[strings.ToUpper(a)] {
			return New(a)
		}
	}

	singularMoot.RLock()
	defer singularMoot.RUnlock()

	ls := strings.ToLower(s)
	if p, ok := pluralToSingle[ls]; ok {
		if s == Capitalize(s) {
			p = Capitalize(p)
		}
		return i.ReplaceSuffix(s, p)
	}

	if _, ok := singleToPlural[ls]; ok {
		return i
	}

	for _, r := range singularRules {
		if strings.HasSuffix(ls, r.suffix) {
			return i.ReplaceSuffix(s, r.fn(s))
		}
	}

	if strings.HasSuffix(s, "s") {
		return i.ReplaceSuffix("s", "")
	}

	return i
}
