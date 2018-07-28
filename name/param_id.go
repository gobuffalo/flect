package name

import "strings"

// ParamID returns the string as parameter with _id added
//	user = user_id
//	UserID = user_id
//	admin/widgets = admin_widgets_id
func ParamID(s string) string {
	return New(s).ParamID()
}

// ParamID returns the string as parameter with _id added
//	user = user_id
//	UserID = user_id
//	admin/widgets = admin_widgets_id
func (i Ident) ParamID() string {
	s := i.Underscore()
	s = strings.ToLower(s)
	if strings.HasSuffix(s, "_id") {
		return s
	}
	return s + "_id"
}
