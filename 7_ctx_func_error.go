package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error

func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
	return f(ctx, p0, p1, p2, p3, p4, p5, p6)
}

func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) Timing(loggers ...func(d time.Duration)) CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
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

func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) WithTimeout(timeout time.Duration) CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}

func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) WithDeadline(deadline time.Time) CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}

func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) WithCancel() CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
		var err error
		attempts := 1
		for {
			err = f(ctx, p0, p1, p2, p3, p4, p5, p6)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func7Value that will panic if the CtxFunc7Result returns an error.
func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) Must() CtxFunc7[P0, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) {
		err := f(ctx, p0, p1, p2, p3, p4, p5, p6)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a CtxFunc7Result that will wrap the error returned by the CtxFunc7Result
// with the provided message.
func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) OnErr(msg string) CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
		err := f(ctx, p0, p1, p2, p3, p4, p5, p6)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) CtxFuncError {
	return func(ctx context.Context) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) CtxFunc1Error[P6] {
	return func(ctx context.Context, p6 P6) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFunc2Error[P5, P6] {
	return func(ctx context.Context, p5 P5, p6 P6) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc3Error[P4, P5, P6] {
	return func(ctx context.Context, p4 P4, p5 P5, p6 P6) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc4Error[P3, P4, P5, P6] {
	return func(ctx context.Context, p3 P3, p4 P4, p5 P5, p6 P6) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) Curry2(p0 P0, p1 P1) CtxFunc5Error[P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	

func (f CtxFunc7Error[P0, P1, P2, P3, P4, P5, P6]) Curry1(p0 P0) CtxFunc6Error[P1, P2, P3, P4, P5, P6] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6)
	}
}
	