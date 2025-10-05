package main

import (
	cli "github.com/kagent-dev/kagent/go/cli/internal/cli/agent"
	"github.com/spf13/cobra"
)


type DeployOpts struct {
	Image string
	APIKey string
	APIKeySecret string
	Namespace string
}

func DefaultDeployOpts() *DeployOpts {
	return &DeployOpts{
		Image:       "",
		APIKey:     "",
		APIKeySecret: "",
		Namespace:   "",
	}
}

func NewDeployCmd() *cobra.Command {
	o := DefaultDeployOpts()

	cmd := &cobra.Command{
		Use:   "deploy [project-directory]",
		Short: "Deploy an agent to Kubernetes",
		Example: "kagent deploy ./my-agent --api-key-secret \"my-existing-secret\"",
		SilenceErrors: true,
		SilenceUsage:  true,
		Long: `Deploy an agent to Kubernetes.

This command will read the kagent.yaml file from the specified project directory,
create or reference a Kubernetes secret with the API key, and create an Agent CRD.

The command will:
1. Load the agent configuration from kagent.yaml
2. Either create a new secret with the provided API key or verify an existing secret
3. Create an Agent CRD with the appropriate configuration

API Key Options:
  --api-key: Convenience option to create a new secret with the provided API key
  --api-key-secret: Canonical way to reference an existing secret by name

Examples:
  kagent deploy ./my-agent --api-key-secret "my-existing-secret"
  kagent deploy ./my-agent --api-key "your-api-key-here" --image "myregistry/myagent:v1.0"
  kagent deploy ./my-agent --api-key-secret "my-secret" --namespace "my-namespace"`,
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
				deployCfg := &cli.DeployCfg{
		Config: cfg,
	}
			deployCfg.ProjectDir = args[0]

			return cli.DeployCmd(cmd.Context(), deployCfg)
		},
	}

	// Add flags for deploy command
	cmd.Flags().StringVarP(&o.Image, "image", "i", o.Image, "Image to use (defaults to localhost:5001/{agentName}:latest)")
	cmd.Flags().StringVar(&o.APIKey, "api-key", o.APIKey, "API key for the model provider (convenience option to create secret)")
	cmd.Flags().StringVar(&o.APIKeySecret, "api-key-secret", o.APIKeySecret, "Name of existing secret containing API key")
	cmd.Flags().StringVar(&o.Namespace, "namespace", o.Namespace, "Kubernetes namespace to deploy to")

	return cmd
}