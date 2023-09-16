package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) R

func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) R {
	return f(ctx, p0, p1, p2, p3, p4, p5, p6)
}

func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) Timing(loggers ...func(d time.Duration)) CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) R {
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
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}

func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) Fallible() CtxFunc7Result[R, P0, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) (R, error) {
		v := f(ctx, p0, p1, p2, p3, p4, p5, p6)
		return v, nil
	}
}

func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) WithTimeout(timeout time.Duration) CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) R {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}

func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) WithDeadline(deadline time.Time) CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) R {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}

func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) WithCancel() CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) R {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}


func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) CtxFuncValue[R] {
	return func(ctx context.Context) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) CtxFunc1Value[R, P6] {
	return func(ctx context.Context, p6 P6) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFunc2Value[R, P5, P6] {
	return func(ctx context.Context, p5 P5, p6 P6) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc3Value[R, P4, P5, P6] {
	return func(ctx context.Context, p4 P4, p5 P5, p6 P6) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc4Value[R, P3, P4, P5, P6] {
	return func(ctx context.Context, p3 P3, p4 P4, p5 P5, p6 P6) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry2(p0 P0, p1 P1) CtxFunc5Value[R, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f CtxFunc7Value[R, P0, P1, P2, P3, P4, P5, P6]) Curry1(p0 P0) CtxFunc6Value[R, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	