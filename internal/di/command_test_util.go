package di

import (
	"bytes"

	"github.com/spf13/cobra"
	stscobra "gitlab.com/stackvista/stackstate-cli2/internal/cobra"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/pkg/pflags"
)

func ExecuteCommandWithContext(cli *Deps, cmd *cobra.Command, args ...string) (output string, err error) {
	buf := new(bytes.Buffer)
	cmd.SetOut(buf)
	cmd.SetErr(buf)
	cmd.SetArgs(args)
	if !cmd.PersistentFlags().HasAvailableFlags() { // to prevent double loading error on second call
		common.AddPersistentFlags(cmd)
	}

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if err := stscobra.ValidateMutexFlags(cmd); err != nil {
			return common.NewCLIArgParseError(err)
		}

		output, err := pflags.GetEnum(cmd.Flags(), common.OutputFlag)
		if err != nil {
			return err
		}

		cli.Output = common.ToOutput(output)
		if cli.Output == common.JSONOutput {
			cli.NoColor = true
		}

		return nil
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
