package main

import (
	cli "github.com/kagent-dev/kagent/go/cli/internal/cli/agent"
	"github.com/kagent-dev/kagent/go/cli/internal/profiles"
	"github.com/spf13/cobra"
)

type InstallOpts struct {
	Profile string
}

func DefaultInstallOpts() *InstallOpts {
	return &InstallOpts{
		Profile: "minimal",
	}
}
func NewInstallCmd() *cobra.Command {
	o := &InstallOpts{}
	cmd := &cobra.Command{
		Use:   "install",
		Short: "Install kagent",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			installCfg := &cli.InstallCfg{
				Config: cfg,
			}

			cli.InstallCmd(cmd.Context(), installCfg)
		},
	}

	cmd.Flags().StringVar(&o.Profile, "profile", o.Profile, "Installation profile (minimal|demo)")
	_ = cmd.RegisterFlagCompletionFunc("profile", func(cmd *cobra.Command, args []string, toComplete string) ([]string, cobra.ShellCompDirective) {
		return profiles.Profiles, cobra.ShellCompDirectiveNoFileComp
	})

	return cmd
}