package flect

func Name(s string) string {
	return New(s).Name()
}

func (i Ident) Name() string {
	s := i.Singularize()
	s = Pascalize(s)
	return s
}
