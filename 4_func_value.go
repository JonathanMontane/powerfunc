package powerfunc

import (
	"fmt"
	"time"
)

// Func4 is a function that takes 0 arguments and returns a value.
type Func4Value[T, P0, P1, P2, P3 any] func(p0 P0, p1 P1, p2 P2, p3 P3) T

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func4Value[T, P0, P1, P2, P3]) Exec(p0 P0, p1 P1, p2 P2, p3 P3) T {
	return f(p0, p1, p2, p3)
}

// Timing returns a Func4 that will log the execution time of the Func4.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func4Value[T, P0, P1, P2, P3]) Timing(loggers ...func(d time.Duration)) Func4Value[T, P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) T {
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
		return f(p0, p1, p2, p3)
	}
}

// Fallible transforms a Func4Value into a Func4Result.
// The returned Func4Result will never return an error.
// Useful when passing a Func4Value to a function that expects a Func4Result.
func (f Func4Value[T, P0, P1, P2, P3]) Fallible() Func4Result[T, P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) (T, error) {
		return f(p0, p1, p2, p3), nil
	}
}


func (f Func4Value[R, P0, P1, P2, P3]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) FuncValue[R] {
	return func() R {
		return f(p0, p1, p2, p3)
	}
}
	

func (f Func4Value[R, P0, P1, P2, P3]) Curry3(p0 P0, p1 P1, p2 P2) Func1Value[R, P3] {
	return func(p3 P3) R {
		return f(p0, p1, p2, p3)
	}
}
	

func (f Func4Value[R, P0, P1, P2, P3]) Curry2(p0 P0, p1 P1) Func2Value[R, P2, P3] {
	return func(p2 P2, p3 P3) R {
		return f(p0, p1, p2, p3)
	}
}
	

func (f Func4Value[R, P0, P1, P2, P3]) Curry1(p0 P0) Func3Value[R, P1, P2, P3] {
	return func(p1 P1, p2 P2, p3 P3) R {
		return f(p0, p1, p2, p3)
	}
}
	