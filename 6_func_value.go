package powerfunc

import (
	"fmt"
	"time"
)

// Func6 is a function that takes 0 arguments and returns a value.
type Func6Value[T, P0, P1, P2, P3, P4, P5 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) T

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func6Value[T, P0, P1, P2, P3, P4, P5]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) T {
	return f(p0, p1, p2, p3, p4, p5)
}

// Timing returns a Func6 that will log the execution time of the Func6.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func6Value[T, P0, P1, P2, P3, P4, P5]) Timing(loggers ...func(d time.Duration)) Func6Value[T, P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) T {
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
		return f(p0, p1, p2, p3, p4, p5)
	}
}

// Fallible transforms a Func6Value into a Func6Result.
// The returned Func6Result will never return an error.
// Useful when passing a Func6Value to a function that expects a Func6Result.
func (f Func6Value[T, P0, P1, P2, P3, P4, P5]) Fallible() Func6Result[T, P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (T, error) {
		return f(p0, p1, p2, p3, p4, p5), nil
	}
}


func (f Func6Value[R, P0, P1, P2, P3, P4, P5]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) FuncValue[R] {
	return func() R {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Value[R, P0, P1, P2, P3, P4, P5]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func1Value[R, P5] {
	return func(p5 P5) R {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Value[R, P0, P1, P2, P3, P4, P5]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func2Value[R, P4, P5] {
	return func(p4 P4, p5 P5) R {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Value[R, P0, P1, P2, P3, P4, P5]) Curry3(p0 P0, p1 P1, p2 P2) Func3Value[R, P3, P4, P5] {
	return func(p3 P3, p4 P4, p5 P5) R {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Value[R, P0, P1, P2, P3, P4, P5]) Curry2(p0 P0, p1 P1) Func4Value[R, P2, P3, P4, P5] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5) R {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Value[R, P0, P1, P2, P3, P4, P5]) Curry1(p0 P0) Func5Value[R, P1, P2, P3, P4, P5] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) R {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	