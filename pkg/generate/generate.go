package generate

import (
	"context"

	"github.com/etiennemarais/go-admin-panel-cli/pkg/env"
	"go.uber.org/zap"
)

type Generate struct {
	ctx     context.Context
	logger  *zap.Logger
	env     env.Env
	options Options
}

type Options struct {
	Repository string
}

func New(ctx context.Context, logger *zap.Logger, env env.Env, options Options) (Generate, error) {
	generate := Generate{
		ctx:     ctx,
		logger:  logger,
		env:     env,
		options: options,
	}

	// Check for repo here

	return generate, nil
}

func (g *Generate) BuildAdminPanelConfig() error {
	g.logger.Info("Generating admin panel config file from resources")

	return nil
}
