package main

import (
	cli "github.com/kagent-dev/kagent/go/cli/internal/cli/agent"
	"github.com/spf13/cobra"
)

func NewUninstallCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "uninstall",
		Short: "Uninstall kagent",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			cli.UninstallCmd(cmd.Context(), cfg)
		},
	}
}
