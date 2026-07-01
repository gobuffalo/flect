/*
Package flect is a new inflection engine to replace [https://github.com/markbates/inflect](https://github.com/markbates/inflect) designed to be more modular, more readable, and easier to fix issues on than the original.
*/
package flect

import (
	"io"

	"github.com/gobuffalo/flect/internal/core"
)

// Ident represents the string and it's parts
type Ident = core.Ident

// CustomDataParser are functions that parse data like acronyms or
// plurals in the shape of a io.Reader it receives.
type CustomDataParser = core.CustomDataParser

// baseAcronyms and dictionary are the same references as in internal/core,
// kept here so that package-level tests in this package can access them
// without importing the internal package directly.
var baseAcronyms = core.BaseAcronyms
var dictionary = core.Dictionary

// New creates a new Ident from the string
func New(s string) Ident {
	return core.New(s)
}

// Camelize returns a camelize version of a string
//
//	bob dylan = bobDylan
//	widget_id = widgetID
//	WidgetID = widgetID
func Camelize(s string) string {
	return core.Camelize(s)
}

// Capitalize will cap the first letter of string
//
//	user = User
//	bob dylan = Bob dylan
//	widget_id = Widget_id
func Capitalize(s string) string {
	return core.Capitalize(s)
}

// Dasherize returns an alphanumeric, lowercased, dashed string
//
//	Donald E. Knuth = donald-e-knuth
//	Test with + sign = test-with-sign
//	admin/WidgetID = admin-widget-id
func Dasherize(s string) string {
	return core.Dasherize(s)
}

// Humanize returns first letter of sentence capitalized.
// Common acronyms are capitalized as well.
// Other capital letters in string are left as provided.
//
//	employee_salary = Employee salary
//	employee_id = employee ID
//	employee_mobile_number = Employee mobile number
//	first_Name = First Name
//	firstName = First Name
func Humanize(s string) string {
	return core.Humanize(s)
}

// Ordinalize converts a number to an ordinal version
//
//	42 = 42nd
//	45 = 45th
//	1 = 1st
func Ordinalize(s string) string {
	return core.Ordinalize(s)
}

// Pascalize returns a string with each segment capitalized
//
//	user = User
//	bob dylan = BobDylan
//	widget_id = WidgetID
func Pascalize(s string) string {
	return core.Pascalize(s)
}

// Pluralize returns a plural version of the string
//
//	user = users
//	person = people
//	datum = data
func Pluralize(s string) string {
	return core.Pluralize(s)
}

// PluralizeWithSize will pluralize a string taking a number into account.
//
//	PluralizeWithSize("user", 1) = user
//	PluralizeWithSize("user", 2) = users
func PluralizeWithSize(s string, i int) string {
	return core.PluralizeWithSize(s, i)
}

// AddPlural adds a rule that will replace the given suffix with the replacement suffix.
// The name is confusing. This function will be deprecated in the next release.
func AddPlural(suffix string, repl string) {
	core.AddPlural(suffix, repl)
}

// InsertPluralRule inserts a rule that will replace the given suffix with
// the repl(acement) at the begining of the list of the pluralize rules.
func InsertPluralRule(suffix, repl string) {
	core.InsertPluralRule(suffix, repl)
}

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

// AddSingular adds a rule that will replace the given suffix with the replacement suffix.
// The name is confusing. This function will be deprecated in the next release.
func AddSingular(ext string, repl string) {
	core.AddSingular(ext, repl)
}

// InsertSingularRule inserts a rule that will replace the given suffix with
// the repl(acement) at the beginning of the list of the singularize rules.
func InsertSingularRule(suffix, repl string) {
	core.InsertSingularRule(suffix, repl)
}

// Titleize will capitalize the start of each part
//
//	"Nice to see you!" = "Nice To See You!"
//	"i've read a book! have you?" = "I've Read A Book! Have You?"
//	"This is `code` ok" = "This Is `code` OK"
func Titleize(s string) string {
	return core.Titleize(s)
}

// Underscore a string
//
//	bob dylan --> bob_dylan
//	Nice to see you! --> nice_to_see_you
//	widgetID --> widget_id
func Underscore(s string) string {
	return core.Underscore(s)
}

// LoadAcronyms loads acronyms from io.Reader param
func LoadAcronyms(r io.Reader) error {
	return core.LoadAcronyms(r)
}

// LoadInflections loads rules from io.Reader param
func LoadInflections(r io.Reader) error {
	return core.LoadInflections(r)
}
