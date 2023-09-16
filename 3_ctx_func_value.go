package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc3Value[R, P0, P1, P2 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2) R

func (f CtxFunc3Value[R, P0, P1, P2]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2) R {
	return f(ctx, p0, p1, p2)
}

func (f CtxFunc3Value[R, P0, P1, P2]) Timing(loggers ...func(d time.Duration)) CtxFunc3Value[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) R {
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
		return f(ctx, p0, p1, p2)
	}
}

func (f CtxFunc3Value[R, P0, P1, P2]) Fallible() CtxFunc3Result[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) (R, error) {
		v := f(ctx, p0, p1, p2)
		return v, nil
	}
}

func (f CtxFunc3Value[R, P0, P1, P2]) WithTimeout(timeout time.Duration) CtxFunc3Value[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) R {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2)
	}
}

func (f CtxFunc3Value[R, P0, P1, P2]) WithDeadline(deadline time.Time) CtxFunc3Value[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) R {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2)
	}
}

func (f CtxFunc3Value[R, P0, P1, P2]) WithCancel() CtxFunc3Value[R, P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) R {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2)
	}
}


func (f CtxFunc3Value[R, P0, P1, P2]) Curry3(p0 P0, p1 P1, p2 P2) CtxFuncValue[R] {
	return func(ctx context.Context) R {
		return f(ctx, p0, p1, p2)
	}
}
	

func (f CtxFunc3Value[R, P0, P1, P2]) Curry2(p0 P0, p1 P1) CtxFunc1Value[R, P2] {
	return func(ctx context.Context, p2 P2) R {
		return f(ctx, p0, p1, p2)
	}
}
	

func (f CtxFunc3Value[R, P0, P1, P2]) Curry1(p0 P0) CtxFunc2Value[R, P1, P2] {
	return func(ctx context.Context, p1 P1, p2 P2) R {
		return f(ctx, p0, p1, p2)
	}
}
	