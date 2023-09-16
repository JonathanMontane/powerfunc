package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc1Result[R, P0 any] func(ctx context.Context, p0 P0) (R, error)

func (f CtxFunc1Result[R, P0]) Exec(ctx context.Context, p0 P0) (R, error) {
	return f(ctx, p0)
}

func (f CtxFunc1Result[R, P0]) Timing(loggers ...func(d time.Duration)) CtxFunc1Result[R, P0] {
	return func(ctx context.Context, p0 P0) (R, error) {
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
		return f(ctx, p0)
	}
}

func (f CtxFunc1Result[R, P0]) WithTimeout(timeout time.Duration) CtxFunc1Result[R, P0] {
	return func(ctx context.Context, p0 P0) (R, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0)
	}
}

func (f CtxFunc1Result[R, P0]) WithDeadline(deadline time.Time) CtxFunc1Result[R, P0] {
	return func(ctx context.Context, p0 P0) (R, error) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0)
	}
}

func (f CtxFunc1Result[R, P0]) WithCancel() CtxFunc1Result[R, P0] {
	return func(ctx context.Context, p0 P0) (R, error) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc1Result[R, P0]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc1Result[R, P0] {
	return func(ctx context.Context, p0 P0) (R, error) {
		var v R
		var err error
		attempts := 1
		for {
			v, err = f(ctx, p0)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func1Value that will panic if the CtxFunc1Result returns an error.
func (f CtxFunc1Result[R, P0]) Must() CtxFunc1Value[R, P0] {
	return func(ctx context.Context, p0 P0) R {
		v, err := f(ctx, p0)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a CtxFunc1Result that will wrap the error returned by the CtxFunc1Result
// with the provided message.
func (f CtxFunc1Result[R, P0]) OnErr(msg string) CtxFunc1Result[R, P0] {
	return func(ctx context.Context, p0 P0) (R, error) {
		v, err := f(ctx, p0)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the CtxFunc1Result,
// if there is no error.
func (f CtxFunc1Result[R, P0]) Map(fn func(R) R) CtxFunc1Result[R, P0] {
	return func(ctx context.Context, p0 P0) (R, error) {
		v, err := f(ctx, p0)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the CtxFunc1Result,
// if there is an error.
func (f CtxFunc1Result[R, P0]) MapErr(fn func(error) error) CtxFunc1Result[R, P0] {
	return func(ctx context.Context, p0 P0) (R, error) {
		v, err := f(ctx, p0)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func1Value that will return the provided value if the
// CtxFunc1Result returns an error.
func (f CtxFunc1Result[R, P0]) Fallback(val R) CtxFunc1Value[R, P0] {
	return func(ctx context.Context, p0 P0) R {
		v, err := f(ctx, p0)
		if err != nil {
			return val
		}
		return v
	}
}


func (f CtxFunc1Result[R, P0]) Curry1(p0 P0) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		return f(ctx, p0)
	}
}
	