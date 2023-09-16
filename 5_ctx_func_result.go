package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc5Result[R, P0, P1, P2, P3, P4 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (R, error)

func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (R, error) {
	return f(ctx, p0, p1, p2, p3, p4)
}

func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) Timing(loggers ...func(d time.Duration)) CtxFunc5Result[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (R, error) {
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
		return f(ctx, p0, p1, p2, p3, p4)
	}
}

func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) WithTimeout(timeout time.Duration) CtxFunc5Result[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (R, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4)
	}
}

func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) WithDeadline(deadline time.Time) CtxFunc5Result[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (R, error) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4)
	}
}

func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) WithCancel() CtxFunc5Result[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (R, error) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc5Result[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (R, error) {
		var v R
		var err error
		attempts := 1
		for {
			v, err = f(ctx, p0, p1, p2, p3, p4)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func5Value that will panic if the CtxFunc5Result returns an error.
func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) Must() CtxFunc5Value[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) R {
		v, err := f(ctx, p0, p1, p2, p3, p4)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a CtxFunc5Result that will wrap the error returned by the CtxFunc5Result
// with the provided message.
func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) OnErr(msg string) CtxFunc5Result[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (R, error) {
		v, err := f(ctx, p0, p1, p2, p3, p4)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the CtxFunc5Result,
// if there is no error.
func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) Map(fn func(R) R) CtxFunc5Result[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (R, error) {
		v, err := f(ctx, p0, p1, p2, p3, p4)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the CtxFunc5Result,
// if there is an error.
func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) MapErr(fn func(error) error) CtxFunc5Result[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (R, error) {
		v, err := f(ctx, p0, p1, p2, p3, p4)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func5Value that will return the provided value if the
// CtxFunc5Result returns an error.
func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) Fallback(val R) CtxFunc5Value[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) R {
		v, err := f(ctx, p0, p1, p2, p3, p4)
		if err != nil {
			return val
		}
		return v
	}
}


func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4)
	}
}
	

func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc1Result[R, P4] {
	return func(ctx context.Context, p4 P4) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4)
	}
}
	

func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc2Result[R, P3, P4] {
	return func(ctx context.Context, p3 P3, p4 P4) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4)
	}
}
	

func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) Curry2(p0 P0, p1 P1) CtxFunc3Result[R, P2, P3, P4] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4)
	}
}
	

func (f CtxFunc5Result[R, P0, P1, P2, P3, P4]) Curry1(p0 P0) CtxFunc4Result[R, P1, P2, P3, P4] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4)
	}
}
	