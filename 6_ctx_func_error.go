package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc6Error[P0, P1, P2, P3, P4, P5 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error

func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
	return f(ctx, p0, p1, p2, p3, p4, p5)
}

func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) Timing(loggers ...func(d time.Duration)) CtxFunc6Error[P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
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
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}

func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) WithTimeout(timeout time.Duration) CtxFunc6Error[P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}

func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) WithDeadline(deadline time.Time) CtxFunc6Error[P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}

func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) WithCancel() CtxFunc6Error[P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc6Error[P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
		var err error
		attempts := 1
		for {
			err = f(ctx, p0, p1, p2, p3, p4, p5)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func6Value that will panic if the CtxFunc6Result returns an error.
func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) Must() CtxFunc6[P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) {
		err := f(ctx, p0, p1, p2, p3, p4, p5)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a CtxFunc6Result that will wrap the error returned by the CtxFunc6Result
// with the provided message.
func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) OnErr(msg string) CtxFunc6Error[P0, P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
		err := f(ctx, p0, p1, p2, p3, p4, p5)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) CtxFuncError {
	return func(ctx context.Context) error {
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFunc1Error[P5] {
	return func(ctx context.Context, p5 P5) error {
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc2Error[P4, P5] {
	return func(ctx context.Context, p4 P4, p5 P5) error {
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc3Error[P3, P4, P5] {
	return func(ctx context.Context, p3 P3, p4 P4, p5 P5) error {
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) Curry2(p0 P0, p1 P1) CtxFunc4Error[P2, P3, P4, P5] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4, p5 P5) error {
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	

func (f CtxFunc6Error[P0, P1, P2, P3, P4, P5]) Curry1(p0 P0) CtxFunc5Error[P1, P2, P3, P4, P5] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) error {
		return f(ctx, p0, p1, p2, p3, p4, p5)
	}
}
	