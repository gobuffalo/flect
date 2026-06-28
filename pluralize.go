package flect

import "github.com/gobuffalo/flect/internal/core"

// Pluralize returns a plural version of the string
//
//	user = users
//	person = people
//	datum = data
func Pluralize(s string) string {
	return core.Pluralize(s)
}

// PluralizeWithSize will pluralize a string taking a number into account.
//
//	PluralizeWithSize("user", 1) = user
//	PluralizeWithSize("user", 2) = users
func PluralizeWithSize(s string, i int) string {
	return core.PluralizeWithSize(s, i)
}
