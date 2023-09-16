package powerfunc

import (
	"fmt"
	"time"
)

// Func is a function that takes 0 arguments and returns a value.
type FuncValue[T any] func() T

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f FuncValue[T]) Exec() T {
	return f()
}

// Timing returns a Func that will log the execution time of the Func.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f FuncValue[T]) Timing(loggers ...func(d time.Duration)) FuncValue[T] {
	return func() T {
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
		return f()
	}
}

// Fallible transforms a FuncValue into a FuncResult.
// The returned FuncResult will never return an error.
// Useful when passing a FuncValue to a function that expects a FuncResult.
func (f FuncValue[T]) Fallible() FuncResult[T] {
	return func() (T, error) {
		return f(), nil
	}
}
