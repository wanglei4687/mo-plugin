package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"log"
)

const (
	moDesc     = "Manage and deploy Matrixone Cluster"
	kubeConfig = "kubeconfig"
)

var (
	moExample = `
# show all resources
%[1]s mo

# install matrixone cluster
$[1]s mo install
`
	confPath string
	rootCmd  = &cobra.Command{
		Use:          "mo [command] [flags]",
		Long:         moDesc,
		SilenceUsage: false,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&confPath, kubeConfig, "", "Custom kubeconfig path")

	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

// New provides a cobra command wrapping mo options
func New(streams genericclioptions.IOStreams) *cobra.Command {
	rootCmd = DisableHelp(rootCmd)
	cobra.EnableCommandSorting = false

	return rootCmd
}
