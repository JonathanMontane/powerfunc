package powerfunc

import (
	"fmt"
	"time"
)

// Func1 is a function that takes 0 arguments and returns a value.
type Func1Value[T, P0 any] func(p0 P0) T

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func1Value[T, P0]) Exec(p0 P0) T {
	return f(p0)
}

// Timing returns a Func1 that will log the execution time of the Func1.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func1Value[T, P0]) Timing(loggers ...func(d time.Duration)) Func1Value[T, P0] {
	return func(p0 P0) T {
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
		return f(p0)
	}
}

// Fallible transforms a Func1Value into a Func1Result.
// The returned Func1Result will never return an error.
// Useful when passing a Func1Value to a function that expects a Func1Result.
func (f Func1Value[T, P0]) Fallible() Func1Result[T, P0] {
	return func(p0 P0) (T, error) {
		return f(p0), nil
	}
}


func (f Func1Value[R, P0]) Curry1(p0 P0) FuncValue[R] {
	return func() R {
		return f(p0)
	}
}
	