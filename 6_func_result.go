package powerfunc

import (
	"fmt"
	"time"
)

// Func6Result is a function that takes 0 arguments and returns a
// value and an error.
type Func6Result[T, P0, P1, P2, P3, P4, P5 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (T, error)

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func6Result[T, P0, P1, P2, P3, P4, P5]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (T, error) {
	return f(p0, p1, p2, p3, p4, p5)
}

// Timing returns a Func6 that will log the execution time of the Func6.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func6Result[T, P0, P1, P2, P3, P4, P5]) Timing(loggers ...func(d time.Duration)) Func6Result[T, P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (T, error) {
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
func (f Func6Result[T, P0, P1, P2, P3, P4, P5]) Retry(tryAgain func(attempts int, err error) bool) Func6Result[T, P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (T, error) {
		var v T
		var err error
		attempts := 1
		for {
			v, err = f(p0, p1, p2, p3, p4, p5)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func6Value that will panic if the Func6Result returns an error.
func (f Func6Result[T, P0, P1, P2, P3, P4, P5]) Must() Func6Value[T, P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) T {
		v, err := f(p0, p1, p2, p3, p4, p5)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a Func6Result that will wrap the error returned by the Func6Result
// with the provided message.
func (f Func6Result[T, P0, P1, P2, P3, P4, P5]) OnErr(msg string) Func6Result[T, P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (T, error) {
		v, err := f(p0, p1, p2, p3, p4, p5)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the Func6Result,
// if there is no error.
func (f Func6Result[T, P0, P1, P2, P3, P4, P5]) Map(fn func(T) T) Func6Result[T, P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (T, error) {
		v, err := f(p0, p1, p2, p3, p4, p5)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the Func6Result,
// if there is an error.
func (f Func6Result[T, P0, P1, P2, P3, P4, P5]) MapErr(fn func(error) error) Func6Result[T, P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (T, error) {
		v, err := f(p0, p1, p2, p3, p4, p5)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func6Value that will return the provided value if the
// Func6Result returns an error.
func (f Func6Result[T, P0, P1, P2, P3, P4, P5]) Fallback(val T) Func6Value[T, P0, P1, P2, P3, P4, P5] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) T {
		v, err := f(p0, p1, p2, p3, p4, p5)
		if err != nil {
			return val
		}
		return v
	}
}


func (f Func6Result[R, P0, P1, P2, P3, P4, P5]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) FuncResult[R] {
	return func() (R, error) {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Result[R, P0, P1, P2, P3, P4, P5]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func1Result[R, P5] {
	return func(p5 P5) (R, error) {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Result[R, P0, P1, P2, P3, P4, P5]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func2Result[R, P4, P5] {
	return func(p4 P4, p5 P5) (R, error) {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Result[R, P0, P1, P2, P3, P4, P5]) Curry3(p0 P0, p1 P1, p2 P2) Func3Result[R, P3, P4, P5] {
	return func(p3 P3, p4 P4, p5 P5) (R, error) {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Result[R, P0, P1, P2, P3, P4, P5]) Curry2(p0 P0, p1 P1) Func4Result[R, P2, P3, P4, P5] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	

func (f Func6Result[R, P0, P1, P2, P3, P4, P5]) Curry1(p0 P0) Func5Result[R, P1, P2, P3, P4, P5] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
		return f(p0, p1, p2, p3, p4, p5)
	}
}
	