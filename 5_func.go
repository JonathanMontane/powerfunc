//go:generate go run ./generator -arity 10

package powerfunc

import (
	"fmt"
	"time"
)

// Func5 is a function that takes 0 arguments and returns no values.
type Func5[P0, P1, P2, P3, P4 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4)

// Exec executes the Function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func5[P0, P1, P2, P3, P4]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) {
	f(p0, p1, p2, p3, p4)
}

// Timing returns a Func5 that will log the execution time of the Func5.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func5[P0, P1, P2, P3, P4]) Timing(loggers ...func(d time.Duration)) Func5[P0, P1, P2, P3, P4] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) {
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
		f(p0, p1, p2, p3, p4)
	}
}

// Fallible transforms a Func5 into a Func5Error.
// The returned Func5Error will never return an error.
// Useful when passing a Func5 to a function that expects a Func5Error.
func (f Func5[P0, P1, P2, P3, P4]) Fallible() Func5Error[P0, P1, P2, P3, P4] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) error {
		f(p0, p1, p2, p3, p4)
		return nil
	}
}


func (f Func5[P0, P1, P2, P3, P4]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func {
	return func()  {
		f(p0, p1, p2, p3, p4)
	}
}
	

func (f Func5[P0, P1, P2, P3, P4]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func1[P4] {
	return func(p4 P4)  {
		f(p0, p1, p2, p3, p4)
	}
}
	

func (f Func5[P0, P1, P2, P3, P4]) Curry3(p0 P0, p1 P1, p2 P2) Func2[P3, P4] {
	return func(p3 P3, p4 P4)  {
		f(p0, p1, p2, p3, p4)
	}
}
	

func (f Func5[P0, P1, P2, P3, P4]) Curry2(p0 P0, p1 P1) Func3[P2, P3, P4] {
	return func(p2 P2, p3 P3, p4 P4)  {
		f(p0, p1, p2, p3, p4)
	}
}
	

func (f Func5[P0, P1, P2, P3, P4]) Curry1(p0 P0) Func4[P1, P2, P3, P4] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4)  {
		f(p0, p1, p2, p3, p4)
	}
}
	