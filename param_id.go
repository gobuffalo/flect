package flect

import "strings"

func ParamID(s string) string {
	return New(s).ParamID()
}

func (i Ident) ParamID() string {
	s := i.Underscore()
	s = strings.ToLower(s)
	if strings.HasSuffix(s, "_id") {
		return s
	}
	return s + "_id"
}
