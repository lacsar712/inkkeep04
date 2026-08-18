package ctxwait

import (
	"context"
	"time"
)

// Until blocks until either d elapses or ctx is cancelled, whichever
// comes first. It returns ctx.Err() (nil while still active) so callers
// can distinguish a clean wait from a cancelled one. Unlike time.Sleep,
// it returns promptly when ctx is cancelled.
func Until(ctx context.Context, d time.Duration) error {
	if d <= 0 {
		return ctx.Err()
	}
	t := time.NewTimer(d)
	defer t.Stop()
	select {
	case <-t.C:
		return ctx.Err()
	case <-ctx.Done():
		return ctx.Err()
	}
}
