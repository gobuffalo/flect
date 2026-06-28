package core

import "strconv"

// Ordinalize converts a number to an ordinal version
//
//	42 = 42nd
//	45 = 45th
//	1 = 1st
func Ordinalize(s string) string {
	return New(s).Ordinalize().String()
}

// Ordinalize converts a number to an ordinal version
//
//	42 = 42nd
//	45 = 45th
//	1 = 1st
func (i Ident) Ordinalize() Ident {
	number, err := strconv.Atoi(i.Original)
	if err != nil {
		return i
	}
	var suffix string
	switch abs(number) % 100 {
	case 11, 12, 13:
		suffix = "th"
	default:
		switch abs(number) % 10 {
		case 1:
			suffix = "st"
		case 2:
			suffix = "nd"
		case 3:
			suffix = "rd"
		default:
			suffix = "th"
		}
	}
	return Ident{Original: strconv.Itoa(number) + suffix}
}
