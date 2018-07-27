package flect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Camelize(t *testing.T) {
	table := []tt{
		{"", ""},
		{"bob dylan", "bobDylan"},
		{"widgetID", "widgetID"},
		{"widget_ID", "widgetID"},
		{"Widget_ID", "widgetID"},
		{"Nice to see you!", "niceToSeeYou"},
		{"*hello*", "hello"},
		{"i've read a book! have you?", "iveReadABookHaveYou"},
		{"This is `code` ok", "thisIsCodeOK"},
	}

	for _, tt := range table {
		t.Run(tt.act, func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, Camelize(tt.act))
			r.Equal(tt.exp, Camelize(tt.exp))
		})
	}
}
