package flect

import (
	"testing"

	th "github.com/gobuffalo/flect/internal/testhelpers"
)

type requireHelper = th.RequireHelper

func newRequire(t testing.TB) *requireHelper {
	return th.NewRequire(t)
}
