package main

import (
	"context"
	"fmt"
	"slog-implementation/logging"

	"log/slog"
)

func main() {
	logger := logging.NewLogger()

	err := fmt.Errorf("imagine this a error message")

	logger.Error(
		"Some troubles", "failed t4.Modify ERROR", err)
	logger.Warn("Warn message")
	logger.Error("Error message")

	ctx := context.Background()
	logger.InfoContext(ctx, "Contextual info")

	logger.LogAttrs(ctx, slog.LevelWarn, "Log with attrs", slog.String("key", "value"))
}
