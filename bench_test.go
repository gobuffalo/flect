package flect

import "testing"

// benchInputs exercises the key code paths: acronym fast-path, camelCase
// splitting, underscore-separated words, multi-segment paths, and plain words.
var benchInputs = []string{
	"widget",
	"widget_id",
	"WidgetID",
	"Widget_ID",
	"HTMLParser",
	"JSONResponse",
	"admin/widget",
	"bob dylan",
	"Nice to see you!",
	"ThisIsALongCamelCaseString",
	"JWTName",
	"foo_bar_baz",
}

func BenchmarkCamelize(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, s := range benchInputs {
			Camelize(s)
		}
	}
}

func BenchmarkPascalize(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, s := range benchInputs {
			Pascalize(s)
		}
	}
}

func BenchmarkDasherize(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, s := range benchInputs {
			Dasherize(s)
		}
	}
}

func BenchmarkTitleize(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, s := range benchInputs {
			Titleize(s)
		}
	}
}

func BenchmarkHumanize(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, s := range benchInputs {
			Humanize(s)
		}
	}
}

func BenchmarkCapitalize(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		for _, s := range benchInputs {
			Capitalize(s)
		}
	}
}

func BenchmarkOrdinalize(b *testing.B) {
	b.ReportAllocs()
	inputs := []string{"1", "2", "3", "11", "12", "13", "21", "42", "100", "101", "111", "1001"}
	for n := 0; n < b.N; n++ {
		for _, s := range inputs {
			Ordinalize(s)
		}
	}
}
