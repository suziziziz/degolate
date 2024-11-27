package degolate

import "reflect"

type Degolated struct {
	fn         reflect.Value
	decorators []reflect.Value
}

type Degolate struct {
	degolated []Degolated
}

var degolate = &Degolate{}
