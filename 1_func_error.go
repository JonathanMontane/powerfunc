package powerfunc

import (
	"fmt"
	"time"
)

// Func1Error is a function that takes 0 arguments and returns an error.
type Func1Error[P0 any] func(p0 P0) error

// Exec executes the function and returns the error.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func1Error[P0]) Exec(p0 P0) error {
	return f(p0)
}

// Timing returns a Func1 that will log the execution time of the Func1.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func1Error[P0]) Timing(loggers ...func(d time.Duration)) Func1Error[P0] {
	return func(p0 P0) error {
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

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func1Error[P0]) Retry(tryAgain func(attempts int, err error) bool) Func1Error[P0] {
	return func(p0 P0) error {
		var err error
		attempts := 1
		for {
			err = f(p0)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func1 that will panic if the Func1Error returns an error.
func (f Func1Error[P0]) Must() Func1[P0] {
	return func(p0 P0) {
		err := f(p0)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a Func1Error that will wrap the error with the provided message.
func (f Func1Error[P0]) OnErr(msg string) Func1Error[P0] {
	return func(p0 P0) error {
		err := f(p0)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f Func1Error[P0]) Curry1(p0 P0) FuncError {
	return func() error {
		return f(p0)
	}
}
	