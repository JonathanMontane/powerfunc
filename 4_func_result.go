package powerfunc

import (
	"fmt"
	"time"
)

// Func4Result is a function that takes 0 arguments and returns a
// value and an error.
type Func4Result[T, P0, P1, P2, P3 any] func(p0 P0, p1 P1, p2 P2, p3 P3) (T, error)

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func4Result[T, P0, P1, P2, P3]) Exec(p0 P0, p1 P1, p2 P2, p3 P3) (T, error) {
	return f(p0, p1, p2, p3)
}

// Timing returns a Func4 that will log the execution time of the Func4.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func4Result[T, P0, P1, P2, P3]) Timing(loggers ...func(d time.Duration)) Func4Result[T, P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) (T, error) {
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
		return f(p0, p1, p2, p3)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f Func4Result[T, P0, P1, P2, P3]) Retry(tryAgain func(attempts int, err error) bool) Func4Result[T, P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) (T, error) {
		var v T
		var err error
		attempts := 1
		for {
			v, err = f(p0, p1, p2, p3)
			if err == nil || !tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func4Value that will panic if the Func4Result returns an error.
func (f Func4Result[T, P0, P1, P2, P3]) Must() Func4Value[T, P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) T {
		v, err := f(p0, p1, p2, p3)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a Func4Result that will wrap the error returned by the Func4Result
// with the provided message.
func (f Func4Result[T, P0, P1, P2, P3]) OnErr(msg string) Func4Result[T, P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) (T, error) {
		v, err := f(p0, p1, p2, p3)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the Func4Result,
// if there is no error.
func (f Func4Result[T, P0, P1, P2, P3]) Map(fn func(T) T) Func4Result[T, P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) (T, error) {
		v, err := f(p0, p1, p2, p3)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the Func4Result,
// if there is an error.
func (f Func4Result[T, P0, P1, P2, P3]) MapErr(fn func(error) error) Func4Result[T, P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) (T, error) {
		v, err := f(p0, p1, p2, p3)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func4Value that will return the provided value if the
// Func4Result returns an error.
func (f Func4Result[T, P0, P1, P2, P3]) Fallback(val T) Func4Value[T, P0, P1, P2, P3] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3) T {
		v, err := f(p0, p1, p2, p3)
		if err != nil {
			return val
		}
		return v
	}
}


func (f Func4Result[R, P0, P1, P2, P3]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) FuncResult[R] {
	return func() (R, error) {
		return f(p0, p1, p2, p3)
	}
}
	

func (f Func4Result[R, P0, P1, P2, P3]) Curry3(p0 P0, p1 P1, p2 P2) Func1Result[R, P3] {
	return func(p3 P3) (R, error) {
		return f(p0, p1, p2, p3)
	}
}
	

func (f Func4Result[R, P0, P1, P2, P3]) Curry2(p0 P0, p1 P1) Func2Result[R, P2, P3] {
	return func(p2 P2, p3 P3) (R, error) {
		return f(p0, p1, p2, p3)
	}
}
	

func (f Func4Result[R, P0, P1, P2, P3]) Curry1(p0 P0) Func3Result[R, P1, P2, P3] {
	return func(p1 P1, p2 P2, p3 P3) (R, error) {
		return f(p0, p1, p2, p3)
	}
}
	