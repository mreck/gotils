package gotils

import (
	"time"
)

// Try reruns the function until it succeeds or runs out of attempts
func Try(noAttempts int, waitDuration time.Duration, f func(attempt int) error) error {
	var err error

	for a := range noAttempts {
		err = f(a)
		if err == nil {
			break
		}
		time.Sleep(waitDuration)
	}

	return err
}
