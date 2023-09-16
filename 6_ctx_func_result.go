package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc6Result[R, P0, P1, P2, P3, P4, P5 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error)

func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
	return f(ctx, p0, p1, p2, p3, p4, p5)
}

func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) Timing(loggers ...func(d time.Duration)) CtxFunc6Result[R, P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
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
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}

func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) WithTimeout(timeout time.Duration) CtxFunc6Result[R, P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}

func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) WithDeadline(deadline time.Time) CtxFunc6Result[R, P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}

func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) WithCancel() CtxFunc6Result[R, P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc6Result[R, P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
		var v R
		var err error
		attempts := 1
		for {
			v, err = f(ctx, p0, p1, p2, p3, p4, p5)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func6Value that will panic if the CtxFunc6Result returns an error.
func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) Must() CtxFunc6Value[R, P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) R {
		v, err := f(ctx, p0, p1, p2, p3, p4, p5)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a CtxFunc6Result that will wrap the error returned by the CtxFunc6Result
// with the provided message.
func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) OnErr(msg string) CtxFunc6Result[R, P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
		v, err := f(ctx, p0, p1, p2, p3, p4, p5)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the CtxFunc6Result,
// if there is no error.
func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) Map(fn func(R) R) CtxFunc6Result[R, P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
		v, err := f(ctx, p0, p1, p2, p3, p4, p5)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the CtxFunc6Result,
// if there is an error.
func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) MapErr(fn func(error) error) CtxFunc6Result[R, P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
		v, err := f(ctx, p0, p1, p2, p3, p4, p5)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func6Value that will return the provided value if the
// CtxFunc6Result returns an error.
func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) Fallback(val R) CtxFunc6Value[R, P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) R {
		v, err := f(ctx, p0, p1, p2, p3, p4, p5)
		if err != nil {
			return val
		}
		return v
	}
}


func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFunc1Result[R, P5] {
	return func(ctx context.Context, p5 P5) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc2Result[R, P4, P5] {
	return func(ctx context.Context, p4 P4, p5 P5) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc3Result[R, P3, P4, P5] {
	return func(ctx context.Context, p3 P3, p4 P4, p5 P5) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) Curry2(p0 P0, p1 P1) CtxFunc4Result[R, P2, P3, P4, P5] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6Result[R, P0, P1, P2, P3, P4, P5]) Curry1(p0 P0) CtxFunc5Result[R, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	