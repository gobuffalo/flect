package name

import (
	"runtime"
	"testing"
)

func Test_Ident_OsPath(t *testing.T) {
	table := map[string]string{
		"foo/bar/baz":   "foo/bar/baz",
		"foo\\bar\\baz": "foo/bar/baz",
	}

	if runtime.GOOS == "windows" {
		table = ident_OsPath_Windows_Table()
	}

	for in, out := range table {
		t.Run(in, func(st *testing.T) {
			r := newRequire(st)
			r.Equal(out, OsPath(in))
		})
	}
}

func ident_OsPath_Windows_Table() map[string]string {
	return map[string]string{
		"foo/bar/baz":   "foo\\bar\\baz",
		"foo\\bar\\baz": "foo\\bar\\baz",
	}
}
