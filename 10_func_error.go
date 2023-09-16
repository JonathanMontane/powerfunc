package powerfunc

import (
	"fmt"
	"time"
)

// Func10Error is a function that takes 0 arguments and returns an error.
type Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error

// Exec executes the function and returns the error.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
	return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
}

// Timing returns a Func10 that will log the execution time of the Func10.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Timing(loggers ...func(d time.Duration)) Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
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
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Retry(tryAgain func(attempts int, err error) bool) Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		var err error
		attempts := 1
		for {
			err = f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func10 that will panic if the Func10Error returns an error.
func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Must() Func10[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) {
		err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a Func10Error that will wrap the error with the provided message.
func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) OnErr(msg string) Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry10(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) FuncError {
	return func() error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry9(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) Func1Error[P9] {
	return func(p9 P9) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) Func2Error[P8, P9] {
	return func(p8 P8, p9 P9) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) Func3Error[P7, P8, P9] {
	return func(p7 P7, p8 P8, p9 P9) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) Func4Error[P6, P7, P8, P9] {
	return func(p6 P6, p7 P7, p8 P8, p9 P9) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func5Error[P5, P6, P7, P8, P9] {
	return func(p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func6Error[P4, P5, P6, P7, P8, P9] {
	return func(p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry3(p0 P0, p1 P1, p2 P2) Func7Error[P3, P4, P5, P6, P7, P8, P9] {
	return func(p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry2(p0 P0, p1 P1) Func8Error[P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry1(p0 P0) Func9Error[P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	