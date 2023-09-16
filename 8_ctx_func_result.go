package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error)

func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
	return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
}

func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Timing(loggers ...func(d time.Duration)) CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
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
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}

func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) WithTimeout(timeout time.Duration) CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}

func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) WithDeadline(deadline time.Time) CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}

func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) WithCancel() CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
		var v R
		var err error
		attempts := 1
		for {
			v, err = f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return v, err
	}
}

// Must returns a Func8Value that will panic if the CtxFunc8Result returns an error.
func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Must() CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) R {
		v, err := f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
		if err != nil {
			panic(err)
		}
		return v
	}
}

// OnErr returns a CtxFunc8Result that will wrap the error returned by the CtxFunc8Result
// with the provided message.
func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) OnErr(msg string) CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
		v, err := f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
		if err != nil {
			return v, fmt.Errorf("%s: %w", msg, err)
		}
		return v, nil
	}
}

// Map applies the provided function to the value returned by the CtxFunc8Result,
// if there is no error.
func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Map(fn func(R) R) CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
		v, err := f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
		if err != nil {
			return v, err
		}
		return fn(v), nil
	}
}

// MapErr applies the provided function to the error returned by the CtxFunc8Result,
// if there is an error.
func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) MapErr(fn func(error) error) CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
		v, err := f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
		if err != nil {
			return v, fn(err)
		}
		return v, nil
	}
}

// Fallback returns a Func8Value that will return the provided value if the
// CtxFunc8Result returns an error.
func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Fallback(val R) CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) R {
		v, err := f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
		if err != nil {
			return val
		}
		return v
	}
}


func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) CtxFunc1Result[R, P7] {
	return func(ctx context.Context, p7 P7) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) CtxFunc2Result[R, P6, P7] {
	return func(ctx context.Context, p6 P6, p7 P7) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFunc3Result[R, P5, P6, P7] {
	return func(ctx context.Context, p5 P5, p6 P6, p7 P7) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc4Result[R, P4, P5, P6, P7] {
	return func(ctx context.Context, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc5Result[R, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry2(p0 P0, p1 P1) CtxFunc6Result[R, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry1(p0 P0) CtxFunc7Result[R, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	