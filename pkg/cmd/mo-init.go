package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"github.com/wanglei4687/mo-plugin/pkg/cmd/resources"
	"io"
	"k8s.io/klog/v2"
)

const (
	operatorInitDesc = `
'init' command creates Matrixone Operator deployment along with all the dependencies.
`
	operatorInitExample = ` kubectl mo init`
)

type operatorInitCmd struct {
	out          io.Writer
	errOut       io.Writer
	output       bool
	operatorOpts resources.OperatorOptions
}

func newInitCmd(out io.Writer, errOut io.Writer) *cobra.Command {
	o := &operatorInitCmd{out: out, errOut: errOut}
	cmd := &cobra.Command{
		Use:     "init",
		Short:   "Initialize MatrixOne Operator",
		Long:    operatorInitDesc,
		Example: operatorInitExample,
		Args:    cobra.MinimumNArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			err := o.run(out)
			if err != nil {
				klog.Warning(err)
				return err
			}
			return nil
		},
	}

	return cmd
}

func (o *operatorInitCmd) run(writer io.Writer) error {
	return nil
}

func (o *operatorInitCmd) serializeJSONPatchOps(jp []interface{}) string {
	jpJSON, _ := json.Marshal(jp)
	return string(jpJSON)
}
