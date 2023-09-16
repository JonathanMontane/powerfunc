package powerfunc

import (
	"fmt"
	"time"
)

// Func10 is a function that takes 0 arguments and returns a value.
type Func10Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9 any] func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) T

// Exec executes the function.
// Convenience method to remove some of the confusion that
// `f.SomeMethod()()` can bring.
func (f Func10Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Exec(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) T {
	return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
}

// Timing returns a Func10 that will log the execution time of the Func10.
// If no loggers are provided, the default logger (fmt.Println) will be used.
func (f Func10Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Timing(loggers ...func(d time.Duration)) Func10Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) T {
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
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}

// Fallible transforms a Func10Value into a Func10Result.
// The returned Func10Result will never return an error.
// Useful when passing a Func10Value to a function that expects a Func10Result.
func (f Func10Value[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Fallible() Func10Result[T, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) (T, error) {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9), nil
	}
}


func (f Func10Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry10(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) FuncValue[R] {
	return func() R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry9(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8) Func1Value[R, P9] {
	return func(p9 P9) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry8(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7) Func2Value[R, P8, P9] {
	return func(p8 P8, p9 P9) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry7(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6) Func3Value[R, P7, P8, P9] {
	return func(p7 P7, p8 P8, p9 P9) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry6(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4, p5 P5) Func4Value[R, P6, P7, P8, P9] {
	return func(p6 P6, p7 P7, p8 P8, p9 P9) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry5(p0 P0, p1 P1, p2 P2, p3 P3, p4 P4) Func5Value[R, P5, P6, P7, P8, P9] {
	return func(p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry4(p0 P0, p1 P1, p2 P2, p3 P3) Func6Value[R, P4, P5, P6, P7, P8, P9] {
	return func(p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry3(p0 P0, p1 P1, p2 P2) Func7Value[R, P3, P4, P5, P6, P7, P8, P9] {
	return func(p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry2(p0 P0, p1 P1) Func8Value[R, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	

func (f Func10Value[R, P0, P1, P2, P3, P4, P5, P6, P7, P8, P9]) Curry1(p0 P0) Func9Value[R, P1, P2, P3, P4, P5, P6, P7, P8, P9] {
	return func(p1 P1, p2 P2, p3 P3, p4 P4, p5 P5, p6 P6, p7 P7, p8 P8, p9 P9) R {
		return f(p0, p1, p2, p3, p4, p5, p6, p7, p8, p9)
	}
}
	