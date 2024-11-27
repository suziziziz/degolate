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
	// Decorators builder
	func(fn interface{}) {
		// See the description of `decorators.Logging` to understand how to
		// implement a decorator with degolate.
		decorators.Logging(fn)
		decorators.HelloDegolate(fn)
	},
	// Function to be decorated
	func(name string) (string, string) {
		return "Hey %s! Do you want to degolate? 🫵", name
	},
)

// Alternatively, you can apply these decorators directly without
// `degolate.Many`, as you can see commented below.
//
// var _, Greetings = decorators.Logging(greetings),
// 	decorators.HelloDegolate(greetings)

func main() {
	usr, _ := user.Current()

	// To use a decorated function, just call it.
	fmt.Printf(Greetings(strings.Split(usr.Name, " ")[0]))
}
