package flect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Pluralize(t *testing.T) {
	for _, tt := range singlePluralAssertions {
		if tt.doPluralizeTest {
			t.Run(tt.singular, func(st *testing.T) {
				r := require.New(st)
				r.Equal(tt.plural, Pluralize(tt.singular), "pluralize %s", tt.singular)
				r.Equal(tt.plural, Pluralize(tt.plural), "pluralize %s", tt.plural)
			})
		}
	}
}

func Test_PluralizeWithSize(t *testing.T) {
	for _, tt := range singlePluralAssertions {
		t.Run(tt.singular, func(st *testing.T) {
			r := require.New(st)
			if tt.doSingularizeTest {
				r.Equal(tt.singular, PluralizeWithSize(tt.singular, -1), "pluralize %d %s", -1, tt.singular)
				r.Equal(tt.singular, PluralizeWithSize(tt.plural, -1), "pluralize %d %s", -1, tt.plural)
				r.Equal(tt.singular, PluralizeWithSize(tt.singular, 1), "pluralize %d %s", 1, tt.singular)
				r.Equal(tt.singular, PluralizeWithSize(tt.plural, 1), "pluralize %d %s", 1, tt.plural)
			}
			if tt.doPluralizeTest {
				r.Equal(tt.plural, PluralizeWithSize(tt.singular, -2), "pluralize %d %s", -2, tt.singular)
				r.Equal(tt.plural, PluralizeWithSize(tt.plural, -2), "pluralize %d %s", -2, tt.plural)
				r.Equal(tt.plural, PluralizeWithSize(tt.singular, 0), "pluralize %d %s", 0, tt.singular)
				r.Equal(tt.plural, PluralizeWithSize(tt.plural, 0), "pluralize %d %s", 0, tt.plural)
				r.Equal(tt.plural, PluralizeWithSize(tt.singular, 2), "pluralize %d %s", 2, tt.singular)
				r.Equal(tt.plural, PluralizeWithSize(tt.plural, 2), "pluralize %d %s", 2, tt.plural)
			}
		})
	}
}

func BenchmarkPluralize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range singlePluralAssertions {
			if tt.doPluralizeTest {
				Pluralize(tt.singular)
				Pluralize(tt.plural)
			}
		}
	}
}
