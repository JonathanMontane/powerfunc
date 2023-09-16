package powerfunc

import (
	"fmt"
	"time"
)

// Func5 is a function that takes 0 arguments and returns a value.
type Func5Value[T, P0, P1, P2, P3, P4 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) T

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func5Value[T, P0, P1, P2, P3, P4]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) T {
	return f(p0, p1, p2, p3, p4)
}

// Timing returns a Func5 that will log the execution time of the Func5.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func5Value[T, P0, P1, P2, P3, P4]) Timing(loggers ...func(d time.Duration)) Func5Value[T, P0, P1, P2, P3, P4] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) T {
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
		return f(p0, p1, p2, p3, p4)
	}
}

// Fallible transforms a Func5Value into a Func5Result.
// The returned Func5Result will never return an error.
// Useful when passing a Func5Value to a function that expects a Func5Result.
func (f Func5Value[T, P0, P1, P2, P3, P4]) Fallible() Func5Result[T, P0, P1, P2, P3, P4] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (T, error) {
		return f(p0, p1, p2, p3, p4), nil
	}
}


func (f Func5Value[R, P0, P1, P2, P3, P4]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) FuncValue[R] {
	return func() R {
		return f(p0, p1, p2, p3, p4)
	}
}
	

func (f Func5Value[R, P0, P1, P2, P3, P4]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func1Value[R, P4] {
	return func(p4 P4) R {
		return f(p0, p1, p2, p3, p4)
	}
}
	

func (f Func5Value[R, P0, P1, P2, P3, P4]) Curry3(p0 P0, p1 P1, p2 P2) Func2Value[R, P3, P4] {
	return func(p3 P3, p4 P4) R {
		return f(p0, p1, p2, p3, p4)
	}
}
	

func (f Func5Value[R, P0, P1, P2, P3, P4]) Curry2(p0 P0, p1 P1) Func3Value[R, P2, P3, P4] {
	return func(p2 P2, p3 P3, p4 P4) R {
		return f(p0, p1, p2, p3, p4)
	}
}
	

func (f Func5Value[R, P0, P1, P2, P3, P4]) Curry1(p0 P0) Func4Value[R, P1, P2, P3, P4] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4) R {
		return f(p0, p1, p2, p3, p4)
	}
}
	