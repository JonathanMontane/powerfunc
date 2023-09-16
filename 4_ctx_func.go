package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc4[P0, P1, P2, P3 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3)

func (f CtxFunc4[P0, P1, P2, P3]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) {
	f(ctx, p0, p1, p2, p3)
}

func (f CtxFunc4[P0, P1, P2, P3]) Timing(loggers ...func(d time.Duration)) CtxFunc4[P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) {
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
		f(ctx, p0, p1, p2, p3)
	}
}

func (f CtxFunc4[P0, P1, P2, P3]) Fallible() CtxFunc4Error[P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) error {
		f(ctx, p0, p1, p2, p3)
		return nil
	}
}

func (f CtxFunc4[P0, P1, P2, P3]) WithTimeout(timeout time.Duration) CtxFunc4[P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		f(ctx, p0, p1, p2, p3)
	}
}

func (f CtxFunc4[P0, P1, P2, P3]) WithDeadline(deadline time.Time) CtxFunc4[P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		f(ctx, p0, p1, p2, p3)
	}
}

func (f CtxFunc4[P0, P1, P2, P3]) WithCancel() CtxFunc4[P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		f(ctx, p0, p1, p2, p3)
	}
}


func (f CtxFunc4[P0, P1, P2, P3]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc {
	return func(ctx context.Context)  {
		f(ctx, p0, p1, p2, p3)
	}
}
	

func (f CtxFunc4[P0, P1, P2, P3]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc1[P3] {
	return func(ctx context.Context, p3 P3)  {
		f(ctx, p0, p1, p2, p3)
	}
}
	

func (f CtxFunc4[P0, P1, P2, P3]) Curry2(p0 P0, p1 P1) CtxFunc2[P2, P3] {
	return func(ctx context.Context, p2 P2, p3 P3)  {
		f(ctx, p0, p1, p2, p3)
	}
}
	

func (f CtxFunc4[P0, P1, P2, P3]) Curry1(p0 P0) CtxFunc3[P1, P2, P3] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3)  {
		f(ctx, p0, p1, p2, p3)
	}
}
	