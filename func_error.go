package powerfunc

import (
	"fmt"
	"time"
)

// FuncError is a function that takes 0 arguments and returns an error.
type FuncError func() error

// Exec executes the function and returns the error.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f FuncError) Exec() error {
	return f()
}

// Timing returns a Func that will log the execution time of the Func.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f FuncError) Timing(loggers ...func(d time.Duration)) FuncError {
	return func() error {
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

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f FuncError) Retry(tryAgain func(attempts int, err error) bool) FuncError {
	return func() error {
		var err error
		attempts := 1
		for {
			err = f()
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func that will panic if the FuncError returns an error.
func (f FuncError) Must() Func {
	return func() {
		err := f()
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a FuncError that will wrap the error with the provided message.
func (f FuncError) OnErr(msg string) FuncError {
	return func() error {
		err := f()
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}
