package degolate

import (
	"reflect"
	"slices"
)

type Degolated struct {
	fn         reflect.Value
	decorators []reflect.Value
}

type Degolate struct {
	degolated []Degolated
}

var degolate = &Degolate{}

func It(function any, decorator any) (idx int) {
	fn, dec := reflect.ValueOf(function), reflect.ValueOf(decorator)

	if fn.Kind() != reflect.Func {
		panic("param `fn` must be a function")
	}
	if dec.Kind() != reflect.Func {
		panic("param `dec` must be a function")
	}

	idx = slices.IndexFunc(degolate.degolated, func(d Degolated) bool {
		return d.fn == fn
	})

	if idx != -1 {
		degolate.degolated[idx].decorators = append(degolate.degolated[idx].decorators, dec)

		return idx
	}

	degolate.degolated = append(degolate.degolated, Degolated{fn, []reflect.Value{dec}})

	return len(degolate.degolated) - 1
}

func D[F any](function F) F {
	idx := slices.IndexFunc(degolate.degolated, func(d Degolated) bool {
		return d.fn == reflect.ValueOf(function)
	})

	return degolate.degolated[idx].fn.Interface().(F) // TODO: MISSING DECORATOR CALL
}
