package cobra_util

import (
	"bytes"
	"context"

	"github.com/spf13/cobra"
)

func ExecuteCommandWithContext(ctx context.Context, root *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	root.SetOut(buf)
	root.SetErr(buf)
	root.SetArgs(args)
	err = root.ExecuteContext(ctx)
	return buf.String(), err
}
