package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"

	"github.com/rudderlabs/rudder-go-kit/config"
	"github.com/rudderlabs/rudder-go-kit/logger"

	ptrans "github.com/rudderlabs/rudder-server/processor/transformer"
)

const (
	processorTransformer = "PROCESSOR_TRANSFORMER"
	warehouseTransformer = "WAREHOUSE_TRANSFORMER"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func mustInt(s string) int { _ = "STUB: not implemented"; return 0 }

func mustString(s string) string { _ = "STUB: not implemented"; return "" }

func selectTransformer(
	mode string,
	conf *config.Config,
	l logger.Logger,
) (ptrans.DestinationClient, error) {
	_ = "STUB: not implemented"
	return *new(ptrans.DestinationClient), nil
}
