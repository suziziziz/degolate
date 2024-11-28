package decorators

import (
	"fmt"

	"github.com/suziziziz/degolate"
)

func HelloDegolate(emoji string) degolate.D {
	return degolate.It(func() {
		fmt.Printf("LET'S DEGOLATE IT ALL! %s\n", emoji)
	})
}
