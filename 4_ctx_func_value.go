package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc4Value[R, P0, P1, P2, P3 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) R

func (f CtxFunc4Value[R, P0, P1, P2, P3]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) R {
	return f(ctx, p0, p1, p2, p3)
}

func (f CtxFunc4Value[R, P0, P1, P2, P3]) Timing(loggers ...func(d time.Duration)) CtxFunc4Value[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) R {
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
		return f(ctx, p0, p1, p2, p3)
	}
}

func (f CtxFunc4Value[R, P0, P1, P2, P3]) Fallible() CtxFunc4Result[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) (R, error) {
		v := f(ctx, p0, p1, p2, p3)
		return v, nil
	}
}

func (f CtxFunc4Value[R, P0, P1, P2, P3]) WithTimeout(timeout time.Duration) CtxFunc4Value[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) R {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3)
	}
}

func (f CtxFunc4Value[R, P0, P1, P2, P3]) WithDeadline(deadline time.Time) CtxFunc4Value[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) R {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3)
	}
}

func (f CtxFunc4Value[R, P0, P1, P2, P3]) WithCancel() CtxFunc4Value[R, P0, P1, P2, P3] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3) R {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3)
	}
}


func (f CtxFunc4Value[R, P0, P1, P2, P3]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFuncValue[R] {
	return func(ctx context.Context) R {
		return f(ctx, p0, p1, p2, p3)
	}
}
	

func (f CtxFunc4Value[R, P0, P1, P2, P3]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc1Value[R, P3] {
	return func(ctx context.Context, p3 P3) R {
		return f(ctx, p0, p1, p2, p3)
	}
}
	

func (f CtxFunc4Value[R, P0, P1, P2, P3]) Curry2(p0 P0, p1 P1) CtxFunc2Value[R, P2, P3] {
	return func(ctx context.Context, p2 P2, p3 P3) R {
		return f(ctx, p0, p1, p2, p3)
	}
}
	

func (f CtxFunc4Value[R, P0, P1, P2, P3]) Curry1(p0 P0) CtxFunc3Value[R, P1, P2, P3] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3) R {
		return f(ctx, p0, p1, p2, p3)
	}
}
	