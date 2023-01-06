package flect

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Singularize(t *testing.T) {
	for _, tt := range singlePluralAssertions {
		if tt.doSingularizeTest {
			t.Run(tt.plural, func(st *testing.T) {
				r := require.New(st)
				r.Equal(tt.singular, Singularize(tt.plural), "singularize %s", tt.plural)
				r.Equal(tt.singular, Singularize(tt.singular), "singularize %s", tt.singular)
			})
		}
	}
}

func Test_SingularizeWithSize(t *testing.T) {
	for _, tt := range singlePluralAssertions {
		t.Run(tt.plural, func(st *testing.T) {
			r := require.New(st)
			if tt.doSingularizeTest {
				r.Equal(tt.singular, SingularizeWithSize(tt.plural, -1), "singularize %d %s", -1, tt.plural)
				r.Equal(tt.singular, SingularizeWithSize(tt.singular, -1), "singularize %d %s", -1, tt.singular)
				r.Equal(tt.singular, SingularizeWithSize(tt.plural, 1), "singularize %d %s", 1, tt.plural)
				r.Equal(tt.singular, SingularizeWithSize(tt.singular, 1), "singularize %d %s", 1, tt.singular)
			}
			if tt.doPluralizeTest {
				r.Equal(tt.plural, SingularizeWithSize(tt.plural, -2), "singularize %d %s", -2, tt.plural)
				r.Equal(tt.plural, SingularizeWithSize(tt.singular, -2), "singularize %d %s", -2, tt.singular)
				r.Equal(tt.plural, SingularizeWithSize(tt.plural, 0), "singularize %d %s", 0, tt.plural)
				r.Equal(tt.plural, SingularizeWithSize(tt.singular, 0), "singularize %d %s", 0, tt.singular)
				r.Equal(tt.plural, SingularizeWithSize(tt.plural, 2), "singularize %d %s", 2, tt.plural)
				r.Equal(tt.plural, SingularizeWithSize(tt.singular, 2), "singularize %d %s", 2, tt.singular)
			}
		})
	}
}

func BenchmarkSingularize(b *testing.B) {
	for n := 0; n < b.N; n++ {
		for _, tt := range singlePluralAssertions {
			if tt.doSingularizeTest {
				Singularize(tt.singular)
				Singularize(tt.plural)
			}
		}
	}
}
