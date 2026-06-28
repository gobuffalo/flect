package flect

import "github.com/gobuffalo/flect/internal/core"

// baseAcronyms and dictionary are the same references as in internal/core,
// kept here so that package-level tests in this package can access them
// without importing the internal package directly.
var baseAcronyms = core.BaseAcronyms
var dictionary = core.Dictionary
