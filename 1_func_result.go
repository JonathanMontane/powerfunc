package powerfunc

import (
	"fmt"
	"time"
)

// Func1Result is a function that takes 0 arguments and returns a
// value and an error.
type Func1Result[T, P0 any] func(p0 P0) (T, error)

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func1Result[T, P0]) Exec(p0 P0) (T, error) {
	return f(p0)
}

// Timing returns a Func1 that will log the execution time of the Func1.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func1Result[T, P0]) Timing(loggers ...func(d time.Duration)) Func1Result[T, P0] {
	return func(p0 P0) (T, error) {
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
		return f(p0)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func1Result[T, P0]) Retry(tryAgain func(attempts int, err error) bool) Func1Result[T, P0] {
	return func(p0 P0) (T, error) {
		var v T
		var err error
		attempts := 1
		for {
			v, err = f(p0)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func1Value that will panic if the Func1Result returns an error.
func (f Func1Result[T, P0]) Must() Func1Value[T, P0] {
	return func(p0 P0) T {
		v, err := f(p0)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a Func1Result that will wrap the error returned by the Func1Result
// with the provided message.
func (f Func1Result[T, P0]) OnErr(msg string) Func1Result[T, P0] {
	return func(p0 P0) (T, error) {
		v, err := f(p0)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the Func1Result,
// if there is no error.
func (f Func1Result[T, P0]) Map(fn func(T) T) Func1Result[T, P0] {
	return func(p0 P0) (T, error) {
		v, err := f(p0)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the Func1Result,
// if there is an error.
func (f Func1Result[T, P0]) MapErr(fn func(error) error) Func1Result[T, P0] {
	return func(p0 P0) (T, error) {
		v, err := f(p0)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func1Value that will return the provided value if the
// Func1Result returns an error.
func (f Func1Result[T, P0]) Fallback(val T) Func1Value[T, P0] {
	return func(p0 P0) T {
		v, err := f(p0)
		if err != nil {
			return val
		}
		return v
	}
}


func (f Func1Result[R, P0]) Curry1(p0 P0) FuncResult[R] {
	return func() (R, error) {
		return f(p0)
	}
}
	