package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc6[P0, P1, P2, P3, P4, P5 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5)

func (f CtxFunc6[P0, P1, P2, P3, P4, P5]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) {
	f(ctx, p0, p1, p2, p3, p4, p5)
}

func (f CtxFunc6[P0, P1, P2, P3, P4, P5]) Timing(loggers ...func(d time.Duration)) CtxFunc6[P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) {
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
		f(ctx, p0, p1, p2, p3, p4, p5)
	}
}

func (f CtxFunc6[P0, P1, P2, P3, P4, P5]) Fallible() CtxFunc6Error[P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
		f(ctx, p0, p1, p2, p3, p4, p5)
		return nil
	}
}

func (f CtxFunc6[P0, P1, P2, P3, P4, P5]) WithTimeout(timeout time.Duration) CtxFunc6[P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		f(ctx, p0, p1, p2, p3, p4, p5)
	}
}

func (f CtxFunc6[P0, P1, P2, P3, P4, P5]) WithDeadline(deadline time.Time) CtxFunc6[P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		f(ctx, p0, p1, p2, p3, p4, p5)
	}
}

func (f CtxFunc6[P0, P1, P2, P3, P4, P5]) WithCancel() CtxFunc6[P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		f(ctx, p0, p1, p2, p3, p4, p5)
	}
}


func (f CtxFunc6[P0, P1, P2, P3, P4, P5]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) CtxFunc {
	return func(ctx context.Context)  {
		f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6[P0, P1, P2, P3, P4, P5]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFunc1[P5] {
	return func(ctx context.Context, p5 P5)  {
		f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6[P0, P1, P2, P3, P4, P5]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc2[P4, P5] {
	return func(ctx context.Context, p4 P4, p5 P5)  {
		f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6[P0, P1, P2, P3, P4, P5]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc3[P3, P4, P5] {
	return func(ctx context.Context, p3 P3, p4 P4, p5 P5)  {
		f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6[P0, P1, P2, P3, P4, P5]) Curry2(p0 P0, p1 P1) CtxFunc4[P2, P3, P4, P5] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4, p5 P5)  {
		f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6[P0, P1, P2, P3, P4, P5]) Curry1(p0 P0) CtxFunc5[P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5)  {
		f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	