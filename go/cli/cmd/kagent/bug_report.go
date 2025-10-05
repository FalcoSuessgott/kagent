package main

import (
	"fmt"
	cli "github.com/kagent-dev/kagent/go/cli/internal/cli/agent"
	"github.com/spf13/cobra"
)

func NewBugReportCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "bug-report",
		Short: "Generate a bug report",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cli.CheckServerConnection(cfg.Client()); err != nil {
				pf, err := cli.NewPortForward(cmd.Context(), cfg)
				if err != nil {
					return fmt.Errorf("Error starting port-forward: %v", err)
				}
				defer pf.Stop()
			}
			cli.BugReportCmd(cfg)

			return nil
		},
	}
}