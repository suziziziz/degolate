package decorators

import (
	"fmt"
	"reflect"
	"runtime"
	"time"

	"github.com/suziziziz/degolate"
)

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
