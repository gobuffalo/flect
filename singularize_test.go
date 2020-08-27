package flect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Singularize(t *testing.T) {
	for _, tt := range pluralSingularAssertions {
		t.Run(tt.exp, func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, Singularize(tt.act))
			r.Equal(tt.exp, Singularize(tt.exp))
		})
	}
}

func Test_SingularizeWithSize(t *testing.T) {
	for _, tt := range pluralSingularAssertions {
		t.Run(tt.exp, func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.act, SingularizeWithSize(tt.act, -2))
			r.Equal(tt.act, SingularizeWithSize(tt.exp, -2))
			r.Equal(tt.exp, SingularizeWithSize(tt.act, -1))
			r.Equal(tt.exp, SingularizeWithSize(tt.exp, -1))
			r.Equal(tt.act, SingularizeWithSize(tt.act, 0))
			r.Equal(tt.act, SingularizeWithSize(tt.exp, 0))
			r.Equal(tt.exp, SingularizeWithSize(tt.act, 1))
			r.Equal(tt.exp, SingularizeWithSize(tt.exp, 1))
			r.Equal(tt.act, SingularizeWithSize(tt.act, 2))
			r.Equal(tt.act, SingularizeWithSize(tt.exp, 2))
		})
	}
}
