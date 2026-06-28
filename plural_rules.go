package flect

import "github.com/gobuffalo/flect/internal/core"

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
