package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc5[P0, P1, P2, P3, P4 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4)

func (f CtxFunc5[P0, P1, P2, P3, P4]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) {
	f(ctx, p0, p1, p2, p3, p4)
}

func (f CtxFunc5[P0, P1, P2, P3, P4]) Timing(loggers ...func(d time.Duration)) CtxFunc5[P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) {
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
		f(ctx, p0, p1, p2, p3, p4)
	}
}

func (f CtxFunc5[P0, P1, P2, P3, P4]) Fallible() CtxFunc5Error[P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) error {
		f(ctx, p0, p1, p2, p3, p4)
		return nil
	}
}

func (f CtxFunc5[P0, P1, P2, P3, P4]) WithTimeout(timeout time.Duration) CtxFunc5[P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		f(ctx, p0, p1, p2, p3, p4)
	}
}

func (f CtxFunc5[P0, P1, P2, P3, P4]) WithDeadline(deadline time.Time) CtxFunc5[P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		f(ctx, p0, p1, p2, p3, p4)
	}
}

func (f CtxFunc5[P0, P1, P2, P3, P4]) WithCancel() CtxFunc5[P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		f(ctx, p0, p1, p2, p3, p4)
	}
}


func (f CtxFunc5[P0, P1, P2, P3, P4]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFunc {
	return func(ctx context.Context)  {
		f(ctx, p0, p1, p2, p3, p4)
	}
}
	

func (f CtxFunc5[P0, P1, P2, P3, P4]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc1[P4] {
	return func(ctx context.Context, p4 P4)  {
		f(ctx, p0, p1, p2, p3, p4)
	}
}
	

func (f CtxFunc5[P0, P1, P2, P3, P4]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc2[P3, P4] {
	return func(ctx context.Context, p3 P3, p4 P4)  {
		f(ctx, p0, p1, p2, p3, p4)
	}
}
	

func (f CtxFunc5[P0, P1, P2, P3, P4]) Curry2(p0 P0, p1 P1) CtxFunc3[P2, P3, P4] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4)  {
		f(ctx, p0, p1, p2, p3, p4)
	}
}
	

func (f CtxFunc5[P0, P1, P2, P3, P4]) Curry1(p0 P0) CtxFunc4[P1, P2, P3, P4] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4)  {
		f(ctx, p0, p1, p2, p3, p4)
	}
}
	