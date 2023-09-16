package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc3[P0, P1, P2 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2)

func (f CtxFunc3[P0, P1, P2]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2) {
	f(ctx, p0, p1, p2)
}

func (f CtxFunc3[P0, P1, P2]) Timing(loggers ...func(d time.Duration)) CtxFunc3[P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) {
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
		f(ctx, p0, p1, p2)
	}
}

func (f CtxFunc3[P0, P1, P2]) Fallible() CtxFunc3Error[P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) error {
		f(ctx, p0, p1, p2)
		return nil
	}
}

func (f CtxFunc3[P0, P1, P2]) WithTimeout(timeout time.Duration) CtxFunc3[P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		f(ctx, p0, p1, p2)
	}
}

func (f CtxFunc3[P0, P1, P2]) WithDeadline(deadline time.Time) CtxFunc3[P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		f(ctx, p0, p1, p2)
	}
}

func (f CtxFunc3[P0, P1, P2]) WithCancel() CtxFunc3[P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		f(ctx, p0, p1, p2)
	}
}


func (f CtxFunc3[P0, P1, P2]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc {
	return func(ctx context.Context)  {
		f(ctx, p0, p1, p2)
	}
}
	

func (f CtxFunc3[P0, P1, P2]) Curry2(p0 P0, p1 P1) CtxFunc1[P2] {
	return func(ctx context.Context, p2 P2)  {
		f(ctx, p0, p1, p2)
	}
}
	

func (f CtxFunc3[P0, P1, P2]) Curry1(p0 P0) CtxFunc2[P1, P2] {
	return func(ctx context.Context, p1 P1, p2 P2)  {
		f(ctx, p0, p1, p2)
	}
}
	