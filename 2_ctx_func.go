package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc2[P0, P1 any] func(ctx context.Context, p0 P0, p1 P1)

func (f CtxFunc2[P0, P1]) Exec(ctx context.Context, p0 P0, p1 P1) {
	f(ctx, p0, p1)
}

func (f CtxFunc2[P0, P1]) Timing(loggers ...func(d time.Duration)) CtxFunc2[P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) {
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
		f(ctx, p0, p1)
	}
}

func (f CtxFunc2[P0, P1]) Fallible() CtxFunc2Error[P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) error {
		f(ctx, p0, p1)
		return nil
	}
}

func (f CtxFunc2[P0, P1]) WithTimeout(timeout time.Duration) CtxFunc2[P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		f(ctx, p0, p1)
	}
}

func (f CtxFunc2[P0, P1]) WithDeadline(deadline time.Time) CtxFunc2[P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		f(ctx, p0, p1)
	}
}

func (f CtxFunc2[P0, P1]) WithCancel() CtxFunc2[P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		f(ctx, p0, p1)
	}
}


func (f CtxFunc2[P0, P1]) Curry2(p0 P0, p1 P1) CtxFunc {
	return func(ctx context.Context)  {
		f(ctx, p0, p1)
	}
}
	

func (f CtxFunc2[P0, P1]) Curry1(p0 P0) CtxFunc1[P1] {
	return func(ctx context.Context, p1 P1)  {
		f(ctx, p0, p1)
	}
}
	