package name

import (
	"strings"

	"github.com/gobuffalo/flect"
)

// Proper pascalizes and singularizes the string
//	person = Person
//	foo_bar = FooBar
//	admin/widgets = AdminWidget
func Proper(s string) string {
	return New(s).Proper()
}

// Proper pascalizes and singularizes the string
//	person = Person
//	foo_bar = FooBar
//	admin/widgets = AdminWidget
func (i Ident) Proper() string {
	s := i.Singularize()
	s = flect.Pascalize(s)
	return s
}

// Group pascalizes and pluralizes the string
//	person = People
//	foo_bar = FooBars
//	admin/widget = AdminWidgets
func Group(s string) string {
	return New(s).Group()
}

// Group pascalizes and pluralizes the string
//	person = People
//	foo_bar = FooBars
//	admin/widget = AdminWidgets
func (i Ident) Group() string {
	var parts []string
	if len(i.Original) == 0 {
		return i.Original
	}
	last := i.Parts[len(i.Parts)-1]
	for _, part := range i.Parts[:len(i.Parts)-1] {
		parts = append(parts, flect.Pascalize(part))
	}
	last = flect.Pascalize(flect.Pluralize(last))
	parts = append(parts, last)
	return strings.Join(parts, "")
}
