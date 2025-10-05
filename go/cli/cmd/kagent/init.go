package main

import (
	cli "github.com/kagent-dev/kagent/go/cli/internal/cli/agent"
	"github.com/spf13/cobra"
)


type InitOpts struct {
	InstructionFile string
	ModelProvider   string
	ModelName       string
	Description     string
}

func DefaultInitOpts() *InitOpts {
	return &InitOpts{
		InstructionFile: "",
		ModelProvider:   "Gemini",
		ModelName:       "gemini-2.0-flash",
		Description:     "",
	}
}

func NewInitCmd() *cobra.Command {
	o := DefaultInitOpts()

	cmd := &cobra.Command{
		Use:   "init [framework] [language] [agent-name]",
		Short: "Initialize a new agent project",
		Example: "kagent init adk python dice",
		SilenceErrors: true,
		SilenceUsage:  true,
		Long: `Initialize a new agent project using the specified framework and language.

You can customize the root agent instructions using the --instruction-file flag.
You can select a specific model using --model-provider and --model-name flags.
If no custom instruction file is provided, a default dice-rolling instruction will be used.
If no model is specified, the agent will need to be configured later.

Examples:
  kagent init adk python dice
  kagent init adk python dice --instruction-file instructions.md
  kagent init adk python dice --model-provider Gemini --model-name gemini-2.0-flash`,
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			initCfg := &cli.InitCfg{
				Config: cfg,
			}

			initCfg.Framework = args[0]
			initCfg.Language = args[1]
			initCfg.AgentName = args[2]

			return  cli.InitCmd(initCfg)
		},
	}

	cmd.Flags().StringVar(&o.InstructionFile, "instruction-file", o.InstructionFile, "Path to file containing custom instructions for the root agent")
	cmd.Flags().StringVar(&o.ModelProvider, "model-provider", o.ModelProvider, "Model provider (OpenAI, Anthropic, Gemini)")
	cmd.Flags().StringVar(&o.ModelName, "model-name", o.ModelName, "Model name (e.g., gpt-4, claude-3-5-sonnet, gemini-2.0-flash)")
	cmd.Flags().StringVar(&o.Description, "description", o.Description, "Description for the agent")

	return cmd
}