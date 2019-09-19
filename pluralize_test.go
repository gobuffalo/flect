package flect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Pluralize(t *testing.T) {
	for _, tt := range singlePluralAssertions {
		t.Run(tt.act, func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, Pluralize(tt.act))
			r.Equal(tt.exp, Pluralize(tt.exp))
		})
	}
}

func Test_PluralizeWithSize(t *testing.T) {
	for _, tt := range singlePluralAssertions {
		t.Run(tt.act, func(st *testing.T) {
			r := require.New(st)
			r.Equal(tt.exp, PluralizeWithSize(tt.act, -2))
			r.Equal(tt.exp, PluralizeWithSize(tt.exp, -2))
			r.Equal(tt.act, PluralizeWithSize(tt.act, -1))
			r.Equal(tt.act, PluralizeWithSize(tt.exp, -1))
			r.Equal(tt.exp, PluralizeWithSize(tt.act, 0))
			r.Equal(tt.exp, PluralizeWithSize(tt.exp, 0))
			r.Equal(tt.act, PluralizeWithSize(tt.act, 1))
			r.Equal(tt.act, PluralizeWithSize(tt.exp, 1))
			r.Equal(tt.exp, PluralizeWithSize(tt.act, 2))
			r.Equal(tt.exp, PluralizeWithSize(tt.exp, 2))
		})
	}
}
