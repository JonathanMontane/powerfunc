package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc3Result[R, P0, P1, P2 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2) (R, error)

func (f CtxFunc3Result[R, P0, P1, P2]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2) (R, error) {
	return f(ctx, p0, p1, p2)
}

func (f CtxFunc3Result[R, P0, P1, P2]) Timing(loggers ...func(d time.Duration)) CtxFunc3Result[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) (R, error) {
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
		return f(ctx, p0, p1, p2)
	}
}

func (f CtxFunc3Result[R, P0, P1, P2]) WithTimeout(timeout time.Duration) CtxFunc3Result[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) (R, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2)
	}
}

func (f CtxFunc3Result[R, P0, P1, P2]) WithDeadline(deadline time.Time) CtxFunc3Result[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) (R, error) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2)
	}
}

func (f CtxFunc3Result[R, P0, P1, P2]) WithCancel() CtxFunc3Result[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) (R, error) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc3Result[R, P0, P1, P2]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc3Result[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) (R, error) {
		var v R
		var err error
		attempts := 1
		for {
			v, err = f(ctx, p0, p1, p2)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func3Value that will panic if the CtxFunc3Result returns an error.
func (f CtxFunc3Result[R, P0, P1, P2]) Must() CtxFunc3Value[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) R {
		v, err := f(ctx, p0, p1, p2)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a CtxFunc3Result that will wrap the error returned by the CtxFunc3Result
// with the provided message.
func (f CtxFunc3Result[R, P0, P1, P2]) OnErr(msg string) CtxFunc3Result[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) (R, error) {
		v, err := f(ctx, p0, p1, p2)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the CtxFunc3Result,
// if there is no error.
func (f CtxFunc3Result[R, P0, P1, P2]) Map(fn func(R) R) CtxFunc3Result[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) (R, error) {
		v, err := f(ctx, p0, p1, p2)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the CtxFunc3Result,
// if there is an error.
func (f CtxFunc3Result[R, P0, P1, P2]) MapErr(fn func(error) error) CtxFunc3Result[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) (R, error) {
		v, err := f(ctx, p0, p1, p2)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func3Value that will return the provided value if the
// CtxFunc3Result returns an error.
func (f CtxFunc3Result[R, P0, P1, P2]) Fallback(val R) CtxFunc3Value[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) R {
		v, err := f(ctx, p0, p1, p2)
		if err != nil {
			return val
		}
		return v
	}
}


func (f CtxFunc3Result[R, P0, P1, P2]) Curry3(p0 P0, p1 P1, p2 P2) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		return f(ctx, p0, p1, p2)
	}
}
	

func (f CtxFunc3Result[R, P0, P1, P2]) Curry2(p0 P0, p1 P1) CtxFunc1Result[R, P2] {
	return func(ctx context.Context, p2 P2) (R, error) {
		return f(ctx, p0, p1, p2)
	}
}
	

func (f CtxFunc3Result[R, P0, P1, P2]) Curry1(p0 P0) CtxFunc2Result[R, P1, P2] {
	return func(ctx context.Context, p1 P1, p2 P2) (R, error) {
		return f(ctx, p0, p1, p2)
	}
}
	