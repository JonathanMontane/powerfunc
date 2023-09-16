package powerfunc

import (
	"fmt"
	"time"
)

// Func3Error is a function that takes 0 arguments and returns an error.
type Func3Error[P0, P1, P2 any] func(p0 P0, p1 P1, p2 P2) error

// Exec executes the function and returns the error.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func3Error[P0, P1, P2]) Exec(p0 P0, p1 P1, p2 P2) error {
	return f(p0, p1, p2)
}

// Timing returns a Func3 that will log the execution time of the Func3.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func3Error[P0, P1, P2]) Timing(loggers ...func(d time.Duration)) Func3Error[P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) error {
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

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func3Error[P0, P1, P2]) Retry(tryAgain func(attempts int, err error) bool) Func3Error[P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) error {
		var err error
		attempts := 1
		for {
			err = f(p0, p1, p2)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func3 that will panic if the Func3Error returns an error.
func (f Func3Error[P0, P1, P2]) Must() Func3[P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) {
		err := f(p0, p1, p2)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a Func3Error that will wrap the error with the provided message.
func (f Func3Error[P0, P1, P2]) OnErr(msg string) Func3Error[P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) error {
		err := f(p0, p1, p2)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f Func3Error[P0, P1, P2]) Curry3(p0 P0, p1 P1, p2 P2) FuncError {
	return func() error {
		return f(p0, p1, p2)
	}
}
	

func (f Func3Error[P0, P1, P2]) Curry2(p0 P0, p1 P1) Func1Error[P2] {
	return func(p2 P2) error {
		return f(p0, p1, p2)
	}
}
	

func (f Func3Error[P0, P1, P2]) Curry1(p0 P0) Func2Error[P1, P2] {
	return func(p1 P1, p2 P2) error {
		return f(p0, p1, p2)
	}
}
	