package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc1Value[R, P0 any] func(ctx context.Context, p0 P0) R

func (f CtxFunc1Value[R, P0]) Exec(ctx context.Context, p0 P0) R {
	return f(ctx, p0)
}

func (f CtxFunc1Value[R, P0]) Timing(loggers ...func(d time.Duration)) CtxFunc1Value[R, P0] {
	return func(ctx context.Context, p0 P0) R {
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

func (f CtxFunc1Value[R, P0]) Fallible() CtxFunc1Result[R, P0] {
	return func(ctx context.Context, p0 P0) (R, error) {
		v := f(ctx, p0)
		return v, nil
	}
}

func (f CtxFunc1Value[R, P0]) WithTimeout(timeout time.Duration) CtxFunc1Value[R, P0] {
	return func(ctx context.Context, p0 P0) R {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0)
	}
}

func (f CtxFunc1Value[R, P0]) WithDeadline(deadline time.Time) CtxFunc1Value[R, P0] {
	return func(ctx context.Context, p0 P0) R {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0)
	}
}

func (f CtxFunc1Value[R, P0]) WithCancel() CtxFunc1Value[R, P0] {
	return func(ctx context.Context, p0 P0) R {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0)
	}
}


func (f CtxFunc1Value[R, P0]) Curry1(p0 P0) CtxFuncValue[R] {
	return func(ctx context.Context) R {
		return f(ctx, p0)
	}
}
	