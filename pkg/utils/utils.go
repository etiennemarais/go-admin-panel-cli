package utils

import (
	"context"

	"github.com/etiennemarais/go-admin-panel-cli/pkg/env"
	"github.com/kelseyhightower/envconfig"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func Bootstrap() (*zap.Logger, *cobra.Command, context.Context, env.Env) {
	logger, _ := zap.NewDevelopment()
	cmd := &cobra.Command{Use: "Go Admin Panel CLI"}

	var e env.Env
	envconfig.MustProcess("", &e)

	return logger, cmd, context.Background(), e
}
