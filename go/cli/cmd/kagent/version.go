package main

import (
	"fmt"
	cli "github.com/kagent-dev/kagent/go/cli/internal/cli/agent"
	"github.com/spf13/cobra"
)

func NewVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "Print the kagent version",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error{
			if err := cli.CheckServerConnection(cfg.Client()); err != nil {
				pf, err := cli.NewPortForward(cmd.Context(), cfg)
				if err != nil {
					return fmt.Errorf("Error starting port-forward: %v", err)
				}
				defer pf.Stop()
			}
			cli.VersionCmd(cfg)

			return nil
		},
	}
}