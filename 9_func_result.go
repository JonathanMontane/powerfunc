package powerfunc

import (
	"fmt"
	"time"
)

// Func9Result is a function that takes 0 arguments and returns a
// value and an error.
type Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (T, error)

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (T, error) {
	return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

// Timing returns a Func9 that will log the execution time of the Func9.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Timing(loggers ...func(d time.Duration)) Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (T, error) {
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
func (f Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Retry(tryAgain func(attempts int, err error) bool) Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (T, error) {
		var v T
		var err error
		attempts := 1
		for {
			v, err = f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func9Value that will panic if the Func9Result returns an error.
func (f Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Must() Func9Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) T {
		v, err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a Func9Result that will wrap the error returned by the Func9Result
// with the provided message.
func (f Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8]) OnErr(msg string) Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (T, error) {
		v, err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the Func9Result,
// if there is no error.
func (f Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Map(fn func(T) T) Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (T, error) {
		v, err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the Func9Result,
// if there is an error.
func (f Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8]) MapErr(fn func(error) error) Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (T, error) {
		v, err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func9Value that will return the provided value if the
// Func9Result returns an error.
func (f Func9Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Fallback(val T) Func9Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) T {
		v, err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
		if err != nil {
			return val
		}
		return v
	}
}


func (f Func9Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry9(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) FuncResult[R] {
	return func() (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) Func1Result[R, P8] {
	return func(p8 P8) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) Func2Result[R, P7, P8] {
	return func(p7 P7, p8 P8) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) Func3Result[R, P6, P7, P8] {
	return func(p6 P6, p7 P7, p8 P8) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func4Result[R, P5, P6, P7, P8] {
	return func(p5 P5, p6 P6, p7 P7, p8 P8) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func5Result[R, P4, P5, P6, P7, P8] {
	return func(p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry3(p0 P0, p1 P1, p2 P2) Func6Result[R, P3, P4, P5, P6, P7, P8] {
	return func(p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry2(p0 P0, p1 P1) Func7Result[R, P2, P3, P4, P5, P6, P7, P8] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f Func9Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry1(p0 P0) Func8Result[R, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	