package main

import (
	"bytes"
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/cmd"
	"gitlab.com/stackvista/stackstate-cli2/internal/common"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

func setupSTSCmd() (*di.MockDeps, *cobra.Command, *bytes.Buffer, *bytes.Buffer) {
	cli := di.NewMockDeps()
	sts := cmd.STSCommand(&cli.Deps)
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
	cli, cmd, stdOut, stdErr := setupSTSCmd()
	cmd.SetArgs([]string{""})
	exitCode := execute(cli.Context, &cli.Deps, cmd)
	assert.Equal(t, 0, exitCode)
	assert.Contains(t, stdOut.String(), "Usage:")
	assert.Equal(t, "", stdErr.String())
}

func TestHelpWhenRunningHelp(t *testing.T) {
	cli, cmd, stdOut, stdErr := setupSTSCmd()
	cmd.SetArgs([]string{"help"})
	exitCode := execute(cli.Context, &cli.Deps, cmd)
	assert.Equal(t, 0, exitCode)
	assert.Contains(t, stdOut.String(), "Usage:")
	assert.Equal(t, "", stdErr.String())
}

func TestVersionRun(t *testing.T) {
	cli, cmd, _, _ := setupSTSCmd()
	cmd.SetArgs([]string{"version"})
	exitCode := execute(cli.Context, &cli.Deps, cmd)
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, []printer.TableData{{
		Header: []string{"Version", "Date", "CLI Type", "Commit"},
		Data:   [][]interface{}{{"1.0.0", "1-1-2022", "full", "123124"}},
	}}, *cli.MockPrinter.TableCalls)
}

func TestVersionJsonRun(t *testing.T) {
	cli, cmd, _, _ := setupSTSCmd()
	cmd.SetArgs([]string{"version", "--json"})
	exitCode := execute(cli.Context, &cli.Deps, cmd)
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, []map[string]interface{}{{
		"cli-type": "full", "commit": "123124", "date": "1-1-2022", "version": "1.0.0",
	}}, *cli.MockPrinter.PrintJsonCalls)
}

func TestWrongFlagError(t *testing.T) {
	cli := di.NewMockDeps()
	sts := cmd.STSCommand(&cli.Deps)
	sts.SetArgs([]string{"version", "--wrongflag"})
	exitCode := execute(cli.Context, &cli.Deps, sts)
	assert.Equal(t, common.CommandFailedRequirementExitCode, exitCode)
	assert.Equal(t, []error{fmt.Errorf("unknown flag: --wrongflag")}, *cli.MockPrinter.PrintErrCalls)
	assert.Contains(t, (*cli.MockPrinter.PrintLnCalls)[0], "Usage:") // show help
}

func TestErrHandling(t *testing.T) {
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

func TestColorUsedByDefault(t *testing.T) {
	cli, cmd, _, _ := setupSTSCmd()
	cmd.SetArgs([]string{"version"})
	execute(cli.Context, &cli.Deps, cmd)
	assert.True(t, cli.MockPrinter.UseColor)
	assert.False(t, cli.NoColor)
}

func TestNoColorFlag(t *testing.T) {
	cli, cmd, _, _ := setupSTSCmd()
	cmd.SetArgs([]string{"version", "--no-color"})
	execute(cli.Context, &cli.Deps, cmd)
	assert.False(t, cli.MockPrinter.UseColor)
	assert.True(t, cli.NoColor)
}

func TestNoColorWhenDumbTerminal(t *testing.T) {
	term := os.Getenv("TERM")
	defer os.Setenv("TERM", term)
	os.Setenv("TERM", "Dumb")

	cli, cmd, _, _ := setupSTSCmd()
	cmd.SetArgs([]string{"version"})
	execute(cli.Context, &cli.Deps, cmd)
	assert.False(t, cli.MockPrinter.UseColor)
	assert.True(t, cli.NoColor)
}

func TestNoColorWhenNoColorEnvExists(t *testing.T) {
	defer os.Unsetenv("NO_COLOR")
	// https://no-color.org says we need to check for existence, not a specific value
	os.Setenv("NO_COLOR", "false")

	cli, cmd, _, _ := setupSTSCmd()
	cmd.SetArgs([]string{"version"})
	execute(cli.Context, &cli.Deps, cmd)
	assert.False(t, cli.MockPrinter.UseColor)
	assert.True(t, cli.NoColor)
}

func TestNoColorWhenJsonOutput(t *testing.T) {
	cli, cmd, _, _ := setupSTSCmd()
	cmd.SetArgs([]string{"version", "--json"})
	execute(cli.Context, &cli.Deps, cmd)
	assert.False(t, cli.MockPrinter.UseColor)
	assert.True(t, cli.NoColor)
}

func TestVerbose(t *testing.T) {
	cli, cmd, _, _ := setupSTSCmd()
	cmd.SetArgs([]string{"version", "--verbose"})
	execute(cli.Context, &cli.Deps, cmd)
	assert.True(t, cli.IsVerBose)
}
