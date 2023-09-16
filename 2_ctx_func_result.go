package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc2Result[R, P0, P1 any] func(ctx context.Context, p0 P0, p1 P1) (R, error)

func (f CtxFunc2Result[R, P0, P1]) Exec(ctx context.Context, p0 P0, p1 P1) (R, error) {
	return f(ctx, p0, p1)
}

func (f CtxFunc2Result[R, P0, P1]) Timing(loggers ...func(d time.Duration)) CtxFunc2Result[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) (R, error) {
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
		return f(ctx, p0, p1)
	}
}

func (f CtxFunc2Result[R, P0, P1]) WithTimeout(timeout time.Duration) CtxFunc2Result[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) (R, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1)
	}
}

func (f CtxFunc2Result[R, P0, P1]) WithDeadline(deadline time.Time) CtxFunc2Result[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) (R, error) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1)
	}
}

func (f CtxFunc2Result[R, P0, P1]) WithCancel() CtxFunc2Result[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) (R, error) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc2Result[R, P0, P1]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc2Result[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) (R, error) {
		var v R
		var err error
		attempts := 1
		for {
			v, err = f(ctx, p0, p1)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func2Value that will panic if the CtxFunc2Result returns an error.
func (f CtxFunc2Result[R, P0, P1]) Must() CtxFunc2Value[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) R {
		v, err := f(ctx, p0, p1)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a CtxFunc2Result that will wrap the error returned by the CtxFunc2Result
// with the provided message.
func (f CtxFunc2Result[R, P0, P1]) OnErr(msg string) CtxFunc2Result[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) (R, error) {
		v, err := f(ctx, p0, p1)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the CtxFunc2Result,
// if there is no error.
func (f CtxFunc2Result[R, P0, P1]) Map(fn func(R) R) CtxFunc2Result[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) (R, error) {
		v, err := f(ctx, p0, p1)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the CtxFunc2Result,
// if there is an error.
func (f CtxFunc2Result[R, P0, P1]) MapErr(fn func(error) error) CtxFunc2Result[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) (R, error) {
		v, err := f(ctx, p0, p1)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func2Value that will return the provided value if the
// CtxFunc2Result returns an error.
func (f CtxFunc2Result[R, P0, P1]) Fallback(val R) CtxFunc2Value[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) R {
		v, err := f(ctx, p0, p1)
		if err != nil {
			return val
		}
		return v
	}
}


func (f CtxFunc2Result[R, P0, P1]) Curry2(p0 P0, p1 P1) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		return f(ctx, p0, p1)
	}
}
	

func (f CtxFunc2Result[R, P0, P1]) Curry1(p0 P0) CtxFunc1Result[R, P1] {
	return func(ctx context.Context, p1 P1) (R, error) {
		return f(ctx, p0, p1)
	}
}
	