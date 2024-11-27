package decorators

import (
	"fmt"

	"github.com/suziziziz/degolate"
)

func HelloDegolate[F any](fn F) F {
	return degolate.It(fn, func() {
		fmt.Println("LET'S DEGOLATE IT ALL! 🤘")
	})
}
