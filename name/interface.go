package name

import (
	"errors"
	"reflect"
)

func Interface(x interface{}) (Ident, error) {
	switch t := x.(type) {
	case string:
		return New(t), nil
	default:
		to := reflect.TypeOf(x)
		if len(to.Name()) > 0 {
			return New(to.Name()), nil
		}
		k := to.Kind()
		switch k {
		case reflect.Slice, reflect.Array:
			e := to.Elem()
			n := New(e.Name())
			return New(n.Pluralize().String()), nil
		}
	}
	return New(""), errors.New("could not convert to Ident")
}
