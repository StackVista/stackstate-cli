package di

import (
	"bytes"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
)

func ExecuteCommandWithContext(cli *Deps, cmd *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)
	if !cmd.PersistentFlags().HasAvailableFlags() { // to prevent double loading error on second call
		common.AddPersistentFlags(cmd)
	}

	for i, s := range args {
		if s == "--output" || s == "-o" {
			cli.Output = common.ToOutput(args[i+1])
		}
	}

	err = cmd.ExecuteContext(cli.Context)
	return buf.String(), err
}

func ExecuteCommandWithContextUnsafe(cli *Deps, cmd *cobra.Command, args ...string) string {
	res, err := ExecuteCommandWithContext(cli, cmd, args...)
	if err != nil {
		panic(err)
	}
	return res
}
