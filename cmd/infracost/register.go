package main

import (
	"github.com/spf13/cobra"

	"github.com/infracost/infracost/pkg/config"
	"github.com/infracost/infracost/pkg/logging"
	"github.com/infracost/infracost/pkg/ui"
)

func registerCmd(ctx *config.RunContext) *cobra.Command {
	login := authLoginCmd(ctx)
	cmd := &cobra.Command{
		Use:    "register",
		Hidden: true,
		Short:  login.Short,
		Long:   login.Long,
		RunE: func(cmd *cobra.Command, args []string) error {
			logging.Logger.Warn().Msgf(
				"this command has been changed to %s, which does the same thing - we’ll run that for you now.\n",
				ui.PrimaryString("infracost auth login"),
			)

			return login.RunE(cmd, args)
		},
	}

	cmd.SetHelpFunc(func(cmd *cobra.Command, strings []string) {
		logging.Logger.Warn().Msgf(
			"this command has been changed to %s, which does the same thing - showing information for that command.\n",
			ui.PrimaryString("infracost auth login"),
		)

		login.HelpFunc()(login, strings)
	})

	return cmd
}
