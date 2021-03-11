package main

import (
	"github.com/etiennemarais/go-admin-panel-cli/cmd/new"
	"github.com/etiennemarais/go-admin-panel-cli/pkg/utils"

	"go.uber.org/zap"
)

func main() {
	logger, cmd, ctx, env := utils.Bootstrap()

	defer func() {
		_ = logger.Sync()
	}()

	cmd.AddCommand(
		new.Command(ctx, logger, env),
	)

	if err := cmd.ExecuteContext(ctx); err != nil {
		logger.Fatal("execute command", zap.Error(err))
	}
}
