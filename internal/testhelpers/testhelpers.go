package testhelpers

import (
	"fmt"
	"reflect"
	"testing"
)

type RequireHelper struct {
	T testing.TB
}

func NewRequire(t testing.TB) *RequireHelper {
	t.Helper()
	return &RequireHelper{T: t}
}

func (r *RequireHelper) Equal(exp, got any, msgAndArgs ...any) {
	r.T.Helper()
	if !reflect.DeepEqual(exp, got) {
		msg := FormatTestMsg(msgAndArgs)
		if msg != "" {
			r.T.Fatalf("%s\nexpected: %v\ngot:      %v", msg, exp, got)
		} else {
			r.T.Fatalf("expected: %v\ngot:      %v", exp, got)
		}
	}
}

func (r *RequireHelper) NoError(err error) {
	r.T.Helper()
	if err != nil {
		r.T.Fatalf("unexpected error: %v", err)
	}
}

func (r *RequireHelper) Error(err error) {
	r.T.Helper()
	if err == nil {
		r.T.Fatal("expected error, got nil")
	}
}

func (r *RequireHelper) True(cond bool) {
	r.T.Helper()
	if !cond {
		r.T.Fatal("expected true, got false")
	}
}

func FormatTestMsg(msgAndArgs []any) string {
	if len(msgAndArgs) == 0 {
		return ""
	}
	if msg, ok := msgAndArgs[0].(string); ok {
		return fmt.Sprintf(msg, msgAndArgs[1:]...)
	}
	return fmt.Sprintf("%v", msgAndArgs[0])
}
