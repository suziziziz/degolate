package main

import (
	"fmt"
	"os/user"
	"reflect"
	"runtime"
	"strings"
	"time"

	"github.com/suziziziz/degolate"
)

func main() {
	usr, _ := user.Current()

	fmt.Printf(Greetings(strings.Split(usr.Name, " ")[0]))
}

var Greetings = degolate.Many(
	func(fn interface{}) {
		logging(fn)
		helloDegolate(fn)
	},
	greetings,
)

func greetings(name string) (string, string) {
	return "Hey %s! Do you want to degolate? 🫵", name
}

func helloDegolate[F any](fn F) F {
	return degolate.It(fn, func() {
		fmt.Println("LET'S DEGOLATE IT ALL! 🤘")
	})
}

func logging[F any](fn F) F {
	return degolate.It(fn, func() {
		fmt.Printf("[%s]: %s %T\n", time.Now(), runtime.FuncForPC(reflect.ValueOf(fn).Pointer()).Name(), fn)
	})
}
