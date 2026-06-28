package flect

import "github.com/gobuffalo/flect/internal/core"

// Singularize returns a singular version of the string
//
//	users = user
//	data = datum
//	people = person
func Singularize(s string) string {
	return core.Singularize(s)
}

// SingularizeWithSize will singularize a string taking a number into account.
//
//	SingularizeWithSize("user", 1) = user
//	SingularizeWithSize("user", 2) = users
func SingularizeWithSize(s string, i int) string {
	return core.SingularizeWithSize(s, i)
}
