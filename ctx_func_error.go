package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFuncError func(ctx context.Context) error

func (f CtxFuncError) Exec(ctx context.Context) error {
	return f(ctx)
}

func (f CtxFuncError) Timing(loggers ...func(d time.Duration)) CtxFuncError {
	return func(ctx context.Context) error {
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

func (f CtxFuncError) WithTimeout(timeout time.Duration) CtxFuncError {
	return func(ctx context.Context) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx)
	}
}

func (f CtxFuncError) WithDeadline(deadline time.Time) CtxFuncError {
	return func(ctx context.Context) error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx)
	}
}

func (f CtxFuncError) WithCancel() CtxFuncError {
	return func(ctx context.Context) error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFuncError) Retry(tryAgain func(attempts int, err error) bool) CtxFuncError {
	return func(ctx context.Context) error {
		var err error
		attempts := 1
		for {
			err = f(ctx)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a FuncValue that will panic if the CtxFuncResult returns an error.
func (f CtxFuncError) Must() CtxFunc {
	return func(ctx context.Context) {
		err := f(ctx)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a CtxFuncResult that will wrap the error returned by the CtxFuncResult
// with the provided message.
func (f CtxFuncError) OnErr(msg string) CtxFuncError {
	return func(ctx context.Context) error {
		err := f(ctx)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}
