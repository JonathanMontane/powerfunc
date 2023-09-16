package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc func(ctx context.Context)

func (f CtxFunc) Exec(ctx context.Context) {
	f(ctx)
}

func (f CtxFunc) Timing(loggers ...func(d time.Duration)) CtxFunc {
	return func(ctx context.Context) {
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
		f(ctx)
	}
}

func (f CtxFunc) Fallible() CtxFuncError {
	return func(ctx context.Context) error {
		f(ctx)
		return nil
	}
}

func (f CtxFunc) WithTimeout(timeout time.Duration) CtxFunc {
	return func(ctx context.Context) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		f(ctx)
	}
}

func (f CtxFunc) WithDeadline(deadline time.Time) CtxFunc {
	return func(ctx context.Context) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		f(ctx)
	}
}

func (f CtxFunc) WithCancel() CtxFunc {
	return func(ctx context.Context) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		f(ctx)
	}
}
