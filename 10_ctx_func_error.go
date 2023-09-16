package powerfunc

import (
	"context"
	"fmt"
	"time"
)

type CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9 any] func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Exec(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
	return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
}

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Timing(loggers ...func(d time.Duration)) CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
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
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) WithTimeout(timeout time.Duration) CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		ctx, cancel := context.WithTimeout(ctx, timeout)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) WithDeadline(deadline time.Time) CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		ctx, cancel := context.WithDeadline(ctx, deadline)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) WithCancel() CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		ctx, cancel := context.WithCancel(ctx)
		defer cancel()
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}

// Retry returns a Function that will retry the Function until it returns
// a nil error or the tryAgain function returns false.
func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Retry(tryAgain func(attempts int, err error) bool) CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		var err error
		attempts := 1
		for {
			err = f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
			if err != nil && tryAgain(attempts, err) {
				break
			}
			attempts++
		}
		return err
	}
}

// Must returns a Func10Value that will panic if the CtxFunc10Result returns an error.
func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Must() CtxFunc10[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) {
		err := f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
		if err != nil {
			panic(err)
		}
	}
}

// OnErr returns a CtxFunc10Result that will wrap the error returned by the CtxFunc10Result
// with the provided message.
func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) OnErr(msg string) CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(ctx context.Context, p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		err := f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
		if err != nil {
			return fmt.Errorf("%s: %w", msg, err)
		}
		return nil
	}
}


func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry10(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) CtxFuncError {
	return func(ctx context.Context) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry9(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) CtxFunc1Error[P9] {
	return func(ctx context.Context, p9 P9) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) CtxFunc2Error[P8, P9] {
	return func(ctx context.Context, p8 P8, p9 P9) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) CtxFunc3Error[P7, P8, P9] {
	return func(ctx context.Context, p7 P7, p8 P8, p9 P9) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) CtxFunc4Error[P6, P7, P8, P9] {
	return func(ctx context.Context, p6 P6, p7 P7, p8 P8, p9 P9) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) CtxFunc5Error[P5, P6, P7, P8, P9] {
	return func(ctx context.Context, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) CtxFunc6Error[P4, P5, P6, P7, P8, P9] {
	return func(ctx context.Context, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry3(p0 P0, p1 P1, p2 P2) CtxFunc7Error[P3, P4, P5, P6, P7, P8, P9] {
	return func(ctx context.Context, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry2(p0 P0, p1 P1) CtxFunc8Error[P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(ctx context.Context, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f CtxFunc10Error[P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry1(p0 P0) CtxFunc9Error[P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(ctx context.Context, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) error {
		return f(ctx, p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	