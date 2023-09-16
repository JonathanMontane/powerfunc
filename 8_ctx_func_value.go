package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) R

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) R {
	return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
}

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) Timing(loggers ...func(d time.Duration)) CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) R {
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
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) Fallible() CtxFunc8Result[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) (R, error) {
		v := f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
		return v, nil
	}
}

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) WithTimeout(timeout time.Duration) CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) R {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) WithDeadline(deadline time.Time) CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) R {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) WithCancel() CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) R {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}


func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) CtxFuncValue[R] {
	return func(ctx context.Context) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) CtxFunc1Value[R, P7] {
	return func(ctx context.Context, p7 P7) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) CtxFunc2Value[R, P6, P7] {
	return func(ctx context.Context, p6 P6, p7 P7) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFunc3Value[R, P5, P6, P7] {
	return func(ctx context.Context, p5 P5, p6 P6, p7 P7) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc4Value[R, P4, P5, P6, P7] {
	return func(ctx context.Context, p4 P4, p5 P5, p6 P6, p7 P7) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc5Value[R, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry2(p0 P0, p1 P1) CtxFunc6Value[R, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	

func (f CtxFunc8Value[R, P0, P1, P2, P3, P4, P5, P6, P7]) Curry1(p0 P0) CtxFunc7Value[R, P1, P2, P3, P4, P5, P6, P7] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) R {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7)
	}
}
	