package powerfunc

import (
	"fmt"
	"time"
)

// Func6Error is a function that takes 0 arguments and returns an error.
type Func6Error[P0, P1, P2, P3, P4, P5 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error

// Exec executes the function and returns the error.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func6Error[P0, P1, P2, P3, P4, P5]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
	return f(p0, p1, p2, p3, p4, p5)
}

// Timing returns a Func6 that will log the execution time of the Func6.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func6Error[P0, P1, P2, P3, P4, P5]) Timing(loggers ...func(d time.Duration)) Func6Error[P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
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
		return f(p0, p1, p2, p3, p4, p5)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func6Error[P0, P1, P2, P3, P4, P5]) Retry(tryAgain func(attempts int, err error) bool) Func6Error[P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
		var err error
		attempts := 1
		for {
			err = f(p0, p1, p2, p3, p4, p5)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func6 that will panic if the Func6Error returns an error.
func (f Func6Error[P0, P1, P2, P3, P4, P5]) Must() Func6[P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) {
		err := f(p0, p1, p2, p3, p4, p5)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a Func6Error that will wrap the error with the provided message.
func (f Func6Error[P0, P1, P2, P3, P4, P5]) OnErr(msg string) Func6Error[P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
		err := f(p0, p1, p2, p3, p4, p5)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f Func6Error[P0, P1, P2, P3, P4, P5]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) FuncError {
	return func() error {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Error[P0, P1, P2, P3, P4, P5]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func1Error[P5] {
	return func(p5 P5) error {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Error[P0, P1, P2, P3, P4, P5]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func2Error[P4, P5] {
	return func(p4 P4, p5 P5) error {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Error[P0, P1, P2, P3, P4, P5]) Curry3(p0 P0, p1 P1, p2 P2) Func3Error[P3, P4, P5] {
	return func(p3 P3, p4 P4, p5 P5) error {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Error[P0, P1, P2, P3, P4, P5]) Curry2(p0 P0, p1 P1) Func4Error[P2, P3, P4, P5] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5) error {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Error[P0, P1, P2, P3, P4, P5]) Curry1(p0 P0) Func5Error[P1, P2, P3, P4, P5] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	