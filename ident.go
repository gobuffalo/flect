package flect

import "github.com/gobuffalo/flect/internal/core"

// Ident represents the string and it's parts
type Ident = core.Ident

// New creates a new Ident from the string
func New(s string) Ident {
	return core.New(s)
}
