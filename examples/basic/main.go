package main

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/suziziziz/degolate"
	"github.com/suziziziz/degolate/examples/basic/decorators"
)

// You can use `degolate.Many` to apply many decorators at once.
//
// `degolate.Many` will return the original function, but decorated.
var Greetings = degolate.Many(
	func(name string) (string, string) {
		return "Hey %s! Do you want to degolate? 🫵", name
	},
	decorators.Logging(),
	decorators.HelloDegolate("🤘"),
)

// Instead of `degolate.Many` you can call decorators directly, as you can see
// below.
//
// var Greetings = decorators.Logging()(
// 	decorators.HelloDegolate("🤘")(
// 		func(name string) (string, string) {
// 			return "Hey %s! Do you want to degolate? 🫵", name
// 		},
// 	),
// ).(func(name string) (string, string))
//
// NOTE: Applying too many decorators this way won't be efficient and is not
// recommended.

func main() {
	usr, _ := user.Current()

	// To use a decorated function, just call it.
	fmt.Printf(Greetings(strings.Split(usr.Name, " ")[0]))
}
