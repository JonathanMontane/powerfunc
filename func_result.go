package powerfunc

import (
	"fmt"
	"time"
)

// FuncResult is a function that takes 0 arguments and returns a
// value and an error.
type FuncResult[T any] func() (T, error)

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f FuncResult[T]) Exec() (T, error) {
	return f()
}

// Timing returns a Func that will log the execution time of the Func.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f FuncResult[T]) Timing(loggers ...func(d time.Duration)) FuncResult[T] {
	return func() (T, error) {
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
		return f()
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f FuncResult[T]) Retry(tryAgain func(attempts int, err error) bool) FuncResult[T] {
	return func() (T, error) {
		var v T
		var err error
		attempts := 1
		for {
			v, err = f()
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a FuncValue that will panic if the FuncResult returns an error.
func (f FuncResult[T]) Must() FuncValue[T] {
	return func() T {
		v, err := f()
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a FuncResult that will wrap the error returned by the FuncResult
// with the provided message.
func (f FuncResult[T]) OnErr(msg string) FuncResult[T] {
	return func() (T, error) {
		v, err := f()
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the FuncResult,
// if there is no error.
func (f FuncResult[T]) Map(fn func(T) T) FuncResult[T] {
	return func() (T, error) {
		v, err := f()
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the FuncResult,
// if there is an error.
func (f FuncResult[T]) MapErr(fn func(error) error) FuncResult[T] {
	return func() (T, error) {
		v, err := f()
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a FuncValue that will return the provided value if the
// FuncResult returns an error.
func (f FuncResult[T]) Fallback(val T) FuncValue[T] {
	return func() T {
		v, err := f()
		if err != nil {
			return val
		}
		return v
	}
}
