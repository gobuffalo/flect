package flect

import "github.com/gobuffalo/flect/internal/core"

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
