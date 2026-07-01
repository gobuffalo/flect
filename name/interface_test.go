package name

import (
	"fmt"
	"testing"
)

type car struct{}

func Test_Interface(t *testing.T) {
	table := []struct {
		in  any
		out string
		err bool
	}{
		{"foo", "foo", false},
		{car{}, "car", false},
		{&car{}, "car", false},
		{[]car{}, "cars", false},
		{false, "bool", false},
	}

	for _, tt := range table {
		t.Run(fmt.Sprint(tt.in), func(st *testing.T) {
			r := newRequire(st)
			n, err := Interface(tt.in)
			if tt.err {
				r.Error(err)
				return
			}
			r.NoError(err)
			r.Equal(tt.out, n.String())
		})
	}
}
