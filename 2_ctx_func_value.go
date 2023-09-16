package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc2Value[R, P0, P1 any] func(ctx context.Context, p0 P0, p1 P1) R

func (f CtxFunc2Value[R, P0, P1]) Exec(ctx context.Context, p0 P0, p1 P1) R {
	return f(ctx, p0, p1)
}

func (f CtxFunc2Value[R, P0, P1]) Timing(loggers ...func(d time.Duration)) CtxFunc2Value[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) R {
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
		return f(ctx, p0, p1)
	}
}

func (f CtxFunc2Value[R, P0, P1]) Fallible() CtxFunc2Result[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) (R, error) {
		v := f(ctx, p0, p1)
		return v, nil
	}
}

func (f CtxFunc2Value[R, P0, P1]) WithTimeout(timeout time.Duration) CtxFunc2Value[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) R {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1)
	}
}

func (f CtxFunc2Value[R, P0, P1]) WithDeadline(deadline time.Time) CtxFunc2Value[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) R {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1)
	}
}

func (f CtxFunc2Value[R, P0, P1]) WithCancel() CtxFunc2Value[R, P0, P1] {
	return func(ctx context.Context, p0 P0, p1 P1) R {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1)
	}
}


func (f CtxFunc2Value[R, P0, P1]) Curry2(p0 P0, p1 P1) CtxFuncValue[R] {
	return func(ctx context.Context) R {
		return f(ctx, p0, p1)
	}
}
	

func (f CtxFunc2Value[R, P0, P1]) Curry1(p0 P0) CtxFunc1Value[R, P1] {
	return func(ctx context.Context, p1 P1) R {
		return f(ctx, p0, p1)
	}
}
	