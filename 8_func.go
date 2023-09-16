//go:generate go run ./generator -arity 10

package powerfunc

import (
	"fmt"
	"time"
)

// Func8 is a function that takes 0 arguments and returns no values.
type Func8[P0, P1, P2, P3, P4, P5, P6, P7 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7)

// Exec executes the Function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func8[P0, P1, P2, P3, P4, P5, P6, P7]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) {
	f(p0, p1, p2, p3, p4, p5, p6, p7)
}

// Timing returns a Func8 that will log the execution time of the Func8.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func8[P0, P1, P2, P3, P4, P5, P6, P7]) Timing(loggers ...func(d time.Duration)) Func8[P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) {
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
		f(p0, p1, p2, p3, p4, p5, p6, p7)
	}
}

// Fallible transforms a Func8 into a Func8Error.
// The returned Func8Error will never return an error.
// Useful when passing a Func8 to a function that expects a Func8Error.
func (f Func8[P0, P1, P2, P3, P4, P5, P6, P7]) Fallible() Func8Error[P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) error {
		f(p0, p1, p2, p3, p4, p5, p6, p7)
		return nil
	}
}


func (f Func8[P0, P1, P2, P3, P4, P5, P6, P7]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) Func {
	return func()  {
		f(p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f Func8[P0, P1, P2, P3, P4, P5, P6, P7]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) Func1[P7] {
	return func(p7 P7)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f Func8[P0, P1, P2, P3, P4, P5, P6, P7]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) Func2[P6, P7] {
	return func(p6 P6, p7 P7)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f Func8[P0, P1, P2, P3, P4, P5, P6, P7]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func3[P5, P6, P7] {
	return func(p5 P5, p6 P6, p7 P7)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f Func8[P0, P1, P2, P3, P4, P5, P6, P7]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func4[P4, P5, P6, P7] {
	return func(p4 P4, p5 P5, p6 P6, p7 P7)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f Func8[P0, P1, P2, P3, P4, P5, P6, P7]) Curry3(p0 P0, p1 P1, p2 P2) Func5[P3, P4, P5, P6, P7] {
	return func(p3 P3, p4 P4, p5 P5, p6 P6, p7 P7)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f Func8[P0, P1, P2, P3, P4, P5, P6, P7]) Curry2(p0 P0, p1 P1) Func6[P2, P3, P4, P5, P6, P7] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f Func8[P0, P1, P2, P3, P4, P5, P6, P7]) Curry1(p0 P0) Func7[P1, P2, P3, P4, P5, P6, P7] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	