package limiter

import (
	"context"
	"time"

	"golang.org/x/time/rate"
)

type rateLimiter interface {
	Wait(ctx context.Context) (err error)
	WaitN(ctx context.Context, n int) (err error)
}

type RateLimiter interface {
	Call(f func() error) error
}

type limiter struct {
	lim rateLimiter
}

func NewLimiter(rpm int) RateLimiter {
	l := rate.NewLimiter(rate.Every(time.Minute/time.Duration(rpm)), 1)
	lim := limiter{lim: l}

	return &lim
}

func (l *limiter) Call(f func() error) error {
	ctx := context.Background()

	err := l.lim.Wait(ctx)
	if err != nil {
		return err
	}

	return f()
}
