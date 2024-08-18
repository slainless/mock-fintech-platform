package platform

import "context"

type ErrorTracker interface {
	// should skip report if err is nil
	Report(ctx context.Context, err error)
}
