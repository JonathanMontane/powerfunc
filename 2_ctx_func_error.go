package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc2Error[P0, P1 any] func(ctx context.Context, p0 P0, p1 P1) error

func (f CtxFunc2Error[P0, P1]) Exec(ctx context.Context, p0 P0, p1 P1) error {
	return f(ctx, p0, p1)
}

func (f CtxFunc2Error[P0, P1]) Timing(loggers ...func(d time.Duration)) CtxFunc2Error[P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) error {
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

func (f CtxFunc2Error[P0, P1]) WithTimeout(timeout time.Duration) CtxFunc2Error[P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1)
	}
}

func (f CtxFunc2Error[P0, P1]) WithDeadline(deadline time.Time) CtxFunc2Error[P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1)
	}
}

func (f CtxFunc2Error[P0, P1]) WithCancel() CtxFunc2Error[P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc2Error[P0, P1]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc2Error[P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) error {
		var err error
		attempts := 1
		for {
			err = f(ctx, p0, p1)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func2Value that will panic if the CtxFunc2Result returns an error.
func (f CtxFunc2Error[P0, P1]) Must() CtxFunc2[P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) {
		err := f(ctx, p0, p1)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a CtxFunc2Result that will wrap the error returned by the CtxFunc2Result
// with the provided message.
func (f CtxFunc2Error[P0, P1]) OnErr(msg string) CtxFunc2Error[P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) error {
		err := f(ctx, p0, p1)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f CtxFunc2Error[P0, P1]) Curry2(p0 P0, p1 P1) CtxFuncError {
	return func(ctx context.Context) error {
		return f(ctx, p0, p1)
	}
}
	

func (f CtxFunc2Error[P0, P1]) Curry1(p0 P0) CtxFunc1Error[P1] {
	return func(ctx context.Context, p1 P1) error {
		return f(ctx, p0, p1)
	}
}
	