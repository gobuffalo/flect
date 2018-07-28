package flect

// Tableize returns an underscore, pluralized string
// User = users
// Person = persons
// Admin/Widget = admin_widgets
func Tableize(s string) string {
	return New(s).Tableize()
}

// Tableize returns an underscore, pluralized string
// User = users
// Person = persons
// Admin/Widget = admin_widgets
func (i Ident) Tableize() string {
	return Underscore(i.Pluralize())
}
