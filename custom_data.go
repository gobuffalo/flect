package flect

import (
	"io"

	"github.com/gobuffalo/flect/internal/core"
)

// CustomDataParser are functions that parse data like acronyms or
// plurals in the shape of a io.Reader it receives.
type CustomDataParser = core.CustomDataParser

// LoadAcronyms loads acronyms from io.Reader param
func LoadAcronyms(r io.Reader) error {
	return core.LoadAcronyms(r)
}

// LoadInflections loads rules from io.Reader param
func LoadInflections(r io.Reader) error {
	return core.LoadInflections(r)
}
