package powerfunc

import (
	"fmt"
	"time"
)

// Func5Error is a function that takes 0 arguments and returns an error.
type Func5Error[P0, P1, P2, P3, P4 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) error

// Exec executes the function and returns the error.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func5Error[P0, P1, P2, P3, P4]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) error {
	return f(p0, p1, p2, p3, p4)
}

// Timing returns a Func5 that will log the execution time of the Func5.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func5Error[P0, P1, P2, P3, P4]) Timing(loggers ...func(d time.Duration)) Func5Error[P0, P1, P2, P3, P4] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) error {
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
		return f(p0, p1, p2, p3, p4)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func5Error[P0, P1, P2, P3, P4]) Retry(tryAgain func(attempts int, err error) bool) Func5Error[P0, P1, P2, P3, P4] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) error {
		var err error
		attempts := 1
		for {
			err = f(p0, p1, p2, p3, p4)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func5 that will panic if the Func5Error returns an error.
func (f Func5Error[P0, P1, P2, P3, P4]) Must() Func5[P0, P1, P2, P3, P4] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) {
		err := f(p0, p1, p2, p3, p4)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a Func5Error that will wrap the error with the provided message.
func (f Func5Error[P0, P1, P2, P3, P4]) OnErr(msg string) Func5Error[P0, P1, P2, P3, P4] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) error {
		err := f(p0, p1, p2, p3, p4)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f Func5Error[P0, P1, P2, P3, P4]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) FuncError {
	return func() error {
		return f(p0, p1, p2, p3, p4)
	}
}
	

func (f Func5Error[P0, P1, P2, P3, P4]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func1Error[P4] {
	return func(p4 P4) error {
		return f(p0, p1, p2, p3, p4)
	}
}
	

func (f Func5Error[P0, P1, P2, P3, P4]) Curry3(p0 P0, p1 P1, p2 P2) Func2Error[P3, P4] {
	return func(p3 P3, p4 P4) error {
		return f(p0, p1, p2, p3, p4)
	}
}
	

func (f Func5Error[P0, P1, P2, P3, P4]) Curry2(p0 P0, p1 P1) Func3Error[P2, P3, P4] {
	return func(p2 P2, p3 P3, p4 P4) error {
		return f(p0, p1, p2, p3, p4)
	}
}
	

func (f Func5Error[P0, P1, P2, P3, P4]) Curry1(p0 P0) Func4Error[P1, P2, P3, P4] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4) error {
		return f(p0, p1, p2, p3, p4)
	}
}
	