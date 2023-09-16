//go:generate go run ./generator -arity 10

package powerfunc

import (
	"fmt"
	"time"
)

// Func2 is a function that takes 0 arguments and returns no values.
type Func2[P0, P1 any] func(p0 P0, p1 P1)

// Exec executes the Function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func2[P0, P1]) Exec(p0 P0, p1 P1) {
	f(p0, p1)
}

// Timing returns a Func2 that will log the execution time of the Func2.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func2[P0, P1]) Timing(loggers ...func(d time.Duration)) Func2[P0, P1] {
	return func(p0 P0, p1 P1) {
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
		f(p0, p1)
	}
}

// Fallible transforms a Func2 into a Func2Error.
// The returned Func2Error will never return an error.
// Useful when passing a Func2 to a function that expects a Func2Error.
func (f Func2[P0, P1]) Fallible() Func2Error[P0, P1] {
	return func(p0 P0, p1 P1) error {
		f(p0, p1)
		return nil
	}
}


func (f Func2[P0, P1]) Curry2(p0 P0, p1 P1) Func {
	return func()  {
		f(p0, p1)
	}
}
	

func (f Func2[P0, P1]) Curry1(p0 P0) Func1[P1] {
	return func(p1 P1)  {
		f(p0, p1)
	}
}
	