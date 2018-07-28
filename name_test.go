package flect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Name(t *testing.T) {
	table := []tt{
		{"", ""},
		{"bob dylan", "BobDylan"},
		{"widgetID", "WidgetID"},
		{"widget_ID", "WidgetID"},
		{"Widget_ID", "WidgetID"},
		{"Widget_Id", "WidgetID"},
		{"Widget_id", "WidgetID"},
		{"Nice to see you!", "NiceToSeeYou"},
		{"*hello*", "Hello"},
		{"i've read a book! have you?", "IveReadABookHaveYou"},
		{"This is `code` ok", "ThisIsCodeOK"},
		{"foo_bar", "FooBar"},
		{"admin/widget", "AdminWidget"},
		{"admin/widgets", "AdminWidget"},
		{"widget", "Widget"},
		{"widgets", "Widget"},
		{"status", "Status"},
		{"Statuses", "Status"},
		{"statuses", "Status"},
		{"People", "Person"},
		{"people", "Person"},
	}

	for _, tt := range table {
		t.Run(tt.act, func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, Name(tt.act))
			r.Equal(tt.exp, Name(tt.exp))
		})
	}
}

func Test_GroupName(t *testing.T) {
	table := []tt{
		{"", ""},
		{"Person", "People"},
		{"foo_bar", "FooBars"},
		{"admin/widget", "AdminWidgets"},
		{"widget", "Widgets"},
		{"widgets", "Widgets"},
		{"greatPerson", "GreatPeople"},
		{"great/person", "GreatPeople"},
		{"status", "Statuses"},
		{"Status", "Statuses"},
		{"Statuses", "Statuses"},
		{"statuses", "Statuses"},
	}

	for _, tt := range table {
		t.Run(tt.act, func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, GroupName(tt.act))
			r.Equal(tt.exp, GroupName(tt.exp))
		})
	}
}
