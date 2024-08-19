package tracker

import (
	"context"
	"fmt"
	"os"
)

type LogTracker struct{}

func (t *LogTracker) Report(ctx context.Context, err error) {
	if err == nil {
		return
	}

	fmt.Fprintf(os.Stderr, "%s\n", err)
}
