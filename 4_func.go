//go:generate go run ./generator -arity 10

package powerfunc

import (
	"fmt"
	"time"
)

// Func4 is a function that takes 0 arguments and returns no values.
type Func4[P0, P1, P2, P3 any] func(p0 P0, p1 P1, p2 P2, p3 P3)

// Exec executes the Function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func4[P0, P1, P2, P3]) Exec(p0 P0, p1 P1, p2 P2, p3 P3) {
	f(p0, p1, p2, p3)
}

// Timing returns a Func4 that will log the execution time of the Func4.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func4[P0, P1, P2, P3]) Timing(loggers ...func(d time.Duration)) Func4[P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) {
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
		f(p0, p1, p2, p3)
	}
}

// Fallible transforms a Func4 into a Func4Error.
// The returned Func4Error will never return an error.
// Useful when passing a Func4 to a function that expects a Func4Error.
func (f Func4[P0, P1, P2, P3]) Fallible() Func4Error[P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) error {
		f(p0, p1, p2, p3)
		return nil
	}
}


func (f Func4[P0, P1, P2, P3]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func {
	return func()  {
		f(p0, p1, p2, p3)
	}
}
	

func (f Func4[P0, P1, P2, P3]) Curry3(p0 P0, p1 P1, p2 P2) Func1[P3] {
	return func(p3 P3)  {
		f(p0, p1, p2, p3)
	}
}
	

func (f Func4[P0, P1, P2, P3]) Curry2(p0 P0, p1 P1) Func2[P2, P3] {
	return func(p2 P2, p3 P3)  {
		f(p0, p1, p2, p3)
	}
}
	

func (f Func4[P0, P1, P2, P3]) Curry1(p0 P0) Func3[P1, P2, P3] {
	return func(p1 P1, p2 P2, p3 P3)  {
		f(p0, p1, p2, p3)
	}
}
	