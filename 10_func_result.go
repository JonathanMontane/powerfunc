package powerfunc

import (
	"fmt"
	"time"
)

// Func10Result is a function that takes 0 arguments and returns a
// value and an error.
type Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (T, error)

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (T, error) {
	return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
}

// Timing returns a Func10 that will log the execution time of the Func10.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Timing(loggers ...func(d time.Duration)) Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (T, error) {
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
func (f Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Retry(tryAgain func(attempts int, err error) bool) Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (T, error) {
		var v T
		var err error
		attempts := 1
		for {
			v, err = f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func10Value that will panic if the Func10Result returns an error.
func (f Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Must() Func10Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) T {
		v, err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a Func10Result that will wrap the error returned by the Func10Result
// with the provided message.
func (f Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) OnErr(msg string) Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (T, error) {
		v, err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the Func10Result,
// if there is no error.
func (f Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Map(fn func(T) T) Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (T, error) {
		v, err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the Func10Result,
// if there is an error.
func (f Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) MapErr(fn func(error) error) Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (T, error) {
		v, err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func10Value that will return the provided value if the
// Func10Result returns an error.
func (f Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Fallback(val T) Func10Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) T {
		v, err := f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
		if err != nil {
			return val
		}
		return v
	}
}


func (f Func10Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry10(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) FuncResult[R] {
	return func() (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry9(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) Func1Result[R, P9] {
	return func(p9 P9) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) Func2Result[R, P8, P9] {
	return func(p8 P8, p9 P9) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) Func3Result[R, P7, P8, P9] {
	return func(p7 P7, p8 P8, p9 P9) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) Func4Result[R, P6, P7, P8, P9] {
	return func(p6 P6, p7 P7, p8 P8, p9 P9) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func5Result[R, P5, P6, P7, P8, P9] {
	return func(p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func6Result[R, P4, P5, P6, P7, P8, P9] {
	return func(p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry3(p0 P0, p1 P1, p2 P2) Func7Result[R, P3, P4, P5, P6, P7, P8, P9] {
	return func(p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry2(p0 P0, p1 P1) Func8Result[R, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry1(p0 P0) Func9Result[R, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (R, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	