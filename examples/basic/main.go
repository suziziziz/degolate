package main

import (
	"fmt"
	"os/user"
	"strings"

	"github.com/suziziziz/degolate"
	"github.com/suziziziz/degolate/examples/basic/decorators"
)

func main() {
	usr, _ := user.Current()

	fmt.Printf(Greetings(strings.Split(usr.Name, " ")[0]))
}

var Greetings = degolate.Many(
	func(fn interface{}) {
		decorators.Logging(fn)
		decorators.HelloDegolate(fn)
	},
	greetings,
)

func greetings(name string) (string, string) {
	return "Hey %s! Do you want to degolate? 🫵", name
}
