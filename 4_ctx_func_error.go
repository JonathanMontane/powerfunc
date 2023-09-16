package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc4Error[P0, P1, P2, P3 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) error

func (f CtxFunc4Error[P0, P1, P2, P3]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) error {
	return f(ctx, p0, p1, p2, p3)
}

func (f CtxFunc4Error[P0, P1, P2, P3]) Timing(loggers ...func(d time.Duration)) CtxFunc4Error[P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) error {
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

func (f CtxFunc4Error[P0, P1, P2, P3]) WithTimeout(timeout time.Duration) CtxFunc4Error[P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3)
	}
}

func (f CtxFunc4Error[P0, P1, P2, P3]) WithDeadline(deadline time.Time) CtxFunc4Error[P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3)
	}
}

func (f CtxFunc4Error[P0, P1, P2, P3]) WithCancel() CtxFunc4Error[P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc4Error[P0, P1, P2, P3]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc4Error[P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) error {
		var err error
		attempts := 1
		for {
			err = f(ctx, p0, p1, p2, p3)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func4Value that will panic if the CtxFunc4Result returns an error.
func (f CtxFunc4Error[P0, P1, P2, P3]) Must() CtxFunc4[P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) {
		err := f(ctx, p0, p1, p2, p3)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a CtxFunc4Result that will wrap the error returned by the CtxFunc4Result
// with the provided message.
func (f CtxFunc4Error[P0, P1, P2, P3]) OnErr(msg string) CtxFunc4Error[P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) error {
		err := f(ctx, p0, p1, p2, p3)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f CtxFunc4Error[P0, P1, P2, P3]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFuncError {
	return func(ctx context.Context) error {
		return f(ctx, p0, p1, p2, p3)
	}
}
	

func (f CtxFunc4Error[P0, P1, P2, P3]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc1Error[P3] {
	return func(ctx context.Context, p3 P3) error {
		return f(ctx, p0, p1, p2, p3)
	}
}
	

func (f CtxFunc4Error[P0, P1, P2, P3]) Curry2(p0 P0, p1 P1) CtxFunc2Error[P2, P3] {
	return func(ctx context.Context, p2 P2, p3 P3) error {
		return f(ctx, p0, p1, p2, p3)
	}
}
	

func (f CtxFunc4Error[P0, P1, P2, P3]) Curry1(p0 P0) CtxFunc3Error[P1, P2, P3] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3) error {
		return f(ctx, p0, p1, p2, p3)
	}
}
	