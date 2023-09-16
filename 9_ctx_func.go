package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8)

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) {
	f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Timing(loggers ...func(d time.Duration)) CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) {
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
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Fallible() CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
		return nil
	}
}

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) WithTimeout(timeout time.Duration) CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) WithDeadline(deadline time.Time) CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) WithCancel() CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}


func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry9(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) CtxFunc {
	return func(ctx context.Context)  {
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) CtxFunc1[P8] {
	return func(ctx context.Context, p8 P8)  {
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) CtxFunc2[P7, P8] {
	return func(ctx context.Context, p7 P7, p8 P8)  {
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) CtxFunc3[P6, P7, P8] {
	return func(ctx context.Context, p6 P6, p7 P7, p8 P8)  {
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFunc4[P5, P6, P7, P8] {
	return func(ctx context.Context, p5 P5, p6 P6, p7 P7, p8 P8)  {
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc5[P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8)  {
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc6[P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8)  {
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry2(p0 P0, p1 P1) CtxFunc7[P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8)  {
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry1(p0 P0) CtxFunc8[P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8)  {
		f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	