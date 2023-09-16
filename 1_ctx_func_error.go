package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc1Error[P0 any] func(ctx context.Context, p0 P0) error

func (f CtxFunc1Error[P0]) Exec(ctx context.Context, p0 P0) error {
	return f(ctx, p0)
}

func (f CtxFunc1Error[P0]) Timing(loggers ...func(d time.Duration)) CtxFunc1Error[P0] {
	return func(ctx context.Context, p0 P0) error {
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

func (f CtxFunc1Error[P0]) WithTimeout(timeout time.Duration) CtxFunc1Error[P0] {
	return func(ctx context.Context, p0 P0) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0)
	}
}

func (f CtxFunc1Error[P0]) WithDeadline(deadline time.Time) CtxFunc1Error[P0] {
	return func(ctx context.Context, p0 P0) error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0)
	}
}

func (f CtxFunc1Error[P0]) WithCancel() CtxFunc1Error[P0] {
	return func(ctx context.Context, p0 P0) error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc1Error[P0]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc1Error[P0] {
	return func(ctx context.Context, p0 P0) error {
		var err error
		attempts := 1
		for {
			err = f(ctx, p0)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func1Value that will panic if the CtxFunc1Result returns an error.
func (f CtxFunc1Error[P0]) Must() CtxFunc1[P0] {
	return func(ctx context.Context, p0 P0) {
		err := f(ctx, p0)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a CtxFunc1Result that will wrap the error returned by the CtxFunc1Result
// with the provided message.
func (f CtxFunc1Error[P0]) OnErr(msg string) CtxFunc1Error[P0] {
	return func(ctx context.Context, p0 P0) error {
		err := f(ctx, p0)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f CtxFunc1Error[P0]) Curry1(p0 P0) CtxFuncError {
	return func(ctx context.Context) error {
		return f(ctx, p0)
	}
}
	