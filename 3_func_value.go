package powerfunc

import (
	"fmt"
	"time"
)

// Func3 is a function that takes 0 arguments and returns a value.
type Func3Value[T, P0, P1, P2 any] func(p0 P0, p1 P1, p2 P2) T

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func3Value[T, P0, P1, P2]) Exec(p0 P0, p1 P1, p2 P2) T {
	return f(p0, p1, p2)
}

// Timing returns a Func3 that will log the execution time of the Func3.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func3Value[T, P0, P1, P2]) Timing(loggers ...func(d time.Duration)) Func3Value[T, P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) T {
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
		return f(p0, p1, p2)
	}
}

// Fallible transforms a Func3Value into a Func3Result.
// The returned Func3Result will never return an error.
// Useful when passing a Func3Value to a function that expects a Func3Result.
func (f Func3Value[T, P0, P1, P2]) Fallible() Func3Result[T, P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) (T, error) {
		return f(p0, p1, p2), nil
	}
}


func (f Func3Value[R, P0, P1, P2]) Curry3(p0 P0, p1 P1, p2 P2) FuncValue[R] {
	return func() R {
		return f(p0, p1, p2)
	}
}
	

func (f Func3Value[R, P0, P1, P2]) Curry2(p0 P0, p1 P1) Func1Value[R, P2] {
	return func(p2 P2) R {
		return f(p0, p1, p2)
	}
}
	

func (f Func3Value[R, P0, P1, P2]) Curry1(p0 P0) Func2Value[R, P1, P2] {
	return func(p1 P1, p2 P2) R {
		return f(p0, p1, p2)
	}
}
	