package name

import "github.com/gobuffalo/flect"

type Ident struct {
	flect.Ident
}

func New(s string) Ident {
	return Ident{flect.New(s)}
}
