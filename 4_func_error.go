package powerfunc

import (
	"fmt"
	"time"
)

// Func4Error is a function that takes 0 arguments and returns an error.
type Func4Error[P0, P1, P2, P3 any] func(p0 P0, p1 P1, p2 P2, p3 P3) error

// Exec executes the function and returns the error.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func4Error[P0, P1, P2, P3]) Exec(p0 P0, p1 P1, p2 P2, p3 P3) error {
	return f(p0, p1, p2, p3)
}

// Timing returns a Func4 that will log the execution time of the Func4.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func4Error[P0, P1, P2, P3]) Timing(loggers ...func(d time.Duration)) Func4Error[P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) error {
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

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func4Error[P0, P1, P2, P3]) Retry(tryAgain func(attempts int, err error) bool) Func4Error[P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) error {
		var err error
		attempts := 1
		for {
			err = f(p0, p1, p2, p3)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func4 that will panic if the Func4Error returns an error.
func (f Func4Error[P0, P1, P2, P3]) Must() Func4[P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) {
		err := f(p0, p1, p2, p3)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a Func4Error that will wrap the error with the provided message.
func (f Func4Error[P0, P1, P2, P3]) OnErr(msg string) Func4Error[P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) error {
		err := f(p0, p1, p2, p3)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f Func4Error[P0, P1, P2, P3]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) FuncError {
	return func() error {
		return f(p0, p1, p2, p3)
	}
}
	

func (f Func4Error[P0, P1, P2, P3]) Curry3(p0 P0, p1 P1, p2 P2) Func1Error[P3] {
	return func(p3 P3) error {
		return f(p0, p1, p2, p3)
	}
}
	

func (f Func4Error[P0, P1, P2, P3]) Curry2(p0 P0, p1 P1) Func2Error[P2, P3] {
	return func(p2 P2, p3 P3) error {
		return f(p0, p1, p2, p3)
	}
}
	

func (f Func4Error[P0, P1, P2, P3]) Curry1(p0 P0) Func3Error[P1, P2, P3] {
	return func(p1 P1, p2 P2, p3 P3) error {
		return f(p0, p1, p2, p3)
	}
}
	