package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc5Value[R, P0, P1, P2, P3, P4 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) R

func (f CtxFunc5Value[R, P0, P1, P2, P3, P4]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) R {
	return f(ctx, p0, p1, p2, p3, p4)
}

func (f CtxFunc5Value[R, P0, P1, P2, P3, P4]) Timing(loggers ...func(d time.Duration)) CtxFunc5Value[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) R {
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
		return f(ctx, p0, p1, p2, p3, p4)
	}
}

func (f CtxFunc5Value[R, P0, P1, P2, P3, P4]) Fallible() CtxFunc5Result[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) (R, error) {
		v := f(ctx, p0, p1, p2, p3, p4)
		return v, nil
	}
}

func (f CtxFunc5Value[R, P0, P1, P2, P3, P4]) WithTimeout(timeout time.Duration) CtxFunc5Value[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) R {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4)
	}
}

func (f CtxFunc5Value[R, P0, P1, P2, P3, P4]) WithDeadline(deadline time.Time) CtxFunc5Value[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) R {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4)
	}
}

func (f CtxFunc5Value[R, P0, P1, P2, P3, P4]) WithCancel() CtxFunc5Value[R, P0, P1, P2, P3, P4] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) R {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4)
	}
}


func (f CtxFunc5Value[R, P0, P1, P2, P3, P4]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFuncValue[R] {
	return func(ctx context.Context) R {
		return f(ctx, p0, p1, p2, p3, p4)
	}
}
	

func (f CtxFunc5Value[R, P0, P1, P2, P3, P4]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc1Value[R, P4] {
	return func(ctx context.Context, p4 P4) R {
		return f(ctx, p0, p1, p2, p3, p4)
	}
}
	

func (f CtxFunc5Value[R, P0, P1, P2, P3, P4]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc2Value[R, P3, P4] {
	return func(ctx context.Context, p3 P3, p4 P4) R {
		return f(ctx, p0, p1, p2, p3, p4)
	}
}
	

func (f CtxFunc5Value[R, P0, P1, P2, P3, P4]) Curry2(p0 P0, p1 P1) CtxFunc3Value[R, P2, P3, P4] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4) R {
		return f(ctx, p0, p1, p2, p3, p4)
	}
}
	

func (f CtxFunc5Value[R, P0, P1, P2, P3, P4]) Curry1(p0 P0) CtxFunc4Value[R, P1, P2, P3, P4] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4) R {
		return f(ctx, p0, p1, p2, p3, p4)
	}
}
	