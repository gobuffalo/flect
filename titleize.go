package flect

import "github.com/gobuffalo/flect/internal/core"

// Titleize will capitalize the start of each part
//
//	"Nice to see you!" = "Nice To See You!"
//	"i've read a book! have you?" = "I've Read A Book! Have You?"
//	"This is `code` ok" = "This Is `code` OK"
func Titleize(s string) string {
	return core.Titleize(s)
}
