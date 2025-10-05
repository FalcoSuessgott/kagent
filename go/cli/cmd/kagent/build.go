package main

import (
	cli "github.com/kagent-dev/kagent/go/cli/internal/cli/agent"
	"github.com/spf13/cobra"
)


type BuildOpts struct {
	Image string
	Push   bool

}
func DefaultBuildOpts() *BuildOpts {
	return &BuildOpts{
		Image:       "",
	}
}

func NewBuildCmd() *cobra.Command {
	o := DefaultBuildOpts()

	cmd := &cobra.Command{
		Use:   "build [project-directory]",
		Short: "Build a Docker image for an agent project",
		Example: "kagent build ./my-agent",
		SilenceErrors: true,
		SilenceUsage:  true,
		Long: `Build a Docker image for an agent project created with the init command.

This command will look for a Dockerfile in the specified project directory and build
a Docker image using docker build. The image can optionally be pushed to a registry.

Image naming:
- If --image is provided, it will be used as the full image specification (e.g., ghcr.io/myorg/my-agent:v1.0.0)
- Otherwise, defaults to localhost:5001/{agentName}:latest where agentName is loaded from kagent.yaml

Examples:
  kagent build ./my-agent
  kagent build ./my-agent --image ghcr.io/myorg/my-agent:v1.0.0
  kagent build ./my-agent --image ghcr.io/myorg/my-agent:v1.0.0 --push`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
				buildCfg := &cli.BuildCfg{
		Config: cfg,
	}
			buildCfg.ProjectDir = args[0]

			return cli.BuildCmd(buildCfg)
		},
	}

	cmd.Flags().StringVar(&o.Image, "image", "", "Full image specification (e.g., ghcr.io/myorg/my-agent:v1.0.0)")
	cmd.Flags().BoolVar(&o.Push, "push", false, "Push the image to the registry")

	return cmd
}