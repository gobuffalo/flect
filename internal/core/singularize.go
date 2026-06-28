package core

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

// SingularizeWithSize will singularize a string taking a number into account.
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

	singularMoot.RLock()

	// check if the Original has an explicit entry in the map
	if p, ok := pluralToSingle[i.Original]; ok {
		singularMoot.RUnlock()
		return i.ReplaceSuffix(i.Original, p)
	}
	if _, ok := singleToPlural[i.Original]; ok {
		singularMoot.RUnlock()
		return i
	}

	ls := strings.ToLower(s)
	if p, ok := pluralToSingle[ls]; ok {
		if s == Capitalize(s) {
			p = Capitalize(p)
		}
		singularMoot.RUnlock()
		return i.ReplaceSuffix(s, p)
	}

	if _, ok := singleToPlural[ls]; ok {
		singularMoot.RUnlock()
		return i
	}

	rules := singularRules
	singularMoot.RUnlock()

	for _, r := range rules {
		if strings.HasSuffix(s, r.suffix) {
			return i.ReplaceSuffix(s, r.fn(s))
		}
	}

	if strings.HasSuffix(s, "s") {
		return i.ReplaceSuffix("s", "")
	}

	return i
}
