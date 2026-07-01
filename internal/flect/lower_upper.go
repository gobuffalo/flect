package flect

import "strings"

// ToUpper is a convenience wrapper for strings.ToUpper
func (i Ident) ToUpper() Ident {
	return Ident{Original: strings.ToUpper(i.Original)}
}

// ToLower is a convenience wrapper for strings.ToLower
func (i Ident) ToLower() Ident {
	return Ident{Original: strings.ToLower(i.Original)}
}
