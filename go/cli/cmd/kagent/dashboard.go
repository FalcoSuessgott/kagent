package main

import (
	cli "github.com/kagent-dev/kagent/go/cli/internal/cli/agent"
	"github.com/spf13/cobra"
)

func NewDashboardCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "dashboard",
		Short: "Open the kagent dashboard",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			cli.DashboardCmd(cmd.Context(), cfg)
		},
	}
}