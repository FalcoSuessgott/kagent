package main

import (
	cli "github.com/kagent-dev/kagent/go/cli/internal/cli/agent"
	"github.com/spf13/cobra"
)

type InvokeOpts struct {
	Task        string
	Session     string
	Agent       string
	Stream      bool
	File        string
	URLOverride string
}

func DefaultInvokeOpts() *InvokeOpts {
	return &InvokeOpts{
		Task:        "",
		Session:     "",
		Agent:       "",
		Stream:      false,
		File:        "",
		URLOverride: "",
	}
}

func NewInvokeCmd() *cobra.Command {
	o := DefaultInvokeOpts()

	cmd := &cobra.Command{
		Use:   "invoke",
		Short: "Invoke a kagent agent",
		SilenceErrors: true,
		SilenceUsage:  true,
		Example: `kagent invoke --agent "k8s-agent" --task "Get all the pods in the kagent namespace"`,
		Run: func(cmd *cobra.Command, args []string) {
			invokeCfg := &cli.InvokeCfg{
				Config: cfg,
			}

			cli.InvokeCmd(cmd.Context(), invokeCfg)
		},
	}

	cmd.Flags().StringVarP(&o.Task, "task", "t", "", "Task")
	cmd.Flags().StringVarP(&o.Session, "session", "s", "", "Session")
	cmd.Flags().StringVarP(&o.Agent, "agent", "a", "", "Agent")
	cmd.Flags().BoolVarP(&o.Stream, "stream", "S", false, "Stream the response")
	cmd.Flags().StringVarP(&o.File, "file", "f", "", "File to read the task from")
	cmd.Flags().StringVarP(&o.URLOverride, "url-override", "u", "", "URL override")
	cmd.Flags().MarkHidden("url-override") //nolint:errcheck

	return cmd
}