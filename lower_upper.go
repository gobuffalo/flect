package flect

import "strings"

// ToUpper is a convience wrapper for strings.ToUpper
var ToUpper = strings.ToUpper

// ToLower is a convience wrapper for strings.ToLower
var ToLower = strings.ToLower

// ToUpper is a convience wrapper for strings.ToUpper
func (i Ident) ToUpper() Ident {
	return New(ToUpper(i.Original))
}

// ToLower is a convience wrapper for strings.ToLower
func (i Ident) ToLower() Ident {
	return New(ToLower(i.Original))
}
