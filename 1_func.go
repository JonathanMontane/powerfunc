//go:generate go run ./generator -arity 10

package powerfunc

import (
	"fmt"
	"time"
)

// Func1 is a function that takes 0 arguments and returns no values.
type Func1[P0 any] func(p0 P0)

// Exec executes the Function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func1[P0]) Exec(p0 P0) {
	f(p0)
}

// Timing returns a Func1 that will log the execution time of the Func1.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func1[P0]) Timing(loggers ...func(d time.Duration)) Func1[P0] {
	return func(p0 P0) {
		start := time.Now()
		defer func() {
			dur := time.Since(start)
			for _, logger := range loggers {
				logger(dur)
			}
			if len(loggers) == 0 {
				// Default logger
				fmt.Println(dur)
			}
		}()
		f(p0)
	}
}

// Fallible transforms a Func1 into a Func1Error.
// The returned Func1Error will never return an error.
// Useful when passing a Func1 to a function that expects a Func1Error.
func (f Func1[P0]) Fallible() Func1Error[P0] {
	return func(p0 P0) error {
		f(p0)
		return nil
	}
}


func (f Func1[P0]) Curry1(p0 P0) Func {
	return func()  {
		f(p0)
	}
}
	