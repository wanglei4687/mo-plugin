package cmd

import "github.com/spf13/cobra"

func DisableHelp(cmd *cobra.Command) *cobra.Command {
	cmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

	return cmd
}
