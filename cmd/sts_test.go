package cmd

import (
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func ErrorCmd() *cobra.Command {
	return &cobra.Command{
		Use: "test",
		RunE: func(cmd *cobra.Command, args []string) error {
			return common.NewCLIArgParseError(fmt.Errorf("test error"))
		},
	}
}

func TestHelp(t *testing.T) {
	cli := di.NewMockDeps()
	sts := STSCommand(&cli.Deps)
	sts.SetArgs([]string{"help"})
	exitCode := execute(cli.Context, &cli.Deps, sts)
	assert.Equal(t, 0, exitCode)
}

func TestWrongFlag(t *testing.T) {
	cli := di.NewMockDeps()
	sts := STSCommand(&cli.Deps)
	sts.SetArgs([]string{"version", "--wrongflag"})
	exitCode := execute(cli.Context, &cli.Deps, sts)
	assert.Equal(t, common.CommandFailedRequirementExitCode, exitCode)
	assert.Equal(t, []error{fmt.Errorf("unknown flag: --wrongflag")}, *cli.MockPrinter.PrintErrCalls)
	assert.Contains(t, (*cli.MockPrinter.PrintLnCalls)[0], "Usage:") // show help
}

func TestErr(t *testing.T) {
	cli := di.NewMockDeps()
	errorCmd := ErrorCmd()
	errorCmd.SetArgs([]string{})
	exitCode := execute(cli.Context, &cli.Deps, errorCmd)
	assert.Equal(t, common.CommandFailedRequirementExitCode, exitCode)
	expectedErrCalls := []error{common.NewCLIArgParseError(fmt.Errorf("test error"))}
	assert.Equal(t, &expectedErrCalls, cli.MockPrinter.PrintErrCalls)
}

func TestErrToJson(t *testing.T) {
	cli := di.NewMockDeps()
	errorCmd := ErrorCmd()
	errorCmd.SetArgs([]string{"--json"})
	exitCode := execute(cli.Context, &cli.Deps, errorCmd)
	assert.Equal(t, common.CommandFailedRequirementExitCode, exitCode)
	expectedJsonCalls := []error{common.NewCLIArgParseError(fmt.Errorf("test error"))}
	assert.Equal(t, &expectedJsonCalls, cli.MockPrinter.PrintErrJsonCalls)
}
