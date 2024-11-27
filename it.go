package degolate

import (
	"reflect"
	"slices"
)

// Use this to register a decorator.
// See `./examples/basic/main.go`
func It[F any](function F, decorator any) F {
	fn, dec := paramsValidate(function, decorator)

	idx := slices.IndexFunc(degolate.degolated, func(d Degolated) bool { return d.fn == fn })

	if idx == -1 {
		degolate.degolated = append(degolate.degolated, Degolated{fn, []reflect.Value{}})
		idx = len(degolate.degolated) - 1
	}

	degolated := &degolate.degolated[idx]

	degolated.decorators = append(degolated.decorators, dec)

	return reflect.MakeFunc(fn.Type(), func(args []reflect.Value) (results []reflect.Value) {
		for _, d := range degolated.decorators {
			d.Call([]reflect.Value{})
		}

		return degolated.fn.Call(args)
	}).Interface().(F)
}

func paramsValidate[F any](function F, decorator any) (reflect.Value, reflect.Value) {
	fn, dec := reflect.ValueOf(function), reflect.ValueOf(decorator)

	if fn.Kind() != reflect.Func {
		panic("param `fn` must be a function")
	}

	if dec.Kind() != reflect.Func {
		panic("param `dec` must be a function")
	}

	return fn, dec
}
