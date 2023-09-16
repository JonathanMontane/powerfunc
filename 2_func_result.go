package powerfunc

import (
	"fmt"
	"time"
)

// Func2Result is a function that takes 0 arguments and returns a
// value and an error.
type Func2Result[T, P0, P1 any] func(p0 P0, p1 P1) (T, error)

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func2Result[T, P0, P1]) Exec(p0 P0, p1 P1) (T, error) {
	return f(p0, p1)
}

// Timing returns a Func2 that will log the execution time of the Func2.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func2Result[T, P0, P1]) Timing(loggers ...func(d time.Duration)) Func2Result[T, P0, P1] {
	return func(p0 P0, p1 P1) (T, error) {
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
		return f(p0, p1)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func2Result[T, P0, P1]) Retry(tryAgain func(attempts int, err error) bool) Func2Result[T, P0, P1] {
	return func(p0 P0, p1 P1) (T, error) {
		var v T
		var err error
		attempts := 1
		for {
			v, err = f(p0, p1)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func2Value that will panic if the Func2Result returns an error.
func (f Func2Result[T, P0, P1]) Must() Func2Value[T, P0, P1] {
	return func(p0 P0, p1 P1) T {
		v, err := f(p0, p1)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a Func2Result that will wrap the error returned by the Func2Result
// with the provided message.
func (f Func2Result[T, P0, P1]) OnErr(msg string) Func2Result[T, P0, P1] {
	return func(p0 P0, p1 P1) (T, error) {
		v, err := f(p0, p1)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the Func2Result,
// if there is no error.
func (f Func2Result[T, P0, P1]) Map(fn func(T) T) Func2Result[T, P0, P1] {
	return func(p0 P0, p1 P1) (T, error) {
		v, err := f(p0, p1)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the Func2Result,
// if there is an error.
func (f Func2Result[T, P0, P1]) MapErr(fn func(error) error) Func2Result[T, P0, P1] {
	return func(p0 P0, p1 P1) (T, error) {
		v, err := f(p0, p1)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func2Value that will return the provided value if the
// Func2Result returns an error.
func (f Func2Result[T, P0, P1]) Fallback(val T) Func2Value[T, P0, P1] {
	return func(p0 P0, p1 P1) T {
		v, err := f(p0, p1)
		if err != nil {
			return val
		}
		return v
	}
}


func (f Func2Result[R, P0, P1]) Curry2(p0 P0, p1 P1) FuncResult[R] {
	return func() (R, error) {
		return f(p0, p1)
	}
}
	

func (f Func2Result[R, P0, P1]) Curry1(p0 P0) Func1Result[R, P1] {
	return func(p1 P1) (R, error) {
		return f(p0, p1)
	}
}
	