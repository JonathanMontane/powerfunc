package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc1[P0 any] func(ctx context.Context, p0 P0)

func (f CtxFunc1[P0]) Exec(ctx context.Context, p0 P0) {
	f(ctx, p0)
}

func (f CtxFunc1[P0]) Timing(loggers ...func(d time.Duration)) CtxFunc1[P0] {
	return func(ctx context.Context, p0 P0) {
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
		f(ctx, p0)
	}
}

func (f CtxFunc1[P0]) Fallible() CtxFunc1Error[P0] {
	return func(ctx context.Context, p0 P0) error {
		f(ctx, p0)
		return nil
	}
}

func (f CtxFunc1[P0]) WithTimeout(timeout time.Duration) CtxFunc1[P0] {
	return func(ctx context.Context, p0 P0) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		f(ctx, p0)
	}
}

func (f CtxFunc1[P0]) WithDeadline(deadline time.Time) CtxFunc1[P0] {
	return func(ctx context.Context, p0 P0) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		f(ctx, p0)
	}
}

func (f CtxFunc1[P0]) WithCancel() CtxFunc1[P0] {
	return func(ctx context.Context, p0 P0) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		f(ctx, p0)
	}
}


func (f CtxFunc1[P0]) Curry1(p0 P0) CtxFunc {
	return func(ctx context.Context)  {
		f(ctx, p0)
	}
}
	