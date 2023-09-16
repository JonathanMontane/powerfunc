//go:generate go run ./generator -arity 10

package powerfunc

import (
	"fmt"
	"time"
)

// Func is a function that takes 0 arguments and returns no values.
type Func func()

// Exec executes the Function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func) Exec() {
	f()
}

// Timing returns a Func that will log the execution time of the Func.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func) Timing(loggers ...func(d time.Duration)) Func {
	return func() {
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
		f()
	}
}

// Fallible transforms a Func into a FuncError.
// The returned FuncError will never return an error.
// Useful when passing a Func to a function that expects a FuncError.
func (f Func) Fallible() FuncError {
	return func() error {
		f()
		return nil
	}
}
