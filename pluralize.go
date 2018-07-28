package flect

import (
	"strings"
	"sync"
)

var pluralMoot = &sync.RWMutex{}

// Pluralize returns a plural version of the string
//	user = users
//	person = people
//	datum = data
func Pluralize(s string) string {
	return New(s).Pluralize()
}

// Pluralize returns a plural version of the string
//	user = users
//	person = people
//	datum = data
func (i Ident) Pluralize() string {
	s := i.Original
	if len(s) == 0 {
		return ""
	}

	pluralMoot.RLock()
	defer pluralMoot.RUnlock()

	ls := strings.ToLower(s)
	if _, ok := pluralToSingle[ls]; ok {
		return s
	}
	if p, ok := singleToPlural[ls]; ok {
		return p
	}
	for _, r := range pluralRules {
		if strings.HasSuffix(ls, r.suffix) {
			return r.fn(s)
		}
	}

	if strings.HasSuffix(ls, "s") {
		return s
	}

	return s + "s"
}
