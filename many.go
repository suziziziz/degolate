package degolate

import (
	"reflect"
	"slices"
)

// Just a simple way to apply many `decorators` to `function`
// See `./examples/basic/main.go`
func Many[F any](decorators func(fn interface{}), function F) (fn F) {
	decorators(function)

	fnV := reflect.ValueOf(function)

	idx := slices.IndexFunc(degolate.degolated, func(d Degolated) bool {
		return d.fn == fnV
	})
	degolated := &degolate.degolated[idx]

	return reflect.MakeFunc(fnV.Type(), func(args []reflect.Value) (results []reflect.Value) {
		for _, d := range degolated.decorators {
			d.Call([]reflect.Value{})
		}

		return degolated.fn.Call(args)
	}).Interface().(F)
}
