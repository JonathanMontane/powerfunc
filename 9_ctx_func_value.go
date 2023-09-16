package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
	return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Timing(loggers ...func(d time.Duration)) CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
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
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Fallible() CtxFunc9Result[R, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) (R, error) {
		v := f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
		return v, nil
	}
}

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) WithTimeout(timeout time.Duration) CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) WithDeadline(deadline time.Time) CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) WithCancel() CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}


func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry9(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) CtxFuncValue[R] {
	return func(ctx context.Context) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) CtxFunc1Value[R, P8] {
	return func(ctx context.Context, p8 P8) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) CtxFunc2Value[R, P7, P8] {
	return func(ctx context.Context, p7 P7, p8 P8) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) CtxFunc3Value[R, P6, P7, P8] {
	return func(ctx context.Context, p6 P6, p7 P7, p8 P8) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFunc4Value[R, P5, P6, P7, P8] {
	return func(ctx context.Context, p5 P5, p6 P6, p7 P7, p8 P8) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc5Value[R, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc6Value[R, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry2(p0 P0, p1 P1) CtxFunc7Value[R, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry1(p0 P0) CtxFunc8Value[R, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	