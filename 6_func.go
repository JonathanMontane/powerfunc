//go:generate go run ./generator -arity 10

package powerfunc

import (
	"fmt"
	"time"
)

// Func6 is a function that takes 0 arguments and returns no values.
type Func6[P0, P1, P2, P3, P4, P5 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5)

// Exec executes the Function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func6[P0, P1, P2, P3, P4, P5]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) {
	f(p0, p1, p2, p3, p4, p5)
}

// Timing returns a Func6 that will log the execution time of the Func6.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func6[P0, P1, P2, P3, P4, P5]) Timing(loggers ...func(d time.Duration)) Func6[P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) {
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
		f(p0, p1, p2, p3, p4, p5)
	}
}

// Fallible transforms a Func6 into a Func6Error.
// The returned Func6Error will never return an error.
// Useful when passing a Func6 to a function that expects a Func6Error.
func (f Func6[P0, P1, P2, P3, P4, P5]) Fallible() Func6Error[P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
		f(p0, p1, p2, p3, p4, p5)
		return nil
	}
}


func (f Func6[P0, P1, P2, P3, P4, P5]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) Func {
	return func()  {
		f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6[P0, P1, P2, P3, P4, P5]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func1[P5] {
	return func(p5 P5)  {
		f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6[P0, P1, P2, P3, P4, P5]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func2[P4, P5] {
	return func(p4 P4, p5 P5)  {
		f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6[P0, P1, P2, P3, P4, P5]) Curry3(p0 P0, p1 P1, p2 P2) Func3[P3, P4, P5] {
	return func(p3 P3, p4 P4, p5 P5)  {
		f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6[P0, P1, P2, P3, P4, P5]) Curry2(p0 P0, p1 P1) Func4[P2, P3, P4, P5] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5)  {
		f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6[P0, P1, P2, P3, P4, P5]) Curry1(p0 P0) Func5[P1, P2, P3, P4, P5] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5)  {
		f(p0, p1, p2, p3, p4, p5)
	}
}
	