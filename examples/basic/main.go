package main

import (
	"fmt"

	"github.com/suziziziz/degolate"
)

func main() {
	degolate.D(greetings)("You")
}

var _ = degolate.It(greetings, helloDegolate())

func greetings(name string) {
	fmt.Printf("%s and I will do this! 🫵", name)
}

func helloDegolate() func(fn interface{}) interface{} {
	return func(fn interface{}) interface{} {
		fmt.Println("LET'S DEGOLATE IT ALL! 🤘")

		return fn
	}
}
