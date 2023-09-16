package powerfunc

import (
	"fmt"
	"time"
)

// Func7Error is a function that takes 0 arguments and returns an error.
type Func7Error[P0, P1, P2, P3, P4, P5, P6 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error

// Exec executes the function and returns the error.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func7Error[P0, P1, P2, P3, P4, P5, P6]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
	return f(p0, p1, p2, p3, p4, p5, p6)
}

// Timing returns a Func7 that will log the execution time of the Func7.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func7Error[P0, P1, P2, P3, P4, P5, P6]) Timing(loggers ...func(d time.Duration)) Func7Error[P0, P1, P2, P3, P4, P5, P6] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
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
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func7Error[P0, P1, P2, P3, P4, P5, P6]) Retry(tryAgain func(attempts int, err error) bool) Func7Error[P0, P1, P2, P3, P4, P5, P6] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
		var err error
		attempts := 1
		for {
			err = f(p0, p1, p2, p3, p4, p5, p6)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func7 that will panic if the Func7Error returns an error.
func (f Func7Error[P0, P1, P2, P3, P4, P5, P6]) Must() Func7[P0, P1, P2, P3, P4, P5, P6] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) {
		err := f(p0, p1, p2, p3, p4, p5, p6)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a Func7Error that will wrap the error with the provided message.
func (f Func7Error[P0, P1, P2, P3, P4, P5, P6]) OnErr(msg string) Func7Error[P0, P1, P2, P3, P4, P5, P6] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
		err := f(p0, p1, p2, p3, p4, p5, p6)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f Func7Error[P0, P1, P2, P3, P4, P5, P6]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) FuncError {
	return func() error {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7Error[P0, P1, P2, P3, P4, P5, P6]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) Func1Error[P6] {
	return func(p6 P6) error {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7Error[P0, P1, P2, P3, P4, P5, P6]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func2Error[P5, P6] {
	return func(p5 P5, p6 P6) error {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7Error[P0, P1, P2, P3, P4, P5, P6]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func3Error[P4, P5, P6] {
	return func(p4 P4, p5 P5, p6 P6) error {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7Error[P0, P1, P2, P3, P4, P5, P6]) Curry3(p0 P0, p1 P1, p2 P2) Func4Error[P3, P4, P5, P6] {
	return func(p3 P3, p4 P4, p5 P5, p6 P6) error {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7Error[P0, P1, P2, P3, P4, P5, P6]) Curry2(p0 P0, p1 P1) Func5Error[P2, P3, P4, P5, P6] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f Func7Error[P0, P1, P2, P3, P4, P5, P6]) Curry1(p0 P0) Func6Error[P1, P2, P3, P4, P5, P6] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
		return f(p0, p1, p2, p3, p4, p5, p6)
	}
}
	