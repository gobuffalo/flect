package flect

import (
	"encoding"
	"strings"
	"unicode"
	"unicode/utf8"
)

// Ident represents the string and it's parts
type Ident struct {
	Original string
	Parts    []string
}

// String implements fmt.Stringer and returns the original string
func (i Ident) String() string {
	return i.Original
}

// New creates a new Ident from the string
func New(s string) Ident {
	i := Ident{
		Original: s,
		Parts:    toParts(s),
	}

	return i
}

func toParts(s string) []string {
	parts := []string{}
	s = strings.TrimSpace(s)
	if len(s) == 0 {
		return parts
	}
	up := strings.ToUpper(s)
	if _, ok := BaseAcronyms[up]; ok {
		return []string{up}
	}
	var prev rune
	var x strings.Builder
	x.Grow(len(s))
	for _, c := range s {
		if !utf8.ValidRune(c) {
			continue
		}

		if isSpace(c) {
			parts = xappend(parts, x.String())
			x.Reset()
			x.WriteRune(c)
			prev = c
			continue
		}

		if unicode.IsUpper(c) && !unicode.IsUpper(prev) {
			parts = xappend(parts, x.String())
			x.Reset()
			x.WriteRune(c)
			prev = c
			continue
		}
		if _, isAcronym := BaseAcronyms[x.String()]; unicode.IsUpper(c) && isAcronym {
			parts = xappend(parts, x.String())
			x.Reset()
			x.WriteRune(c)
			prev = c
			continue
		}
		if unicode.IsLetter(c) || unicode.IsDigit(c) || unicode.IsPunct(c) || c == '`' {
			prev = c
			x.WriteRune(c)
			continue
		}

		parts = xappend(parts, x.String())
		x.Reset()
		prev = c
	}
	parts = xappend(parts, x.String())

	return parts
}

var _ encoding.TextUnmarshaler = &Ident{}
var _ encoding.TextMarshaler = &Ident{}

// LastPart returns the last part/word of the original string
func (i *Ident) LastPart() string {
	if len(i.Parts) == 0 {
		return ""
	}
	return i.Parts[len(i.Parts)-1]
}

// ReplaceSuffix creates a new Ident with the original suffix replaced by new
func (i Ident) ReplaceSuffix(orig, new string) Ident {
	return New(strings.TrimSuffix(i.Original, orig) + new)
}

// UnmarshalText unmarshalls byte array into the Ident
func (i *Ident) UnmarshalText(data []byte) error {
	(*i) = New(string(data))
	return nil
}

// MarshalText marshals Ident into byte array
func (i Ident) MarshalText() ([]byte, error) {
	return []byte(i.Original), nil
}
