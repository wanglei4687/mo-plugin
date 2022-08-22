package main

import (
	"github.com/spf13/pflag"
	"github.com/wanglei4687/mo-plugin/pkg/cmd"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"os"
)

func main() {
	flags := pflag.NewFlagSet("kubectl-mo", pflag.ExitOnError)
	pflag.CommandLine = flags

	root := cmd.New(genericclioptions.IOStreams{
		In:     os.Stdin,
		Out:    os.Stdout,
		ErrOut: os.Stderr})

	if err := root.Execute(); err != nil {
		os.Exit(1)
	}
}
