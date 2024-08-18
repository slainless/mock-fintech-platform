package tracker

import "context"

type NilTracker struct{}

func (t *NilTracker) Report(ctx context.Context, err error) {}
