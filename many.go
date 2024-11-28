package degolate

import "slices"

// Just a simple way to apply many `decorators` to `function`
//
// See `./examples/basic/main.go`
func Many[F any](function F, decorators ...D) (fn F) {
	fn = function
	slices.Reverse(decorators)

	for _, dec := range decorators {
		fn = dec(fn).(F)
	}

	return
}
