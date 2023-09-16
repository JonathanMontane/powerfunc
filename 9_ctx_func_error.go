package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
	return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
}

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Timing(loggers ...func(d time.Duration)) CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
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

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) WithTimeout(timeout time.Duration) CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) WithDeadline(deadline time.Time) CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) WithCancel() CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		var err error
		attempts := 1
		for {
			err = f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func9Value that will panic if the CtxFunc9Result returns an error.
func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Must() CtxFunc9[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) {
		err := f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a CtxFunc9Result that will wrap the error returned by the CtxFunc9Result
// with the provided message.
func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) OnErr(msg string) CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		err := f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry9(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) CtxFuncError {
	return func(ctx context.Context) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) CtxFunc1Error[P8] {
	return func(ctx context.Context, p8 P8) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) CtxFunc2Error[P7, P8] {
	return func(ctx context.Context, p7 P7, p8 P8) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) CtxFunc3Error[P6, P7, P8] {
	return func(ctx context.Context, p6 P6, p7 P7, p8 P8) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFunc4Error[P5, P6, P7, P8] {
	return func(ctx context.Context, p5 P5, p6 P6, p7 P7, p8 P8) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc5Error[P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc6Error[P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry2(p0 P0, p1 P1) CtxFunc7Error[P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	

func (f CtxFunc9Error[P0, P1, P2, P3, P4, P5, P6, P7, P8]) Curry1(p0 P0) CtxFunc8Error[P1, P2, P3, P4, P5, P6, P7, P8] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8)
	}
}
	