package flect

import (
	"strings"
	"sync"
)

var singularMoot = &sync.RWMutex{}

// Singularize returns a singular version of the string
//	users = user
//	data = datum
//	people = person
func Singularize(s string) string {
	return New(s).Singularize()
}

// Singularize returns a singular version of the string
//	users = user
//	data = datum
//	people = person
func (i Ident) Singularize() string {
	s := i.Original
	if len(s) == 0 {
		return ""
	}

	singularMoot.RLock()
	defer singularMoot.RUnlock()
	ls := strings.ToLower(s)
	if p, ok := pluralToSingle[ls]; ok {
		return p
	}
	if _, ok := singleToPlural[ls]; ok {
		return s
	}
	for _, r := range singularRules {
		if strings.HasSuffix(ls, r.suffix) {
			return r.fn(s)
		}
	}

	return s
}
