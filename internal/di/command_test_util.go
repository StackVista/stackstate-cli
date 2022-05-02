package di

import (
	"bytes"

	"github.com/spf13/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func ExecuteCommandWithContext(cli *Deps, cmd *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)
	if !cmd.PersistentFlags().HasAvailableFlags() { // to prevent double loading error on second call
		common.AddPersistentFlags(cmd)
	}
	if util.StringInSlice("--json", args) {
		cli.IsJson = true
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
