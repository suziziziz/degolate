package decorators

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/suziziziz/degolate"
)

// A decorator is declared as a function.
//
// Using `degolate.It` you can apply a decorator to a function. To be simpler
// and avoid code replication, just put `degolate.It` directly in the return of
// a decorator, as you can see below.
func Logging[F any](fn F) F {
	return degolate.It(fn, func() {
		fmt.Printf(
			"[%s]: %s %T\n",
			time.Now().Format(time.DateTime),
			runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(),
			fn,
		)
	})
}
