package flect

import "github.com/gobuffalo/flect/internal/core"

// Ordinalize converts a number to an ordinal version
//
//	42 = 42nd
//	45 = 45th
//	1 = 1st
func Ordinalize(s string) string {
	return core.Ordinalize(s)
}
