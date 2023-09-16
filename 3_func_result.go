package powerfunc

import (
	"fmt"
	"time"
)

// Func3Result is a function that takes 0 arguments and returns a
// value and an error.
type Func3Result[T, P0, P1, P2 any] func(p0 P0, p1 P1, p2 P2) (T, error)

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func3Result[T, P0, P1, P2]) Exec(p0 P0, p1 P1, p2 P2) (T, error) {
	return f(p0, p1, p2)
}

// Timing returns a Func3 that will log the execution time of the Func3.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func3Result[T, P0, P1, P2]) Timing(loggers ...func(d time.Duration)) Func3Result[T, P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) (T, error) {
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
		return f(p0, p1, p2)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func3Result[T, P0, P1, P2]) Retry(tryAgain func(attempts int, err error) bool) Func3Result[T, P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) (T, error) {
		var v T
		var err error
		attempts := 1
		for {
			v, err = f(p0, p1, p2)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func3Value that will panic if the Func3Result returns an error.
func (f Func3Result[T, P0, P1, P2]) Must() Func3Value[T, P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) T {
		v, err := f(p0, p1, p2)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a Func3Result that will wrap the error returned by the Func3Result
// with the provided message.
func (f Func3Result[T, P0, P1, P2]) OnErr(msg string) Func3Result[T, P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) (T, error) {
		v, err := f(p0, p1, p2)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the Func3Result,
// if there is no error.
func (f Func3Result[T, P0, P1, P2]) Map(fn func(T) T) Func3Result[T, P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) (T, error) {
		v, err := f(p0, p1, p2)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the Func3Result,
// if there is an error.
func (f Func3Result[T, P0, P1, P2]) MapErr(fn func(error) error) Func3Result[T, P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) (T, error) {
		v, err := f(p0, p1, p2)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func3Value that will return the provided value if the
// Func3Result returns an error.
func (f Func3Result[T, P0, P1, P2]) Fallback(val T) Func3Value[T, P0, P1, P2] {
	return func(p0 P0, p1 P1, p2 P2) T {
		v, err := f(p0, p1, p2)
		if err != nil {
			return val
		}
		return v
	}
}


func (f Func3Result[R, P0, P1, P2]) Curry3(p0 P0, p1 P1, p2 P2) FuncResult[R] {
	return func() (R, error) {
		return f(p0, p1, p2)
	}
}
	

func (f Func3Result[R, P0, P1, P2]) Curry2(p0 P0, p1 P1) Func1Result[R, P2] {
	return func(p2 P2) (R, error) {
		return f(p0, p1, p2)
	}
}
	

func (f Func3Result[R, P0, P1, P2]) Curry1(p0 P0) Func2Result[R, P1, P2] {
	return func(p1 P1, p2 P2) (R, error) {
		return f(p0, p1, p2)
	}
}
	