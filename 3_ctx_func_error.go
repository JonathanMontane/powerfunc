package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc3Error[P0, P1, P2 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2) error

func (f CtxFunc3Error[P0, P1, P2]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2) error {
	return f(ctx, p0, p1, p2)
}

func (f CtxFunc3Error[P0, P1, P2]) Timing(loggers ...func(d time.Duration)) CtxFunc3Error[P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) error {
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

func (f CtxFunc3Error[P0, P1, P2]) WithTimeout(timeout time.Duration) CtxFunc3Error[P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2)
	}
}

func (f CtxFunc3Error[P0, P1, P2]) WithDeadline(deadline time.Time) CtxFunc3Error[P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2)
	}
}

func (f CtxFunc3Error[P0, P1, P2]) WithCancel() CtxFunc3Error[P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc3Error[P0, P1, P2]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc3Error[P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) error {
		var err error
		attempts := 1
		for {
			err = f(ctx, p0, p1, p2)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func3Value that will panic if the CtxFunc3Result returns an error.
func (f CtxFunc3Error[P0, P1, P2]) Must() CtxFunc3[P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) {
		err := f(ctx, p0, p1, p2)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a CtxFunc3Result that will wrap the error returned by the CtxFunc3Result
// with the provided message.
func (f CtxFunc3Error[P0, P1, P2]) OnErr(msg string) CtxFunc3Error[P0, P1, P2] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2) error {
		err := f(ctx, p0, p1, p2)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f CtxFunc3Error[P0, P1, P2]) Curry3(p0 P0, p1 P1, p2 P2) CtxFuncError {
	return func(ctx context.Context) error {
		return f(ctx, p0, p1, p2)
	}
}
	

func (f CtxFunc3Error[P0, P1, P2]) Curry2(p0 P0, p1 P1) CtxFunc1Error[P2] {
	return func(ctx context.Context, p2 P2) error {
		return f(ctx, p0, p1, p2)
	}
}
	

func (f CtxFunc3Error[P0, P1, P2]) Curry1(p0 P0) CtxFunc2Error[P1, P2] {
	return func(ctx context.Context, p1 P1, p2 P2) error {
		return f(ctx, p0, p1, p2)
	}
}
	