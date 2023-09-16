package powerfunc

import (
	"fmt"
	"time"
)

// Func7 is a function that takes 0 arguments and returns a value.
type Func7Value[T, P0, P1, P2, P3, P4, P5, P6 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) T

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func7Value[T, P0, P1, P2, P3, P4, P5, P6]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) T {
	return f(p0, p1, p2, p3, p4, p5, p6)
}

// Timing returns a Func7 that will log the execution time of the Func7.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func7Value[T, P0, P1, P2, P3, P4, P5, P6]) Timing(loggers ...func(d time.Duration)) Func7Value[T, P0, P1, P2, P3, P4, P5, P6] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) T {
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
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}

// Fallible transforms a Func7Value into a Func7Result.
// The returned Func7Result will never return an error.
// Useful when passing a Func7Value to a function that expects a Func7Result.
func (f Func7Value[T, P0, P1, P2, P3, P4, P5, P6]) Fallible() Func7Result[T, P0, P1, P2, P3, P4, P5, P6] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) (T, error) {
		return f(p0, p1, p2, p3, p4, p5, p6), nil
	}
}


func (f Func7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) FuncValue[R] {
	return func() R {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) Func1Value[R, P6] {
	return func(p6 P6) R {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func2Value[R, P5, P6] {
	return func(p5 P5, p6 P6) R {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func3Value[R, P4, P5, P6] {
	return func(p4 P4, p5 P5, p6 P6) R {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry3(p0 P0, p1 P1, p2 P2) Func4Value[R, P3, P4, P5, P6] {
	return func(p3 P3, p4 P4, p5 P5, p6 P6) R {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry2(p0 P0, p1 P1) Func5Value[R, P2, P3, P4, P5, P6] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) R {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry1(p0 P0) Func6Value[R, P1, P2, P3, P4, P5, P6] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) R {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	