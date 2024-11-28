package decorators

import (
	"fmt"
	"time"

	"github.com/suziziziz/degolate"
)

// A decorator is declared as a function.
//
// Using `degolate.It` you can apply a decorator to a function. To be simpler
// and avoid code replication, just put `degolate.It` directly in the return of
// a decorator, as you can see below.
func Logging() degolate.D {
	return degolate.It(func() {
		fmt.Printf(
			"[%s]: %s\n",
			time.Now().Format(time.DateTime),
			"Hello World!",
		)
	})
}
