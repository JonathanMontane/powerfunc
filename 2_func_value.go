package powerfunc

import (
	"fmt"
	"time"
)

// Func2 is a function that takes 0 arguments and returns a value.
type Func2Value[T, P0, P1 any] func(p0 P0, p1 P1) T

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func2Value[T, P0, P1]) Exec(p0 P0, p1 P1) T {
	return f(p0, p1)
}

// Timing returns a Func2 that will log the execution time of the Func2.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func2Value[T, P0, P1]) Timing(loggers ...func(d time.Duration)) Func2Value[T, P0, P1] {
	return func(p0 P0, p1 P1) T {
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
		return f(p0, p1)
	}
}

// Fallible transforms a Func2Value into a Func2Result.
// The returned Func2Result will never return an error.
// Useful when passing a Func2Value to a function that expects a Func2Result.
func (f Func2Value[T, P0, P1]) Fallible() Func2Result[T, P0, P1] {
	return func(p0 P0, p1 P1) (T, error) {
		return f(p0, p1), nil
	}
}


func (f Func2Value[R, P0, P1]) Curry2(p0 P0, p1 P1) FuncValue[R] {
	return func() R {
		return f(p0, p1)
	}
}
	

func (f Func2Value[R, P0, P1]) Curry1(p0 P0) Func1Value[R, P1] {
	return func(p1 P1) R {
		return f(p0, p1)
	}
}
	