package degolate

import (
	"reflect"
)

// Use this to register a decorator.
//
// See `./examples/basic/decorators/logging.go`
func It(decorator any) D {
	dec := funcValidate(decorator)

	return func(function any) any {
		fn := funcValidate(function)

		return reflect.MakeFunc(fn.Type(), func(args []reflect.Value) (results []reflect.Value) {
			dec.Call([]reflect.Value{})

			return fn.Call(args)
		}).Interface()
	}
}

func funcValidate(value any) reflect.Value {
	val := reflect.ValueOf(value)

	if val.Kind() != reflect.Func {
		panic("param `value` must be a function")
	}

	return val
}
