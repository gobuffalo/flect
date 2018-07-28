package flect

import "strings"

var ToUpper = strings.ToUpper
var ToLower = strings.ToLower

func (i Ident) ToUpper() Ident {
	return New(ToUpper(i.Original))
}

func (i Ident) ToLower() Ident {
	return New(ToLower(i.Original))
}
