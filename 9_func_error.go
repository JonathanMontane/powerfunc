package powerfunc

import (
	"fmt"
	"time"
)

// Func9Error is a function that takes 0 arguments and returns an error.
type Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error

// Exec executes the function and returns the error.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
	return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

// Timing returns a Func9 that will log the execution time of the Func9.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Timing(loggers ...func(d time.Duration)) Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
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
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Retry(tryAgain func(attempts int, err error) bool) Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		var err error
		attempts := 1
		for {
			err = f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func9 that will panic if the Func9Error returns an error.
func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Must() Func9[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) {
		err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a Func9Error that will wrap the error with the provided message.
func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) OnErr(msg string) Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry9(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) FuncError {
	return func() error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) Func1Error[P8] {
	return func(p8 P8) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) Func2Error[P7, P8] {
	return func(p7 P7, p8 P8) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) Func3Error[P6, P7, P8] {
	return func(p6 P6, p7 P7, p8 P8) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func4Error[P5, P6, P7, P8] {
	return func(p5 P5, p6 P6, p7 P7, p8 P8) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func5Error[P4, P5, P6, P7, P8] {
	return func(p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry3(p0 P0, p1 P1, p2 P2) Func6Error[P3, P4, P5, P6, P7, P8] {
	return func(p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry2(p0 P0, p1 P1) Func7Error[P2, P3, P4, P5, P6, P7, P8] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry1(p0 P0) Func8Error[P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	