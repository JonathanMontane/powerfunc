//go:generate go run ./generator -arity 10

package powerfunc

import (
	"fmt"
	"time"
)

// Func9 is a function that takes 0 arguments and returns no values.
type Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8)

// Exec executes the Function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) {
	f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

// Timing returns a Func9 that will log the execution time of the Func9.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Timing(loggers ...func(d time.Duration)) Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) {
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
		f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

// Fallible transforms a Func9 into a Func9Error.
// The returned Func9Error will never return an error.
// Useful when passing a Func9 to a function that expects a Func9Error.
func (f Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Fallible() Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
		return nil
	}
}


func (f Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry9(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) Func {
	return func()  {
		f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) Func1[P8] {
	return func(p8 P8)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) Func2[P7, P8] {
	return func(p7 P7, p8 P8)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) Func3[P6, P7, P8] {
	return func(p6 P6, p7 P7, p8 P8)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func4[P5, P6, P7, P8] {
	return func(p5 P5, p6 P6, p7 P7, p8 P8)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func5[P4, P5, P6, P7, P8] {
	return func(p4 P4, p5 P5, p6 P6, p7 P7, p8 P8)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry3(p0 P0, p1 P1, p2 P2) Func6[P3, P4, P5, P6, P7, P8] {
	return func(p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry2(p0 P0, p1 P1) Func7[P2, P3, P4, P5, P6, P7, P8] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry1(p0 P0) Func8[P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8)  {
		f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	