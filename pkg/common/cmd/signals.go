package cmd

import (
	"context"
	"os/signal"
	"syscall"
)

func Context() context.Context {
	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	return ctx
}
