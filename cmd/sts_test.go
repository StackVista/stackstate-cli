package cmd

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupStsCmd() (*di.MockDeps, *cobra.Command, *bytes.Buffer, *bytes.Buffer) {
	cli := di.NewMockDeps()
	sts := STSCommand(&cli.Deps)
	stdOut := new(bytes.Buffer)
	stdErr := new(bytes.Buffer)
	sts.SetOut(stdOut)
	sts.SetErr(stdErr)
	return &cli, sts, stdOut, stdErr
}

func ErrorCmd() *cobra.Command {
	return &cobra.Command{
		Use: "test",
		RunE: func(cmd *cobra.Command, args []string) error {
			return common.NewCLIArgParseError(fmt.Errorf("test error"))
		},
	}
}

func TestHelpWhenRunningWithoutArgs(t *testing.T) {
	cli, cmd, stdOut, stdErr := setupStsCmd()
	cmd.SetArgs([]string{""})
	exitCode := execute(cli.Context, &cli.Deps, cmd)
	assert.Equal(t, 0, exitCode)
	assert.Contains(t, stdOut.String(), "Usage:")
	assert.Equal(t, "", stdErr.String())
}

func TestHelpWhenRunningHelp(t *testing.T) {
	cli, cmd, stdOut, stdErr := setupStsCmd()
	cmd.SetArgs([]string{"help"})
	exitCode := execute(cli.Context, &cli.Deps, cmd)
	assert.Equal(t, 0, exitCode)
	assert.Contains(t, stdOut.String(), "Usage:")
	assert.Equal(t, "", stdErr.String())
}

func TestVersion(t *testing.T) {
	cli := di.NewMockDeps()
	sts := STSCommand(&cli.Deps)
	sts.SetArgs([]string{"version"})
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
