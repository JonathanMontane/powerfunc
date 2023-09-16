package powerfunc

import (
	"context"
	"errors"
)

func RetryImmediately(nbAttempts int) func(int, error) bool {
	return func(attempts int, err error) bool {
		if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
			return false
		}
		return attempts < nbAttempts
	}
}
