package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFuncResult[R any] func(ctx context.Context) (R, error)

func (f CtxFuncResult[R]) Exec(ctx context.Context) (R, error) {
	return f(ctx)
}

func (f CtxFuncResult[R]) Timing(loggers ...func(d time.Duration)) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
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
		return f(ctx)
	}
}

func (f CtxFuncResult[R]) WithTimeout(timeout time.Duration) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx)
	}
}

func (f CtxFuncResult[R]) WithDeadline(deadline time.Time) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx)
	}
}

func (f CtxFuncResult[R]) WithCancel() CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFuncResult[R]) Retry(tryAgain func(attempts int, err error) bool) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		var v R
		var err error
		attempts := 1
		for {
			v, err = f(ctx)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a FuncValue that will panic if the CtxFuncResult returns an error.
func (f CtxFuncResult[R]) Must() CtxFuncValue[R] {
	return func(ctx context.Context) R {
		v, err := f(ctx)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a CtxFuncResult that will wrap the error returned by the CtxFuncResult
// with the provided message.
func (f CtxFuncResult[R]) OnErr(msg string) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		v, err := f(ctx)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the CtxFuncResult,
// if there is no error.
func (f CtxFuncResult[R]) Map(fn func(R) R) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		v, err := f(ctx)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the CtxFuncResult,
// if there is an error.
func (f CtxFuncResult[R]) MapErr(fn func(error) error) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		v, err := f(ctx)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a FuncValue that will return the provided value if the
// CtxFuncResult returns an error.
func (f CtxFuncResult[R]) Fallback(val R) CtxFuncValue[R] {
	return func(ctx context.Context) R {
		v, err := f(ctx)
		if err != nil {
			return val
		}
		return v
	}
}
