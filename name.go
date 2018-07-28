package flect

import "strings"

// GroupName pascalizes and singularizes the string
// person = Person
// foo_bar = FooBar
// admin/widgets = AdminWidget
func Name(s string) string {
	return New(s).Name()
}

// GroupName pascalizes and singularizes the string
// person = Person
// foo_bar = FooBar
// admin/widgets = AdminWidget
func (i Ident) Name() string {
	s := i.Singularize()
	s = Pascalize(s)
	return s
}

// GroupName pascalizes and pluralizes the string
// person = People
// foo_bar = FooBars
// admin/widget = AdminWidgets
func GroupName(s string) string {
	return New(s).GroupName()
}

// GroupName pascalizes and pluralizes the string
// person = People
// foo_bar = FooBars
// admin/widget = AdminWidgets
func (i Ident) GroupName() string {
	var parts []string
	if len(i.original) == 0 {
		return i.original
	}
	last := i.parts[len(i.parts)-1]
	for _, part := range i.parts[:len(i.parts)-1] {
		parts = append(parts, Pascalize(part))
	}
	last = Pascalize(Pluralize(last))
	parts = append(parts, last)
	return strings.Join(parts, "")
}
