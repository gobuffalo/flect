package name

import core "github.com/gobuffalo/flect/internal/flect"

// Ident represents the string and it's parts
type Ident struct {
	core.Ident
}

// New creates a new Ident from the string
func New(s string) Ident {
	return Ident{core.New(s)}
}
