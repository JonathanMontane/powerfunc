//go:generate go run ./generator -arity 10

package powerfunc

import (
	"fmt"
	"time"
)

// Func7 is a function that takes 0 arguments and returns no values.
type Func7[P0, P1, P2, P3, P4, P5, P6 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6)

// Exec executes the Function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func7[P0, P1, P2, P3, P4, P5, P6]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) {
	f(p0, p1, p2, p3, p4, p5, p6)
}

// Timing returns a Func7 that will log the execution time of the Func7.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func7[P0, P1, P2, P3, P4, P5, P6]) Timing(loggers ...func(d time.Duration)) Func7[P0, P1, P2, P3, P4, P5, P6] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) {
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
		f(p0, p1, p2, p3, p4, p5, p6)
	}
}

// Fallible transforms a Func7 into a Func7Error.
// The returned Func7Error will never return an error.
// Useful when passing a Func7 to a function that expects a Func7Error.
func (f Func7[P0, P1, P2, P3, P4, P5, P6]) Fallible() Func7Error[P0, P1, P2, P3, P4, P5, P6] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
		f(p0, p1, p2, p3, p4, p5, p6)
		return nil
	}
}


func (f Func7[P0, P1, P2, P3, P4, P5, P6]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) Func {
	return func()  {
		f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7[P0, P1, P2, P3, P4, P5, P6]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) Func1[P6] {
	return func(p6 P6)  {
		f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7[P0, P1, P2, P3, P4, P5, P6]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func2[P5, P6] {
	return func(p5 P5, p6 P6)  {
		f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7[P0, P1, P2, P3, P4, P5, P6]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func3[P4, P5, P6] {
	return func(p4 P4, p5 P5, p6 P6)  {
		f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7[P0, P1, P2, P3, P4, P5, P6]) Curry3(p0 P0, p1 P1, p2 P2) Func4[P3, P4, P5, P6] {
	return func(p3 P3, p4 P4, p5 P5, p6 P6)  {
		f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7[P0, P1, P2, P3, P4, P5, P6]) Curry2(p0 P0, p1 P1) Func5[P2, P3, P4, P5, P6] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5, p6 P6)  {
		f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7[P0, P1, P2, P3, P4, P5, P6]) Curry1(p0 P0) Func6[P1, P2, P3, P4, P5, P6] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6)  {
		f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	