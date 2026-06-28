package flect

import (
	"testing"
)

func Test_Capitalize(t *testing.T) {
	table := []tt{
		{"", ""},
		{"foo", "Foo"},
		{"bob dylan", "Bob dylan"},
		{"WidgetID", "WidgetID"},
		{"widget_id", "Widget_id"},
		{"widget_ID", "Widget_ID"},
		{"widget ID", "Widget ID"},
		{"гофер", "Гофер"}, // it's "gopher" in Ukrainian
	}

	for _, tt := range table {
		t.Run(tt.act, func(st *testing.T) {
			r := newRequire(st)
			r.Equal(tt.exp, Capitalize(tt.act))
			r.Equal(tt.exp, Capitalize(tt.exp))
		})
	}
}
