package powerfunc

import (
	"fmt"
	"time"
)

// Func9 is a function that takes 0 arguments and returns a value.
type Func9Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) T

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func9Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) T {
	return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

// Timing returns a Func9 that will log the execution time of the Func9.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func9Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Timing(loggers ...func(d time.Duration)) Func9Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) T {
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
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

// Fallible transforms a Func9Value into a Func9Result.
// The returned Func9Result will never return an error.
// Useful when passing a Func9Value to a function that expects a Func9Result.
func (f Func9Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Fallible() Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (T, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8), nil
	}
}


func (f Func9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry9(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) FuncValue[R] {
	return func() R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) Func1Value[R, P8] {
	return func(p8 P8) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) Func2Value[R, P7, P8] {
	return func(p7 P7, p8 P8) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) Func3Value[R, P6, P7, P8] {
	return func(p6 P6, p7 P7, p8 P8) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func4Value[R, P5, P6, P7, P8] {
	return func(p5 P5, p6 P6, p7 P7, p8 P8) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func5Value[R, P4, P5, P6, P7, P8] {
	return func(p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry3(p0 P0, p1 P1, p2 P2) Func6Value[R, P3, P4, P5, P6, P7, P8] {
	return func(p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry2(p0 P0, p1 P1) Func7Value[R, P2, P3, P4, P5, P6, P7, P8] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry1(p0 P0) Func8Value[R, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	