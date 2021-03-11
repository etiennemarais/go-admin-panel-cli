package new

import (
	"context"
	"fmt"

	"github.com/etiennemarais/go-admin-panel-cli/pkg/env"
	"github.com/etiennemarais/go-admin-panel-cli/pkg/generate"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

func Command(ctx context.Context, logger *zap.Logger, env env.Env) *cobra.Command {
	options := generate.Options{}

	newCommand := &cobra.Command{
		Use:   "new",
		Short: "n",
		RunE: func(cmd *cobra.Command, args []string) error {
			logger.Info(fmt.Sprintf("Initializing admin panel config for (%s)", options.Repository))

			generate, _ := generate.New(ctx, logger, env, options)
			_ = generate.BuildAdminPanelConfig()

			return nil
		},
	}

	newCommand.Flags().StringVarP(&options.Repository, "repository", "r", "", "The source repository to scan for migrations (required)")
	newCommand.MarkFlagRequired("repository")

	return newCommand
}
