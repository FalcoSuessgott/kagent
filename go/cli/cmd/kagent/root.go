package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	cli "github.com/kagent-dev/kagent/go/cli/internal/cli/agent"
	"github.com/kagent-dev/kagent/go/cli/internal/cli/mcp"
	"github.com/kagent-dev/kagent/go/cli/internal/config"
	"github.com/kagent-dev/kagent/go/cli/internal/tui"
	"github.com/spf13/cobra"
)

var cfg *config.Config

type ConfigOptions struct {
	KAgentURL   string
	Namespace   string
	OutputFormat string
	Verbose     bool
	Timeout     time.Duration
}

func DefaultConfigOptions() *ConfigOptions{
	return &ConfigOptions{
		KAgentURL:   "http://localhost:8083",
		Namespace:   "kagent",
		OutputFormat: "table",
		Verbose:     false,
		Timeout:     300 * time.Second,
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// listen for signals to cancel the context throughout the application
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-done

		fmt.Fprintf(os.Stderr, "kagent aborted.\n")
		fmt.Fprintf(os.Stderr, "Exiting.\n")

		cancel()
	}()

	cfg = &config.Config{}
	
	if err := NewRootCmd().ExecuteContext(ctx); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)

		os.Exit(1)
	}
}

func NewRootCmd() *cobra.Command {
	o := DefaultConfigOptions()

	cmd := &cobra.Command{
		Use:   "kagent",
		Short: "kagent is a CLI and TUI for kagent",
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			return config.Init()
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			cfg, err := config.Get()
			if err != nil {
				return fmt.Errorf("Error getting config: %v", err)
			}

			client := cfg.Client()
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()

			// Start port forward and ensure it is healthy.
			var pf *cli.PortForward
			if err := cli.CheckServerConnection(client); err != nil {
				pf, err = cli.NewPortForward(ctx, cfg)
				if err != nil {
					return fmt.Errorf("Error starting port-forward: %v", err)
				}
				defer pf.Stop()
			}

			if err := tui.RunWorkspace(cfg, cfg.Client(), cfg.Verbose); err != nil {
				return fmt.Errorf("TUI error: %v", err)
			}

			return nil
		},
	}

	cmd.PersistentFlags().StringVar(&o.KAgentURL, "kagent-url", o.KAgentURL, "KAgent URL")
	cmd.PersistentFlags().StringVarP(&o.Namespace, "namespace", "n", o.Namespace, "Namespace")
	cmd.PersistentFlags().StringVarP(&o.OutputFormat, "output-format", "o", o.OutputFormat, "Output format")
	cmd.PersistentFlags().BoolVarP(&o.Verbose, "verbose", "v", o.Verbose, "Verbose output")
	cmd.PersistentFlags().DurationVar(&o.Timeout, "timeout", o.Timeout, "Timeout")

	cmd.AddCommand(
		NewInstallCmd(), 
		NewUninstallCmd(),
		NewInvokeCmd(),
		NewBugReportCmd(),
		NewVersionCmd(),
		NewDashboardCmd(),
		NewBugReportCmd(),
		NewGetCmd(),
		NewInitCmd(),
		NewBuildCmd(),
		NewDeployCmd(),
		mcp.NewMCPCmd(),
	)

	return cmd
}
