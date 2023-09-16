package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc4Result[R, P0, P1, P2, P3 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) (R, error)

func (f CtxFunc4Result[R, P0, P1, P2, P3]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) (R, error) {
	return f(ctx, p0, p1, p2, p3)
}

func (f CtxFunc4Result[R, P0, P1, P2, P3]) Timing(loggers ...func(d time.Duration)) CtxFunc4Result[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) (R, error) {
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
		return f(ctx, p0, p1, p2, p3)
	}
}

func (f CtxFunc4Result[R, P0, P1, P2, P3]) WithTimeout(timeout time.Duration) CtxFunc4Result[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) (R, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3)
	}
}

func (f CtxFunc4Result[R, P0, P1, P2, P3]) WithDeadline(deadline time.Time) CtxFunc4Result[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) (R, error) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3)
	}
}

func (f CtxFunc4Result[R, P0, P1, P2, P3]) WithCancel() CtxFunc4Result[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) (R, error) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc4Result[R, P0, P1, P2, P3]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc4Result[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) (R, error) {
		var v R
		var err error
		attempts := 1
		for {
			v, err = f(ctx, p0, p1, p2, p3)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func4Value that will panic if the CtxFunc4Result returns an error.
func (f CtxFunc4Result[R, P0, P1, P2, P3]) Must() CtxFunc4Value[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) R {
		v, err := f(ctx, p0, p1, p2, p3)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a CtxFunc4Result that will wrap the error returned by the CtxFunc4Result
// with the provided message.
func (f CtxFunc4Result[R, P0, P1, P2, P3]) OnErr(msg string) CtxFunc4Result[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) (R, error) {
		v, err := f(ctx, p0, p1, p2, p3)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the CtxFunc4Result,
// if there is no error.
func (f CtxFunc4Result[R, P0, P1, P2, P3]) Map(fn func(R) R) CtxFunc4Result[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) (R, error) {
		v, err := f(ctx, p0, p1, p2, p3)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the CtxFunc4Result,
// if there is an error.
func (f CtxFunc4Result[R, P0, P1, P2, P3]) MapErr(fn func(error) error) CtxFunc4Result[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) (R, error) {
		v, err := f(ctx, p0, p1, p2, p3)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func4Value that will return the provided value if the
// CtxFunc4Result returns an error.
func (f CtxFunc4Result[R, P0, P1, P2, P3]) Fallback(val R) CtxFunc4Value[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) R {
		v, err := f(ctx, p0, p1, p2, p3)
		if err != nil {
			return val
		}
		return v
	}
}


func (f CtxFunc4Result[R, P0, P1, P2, P3]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		return f(ctx, p0, p1, p2, p3)
	}
}
	

func (f CtxFunc4Result[R, P0, P1, P2, P3]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc1Result[R, P3] {
	return func(ctx context.Context, p3 P3) (R, error) {
		return f(ctx, p0, p1, p2, p3)
	}
}
	

func (f CtxFunc4Result[R, P0, P1, P2, P3]) Curry2(p0 P0, p1 P1) CtxFunc2Result[R, P2, P3] {
	return func(ctx context.Context, p2 P2, p3 P3) (R, error) {
		return f(ctx, p0, p1, p2, p3)
	}
}
	

func (f CtxFunc4Result[R, P0, P1, P2, P3]) Curry1(p0 P0) CtxFunc3Result[R, P1, P2, P3] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3) (R, error) {
		return f(ctx, p0, p1, p2, p3)
	}
}
	