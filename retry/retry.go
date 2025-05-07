package retry

import (
	"math"
	"time"
)

type retryOptions struct {
	maxRetries int
	delay      time.Duration
}

func (r *retryOptions) retry(fn func() error) error {
	var err error
	for range r.maxRetries {
		err = fn()
		if err == nil {
			break
		}

		time.Sleep(r.delay)
		r.delay *= 2
	}

	return err
}

func RetrySuperShort(fn func() error) error {
	ro := retryOptions{
		maxRetries: 5,
		delay:      5 * time.Second,
	}
	return ro.retry(fn)
}

func RetrySimple(fn func() error) error {
	ro := retryOptions{
		maxRetries: 3,
		delay:      30 * time.Second,
	}
	return ro.retry(fn)
}

func RetryMedium(fn func() error) error {
	ro := retryOptions{
		maxRetries: 5,
		delay:      30 * time.Second,
	}
	return ro.retry(fn)
}

func RetryDebug(fn func() error) error {
	ro := retryOptions{
		maxRetries: 1,
		delay:      1 * time.Second,
	}
	return ro.retry(fn)
}

func RetryLongForever(fn func() error) error {
	ro := retryOptions{
		maxRetries: math.MaxInt32,
		delay:      60 * time.Second,
	}
	return ro.retry(fn)
}
