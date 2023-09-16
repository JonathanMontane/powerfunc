//go:generate go run ./generator -arity 10

package powerfunc

import (
	"fmt"
	"time"
)

// Func3 is a function that takes 0 arguments and returns no values.
type Func3[P0, P1, P2 any] func(p0 P0, p1 P1, p2 P2)

// Exec executes the Function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func3[P0, P1, P2]) Exec(p0 P0, p1 P1, p2 P2) {
	f(p0, p1, p2)
}

// Timing returns a Func3 that will log the execution time of the Func3.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func3[P0, P1, P2]) Timing(loggers ...func(d time.Duration)) Func3[P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) {
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
		f(p0, p1, p2)
	}
}

// Fallible transforms a Func3 into a Func3Error.
// The returned Func3Error will never return an error.
// Useful when passing a Func3 to a function that expects a Func3Error.
func (f Func3[P0, P1, P2]) Fallible() Func3Error[P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) error {
		f(p0, p1, p2)
		return nil
	}
}


func (f Func3[P0, P1, P2]) Curry3(p0 P0, p1 P1, p2 P2) Func {
	return func()  {
		f(p0, p1, p2)
	}
}
	

func (f Func3[P0, P1, P2]) Curry2(p0 P0, p1 P1) Func1[P2] {
	return func(p2 P2)  {
		f(p0, p1, p2)
	}
}
	

func (f Func3[P0, P1, P2]) Curry1(p0 P0) Func2[P1, P2] {
	return func(p1 P1, p2 P2)  {
		f(p0, p1, p2)
	}
}
	