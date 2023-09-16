package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFuncValue[R any] func(ctx context.Context) R

func (f CtxFuncValue[R]) Exec(ctx context.Context) R {
	return f(ctx)
}

func (f CtxFuncValue[R]) Timing(loggers ...func(d time.Duration)) CtxFuncValue[R] {
	return func(ctx context.Context) R {
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

func (f CtxFuncValue[R]) Fallible() CtxFuncResult[R] {
	return func(ctx context.Context) (R, error) {
		v := f(ctx)
		return v, nil
	}
}

func (f CtxFuncValue[R]) WithTimeout(timeout time.Duration) CtxFuncValue[R] {
	return func(ctx context.Context) R {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx)
	}
}

func (f CtxFuncValue[R]) WithDeadline(deadline time.Time) CtxFuncValue[R] {
	return func(ctx context.Context) R {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx)
	}
}

func (f CtxFuncValue[R]) WithCancel() CtxFuncValue[R] {
	return func(ctx context.Context) R {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx)
	}
}
