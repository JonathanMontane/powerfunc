package powerfunc

import (
	"fmt"
	"time"
)

// Func2Error is a function that takes 0 arguments and returns an error.
type Func2Error[P0, P1 any] func(p0 P0, p1 P1) error

// Exec executes the function and returns the error.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func2Error[P0, P1]) Exec(p0 P0, p1 P1) error {
	return f(p0, p1)
}

// Timing returns a Func2 that will log the execution time of the Func2.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func2Error[P0, P1]) Timing(loggers ...func(d time.Duration)) Func2Error[P0, P1] {
	return func(p0 P0, p1 P1) error {
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

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func2Error[P0, P1]) Retry(tryAgain func(attempts int, err error) bool) Func2Error[P0, P1] {
	return func(p0 P0, p1 P1) error {
		var err error
		attempts := 1
		for {
			err = f(p0, p1)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func2 that will panic if the Func2Error returns an error.
func (f Func2Error[P0, P1]) Must() Func2[P0, P1] {
	return func(p0 P0, p1 P1) {
		err := f(p0, p1)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a Func2Error that will wrap the error with the provided message.
func (f Func2Error[P0, P1]) OnErr(msg string) Func2Error[P0, P1] {
	return func(p0 P0, p1 P1) error {
		err := f(p0, p1)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f Func2Error[P0, P1]) Curry2(p0 P0, p1 P1) FuncError {
	return func() error {
		return f(p0, p1)
	}
}
	

func (f Func2Error[P0, P1]) Curry1(p0 P0) Func1Error[P1] {
	return func(p1 P1) error {
		return f(p0, p1)
	}
}
	