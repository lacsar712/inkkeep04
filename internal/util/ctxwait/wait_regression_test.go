package ctxwait

import (
	"context"
	"testing"
	"time"
)

// TestUntilReturnsImmediatelyOnPreCancelledCtx ensures that a context
// cancelled before Until is even called does not block for the full
// duration; it must return the cancel error right away.
func TestUntilReturnsImmediatelyOnPreCancelledCtx(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	start := time.Now()
	err := Until(ctx, 5*time.Second)
	elapsed := time.Since(start)

	if err == nil {
		t.Fatal("expected cancel error from pre-cancelled ctx")
	}
	if elapsed > 250*time.Millisecond {
		t.Fatalf("Until blocked on pre-cancelled ctx, elapsed=%s", elapsed)
	}
}

// TestUntilNegativeDurationReturnsAtOnce covers the d<=0 short-circuit:
// it must not sleep and must surface ctx.Err() immediately.
func TestUntilNegativeDurationReturnsAtOnce(t *testing.T) {
	ctx := context.Background()
	start := time.Now()
	if err := Until(ctx, -1); err != nil {
		t.Fatalf("expected nil for non-cancelled ctx with d<=0, got %v", err)
	}
	if time.Since(start) > 100*time.Millisecond {
		t.Fatal("negative duration should not sleep")
	}
}
