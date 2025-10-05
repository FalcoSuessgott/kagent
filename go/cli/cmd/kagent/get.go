package main

import (
	"os"
	"fmt"

	"github.com/spf13/cobra"
	cli "github.com/kagent-dev/kagent/go/cli/internal/cli/agent"
)

func NewGetCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a kagent resource",
		Long:  "Get a kagent resource",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Fprintf(os.Stderr, "No resource type provided\n\n")

			return cmd.Help()
		},
	}

	cmd.AddCommand(
		NewGetSessionCmd(),
		NewGetAgentCmd(),
		NewGetToolCmd(),
	)

	return cmd	
}

func NewGetSessionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "session [session_id]",
		Short: "Get a session or list all sessions",
		Long:  "Get a session by ID or list all sessions",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cli.CheckServerConnection(cfg.Client()); err != nil {
				pf, err := cli.NewPortForward(cmd.Context(), cfg)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error starting port-forward: %v\n", err)
					return
				}
				defer pf.Stop()
			}
			resourceName := ""
			if len(args) > 0 {
				resourceName = args[0]
			}
			cli.GetSessionCmd(cfg, resourceName)
		},
	}
}

func NewGetAgentCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "agent [agent_name]",
		Short: "Get an agent or list all agents",
		Long:  "Get an agent by name or list all agents",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cli.CheckServerConnection(cfg.Client()); err != nil {
				pf, err := cli.NewPortForward(cmd.Context(), cfg)
				if err != nil {
					return
				}
				defer pf.Stop()
			}
			resourceName := ""
			if len(args) > 0 {
				resourceName = args[0]
			}
			cli.GetAgentCmd(cfg, resourceName)
		},
	}
}

func NewGetToolCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tool",
		Short: "Get tools",
		Long:  "List all available tools",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			if err := cli.CheckServerConnection(cfg.Client()); err != nil {
				pf, err := cli.NewPortForward(cmd.Context(), cfg)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Error starting port-forward: %v\n", err)
					return
				}
				defer pf.Stop()
			}
			cli.GetToolCmd(cfg)
		},
	}
}