package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/cmd"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

func setupSTSCmd(t *testing.T) (*di.MockDeps, *cobra.Command) {
	errorCmd := cobra.Command{
		Use:   "test-error",
		Short: "some short help text for ErrorCmd",
		Long:  "Some long help text for ErrorCmd.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return common.NewCLIArgParseError(fmt.Errorf("test error"))
		},
	}
	cli := di.NewMockDeps(t)
	sts := cmd.STSCommand(&cli.Deps)
	sts.AddCommand(&errorCmd)
	sts.SetOut(cli.StdOut)
	sts.SetErr(cli.StdErr)

	return &cli, sts
}

func TestHelpWhenRunningWithoutArgs(t *testing.T) {
	cli, cmd := setupSTSCmd(t)
	cmd.SetArgs([]string{""})
	exitCode := execute(cli.Context, &cli.Deps, cmd)
	assert.Equal(t, 0, exitCode)
	assert.Contains(t, cli.StdOut.String(), "Usage:")
	assert.NotContains(t, cli.StdOut.String(), "flags")
	assert.Equal(t, "", cli.StdErr.String())
}

func TestHelpWhenRunningHelp(t *testing.T) {
	cli, cmd := setupSTSCmd(t)
	cmd.SetArgs([]string{"help"})
	exitCode := execute(cli.Context, &cli.Deps, cmd)
	assert.Equal(t, 0, exitCode)
	assert.Contains(t, cli.StdOut.String(), "Usage:")
	assert.Equal(t, "", cli.StdErr.String())
}

func TestVersionRun(t *testing.T) {
	cli, cmd := setupSTSCmd(t)
	cmd.SetArgs([]string{"version"})
	exitCode := execute(cli.Context, &cli.Deps, cmd)
	assert.Equal(t, 0, exitCode)
	assert.Equal(t, []printer.TableData{{
		Header: []string{"Version", "Build Date", "Commit"},
		Data:   [][]interface{}{{"1.0.0", "1-1-2022", "123124"}},
	}}, *cli.MockPrinter.TableCalls)
}

func TestVersionJsonRun(t *testing.T) {
	cli, cmd := setupSTSCmd(t)

	cmd.SetArgs([]string{"version", "-o", "json"})
	exitCode := execute(cli.Context, &cli.Deps, cmd)

	assert.Equal(t, 0, exitCode)
	assert.Equal(t, []map[string]interface{}{{
		"commit": "123124", "build-date": "1-1-2022", "version": "1.0.0",
	}}, *cli.MockPrinter.PrintJsonCalls)
}

func TestSetFlagErrorFuncSetsCLIArgParseError(t *testing.T) {
	cli, sts := setupSTSCmd(t)

	sts.SetArgs([]string{"version", "--wrongflag"})
	exitCode := execute(cli.Context, &cli.Deps, sts)

	assert.Equal(t, common.CommandFailedRequirementExitCode, exitCode)
	assert.Equal(t, []error{common.NewCLIArgParseError(fmt.Errorf("unknown flag: --wrongflag"))}, *cli.MockPrinter.PrintErrCalls)
	assert.Contains(t, (*cli.MockPrinter.PrintLnCalls)[0], "Usage:") // show help
}

func TestErrHandling(t *testing.T) {
	cli, sts := setupSTSCmd(t)

	sts.SetArgs([]string{"test-error"})
	exitCode := execute(cli.Context, &cli.Deps, sts)

	assert.Equal(t, common.CommandFailedRequirementExitCode, exitCode)
	expectedErrCalls := []error{common.NewCLIArgParseError(fmt.Errorf("test error"))}
	assert.Equal(t, &expectedErrCalls, cli.MockPrinter.PrintErrCalls)
	assert.Contains(t, cli.StdOut.String(), "Usage:\n  sts test-error [flags]") // show help
}

func TestErrToJson(t *testing.T) {
	cli, sts := setupSTSCmd(t)

	sts.SetArgs([]string{"test-error", "-o", "json"})
	exitCode := execute(cli.Context, &cli.Deps, sts)

	assert.Equal(t, common.CommandFailedRequirementExitCode, exitCode)
	expectedJsonCalls := []error{common.NewCLIArgParseError(fmt.Errorf("test error"))}
	assert.Equal(t, &expectedJsonCalls, cli.MockPrinter.PrintErrJsonCalls)
}

func TestColorUsedByDefault(t *testing.T) {
	term := os.Getenv("TERM")
	defer os.Setenv("TERM", term)
	os.Setenv("TERM", "xterm-256color")

	noColor := os.Getenv("NO_COLOR")
	if _, exists := os.LookupEnv("NO_COLOR"); exists {
		defer func() {
			os.Setenv("NO_COLOR", noColor)
		}()
		os.Unsetenv("NO_COLOR")
	}

	cli, cmd := setupSTSCmd(t)

	cmd.SetArgs([]string{"version"})
	execute(cli.Context, &cli.Deps, cmd)

	assert.True(t, cli.MockPrinter.UseColor)
	assert.False(t, cli.NoColor)
}

func TestNoColorFlag(t *testing.T) {
	cli, cmd := setupSTSCmd(t)
	cmd.SetArgs([]string{"version", "--no-color"})
	execute(cli.Context, &cli.Deps, cmd)
	assert.False(t, cli.MockPrinter.UseColor)
	assert.True(t, cli.NoColor)
}

func TestNoColorWhenDumbTerminal(t *testing.T) {
	term := os.Getenv("TERM")
	defer os.Setenv("TERM", term)
	os.Setenv("TERM", "Dumb")

	cli, cmd := setupSTSCmd(t)
	cmd.SetArgs([]string{"version"})
	execute(cli.Context, &cli.Deps, cmd)
	assert.False(t, cli.MockPrinter.UseColor)
	assert.True(t, cli.NoColor)
}

func TestNoColorWhenNoColorEnvExists(t *testing.T) {
	defer os.Unsetenv("NO_COLOR")
	// https://no-color.org says we need to check for existence, not a specific value
	os.Setenv("NO_COLOR", "false")

	cli, cmd := setupSTSCmd(t)
	cmd.SetArgs([]string{"version"})
	execute(cli.Context, &cli.Deps, cmd)
	assert.False(t, cli.MockPrinter.UseColor)
	assert.True(t, cli.NoColor)
}

func TestNoColorWhenJsonOutput(t *testing.T) {
	cli, cmd := setupSTSCmd(t)
	cmd.SetArgs([]string{"version", "-o", "json"})
	execute(cli.Context, &cli.Deps, cmd)
	assert.False(t, cli.MockPrinter.UseColor)
	assert.True(t, cli.NoColor)
}

func TestVerbose(t *testing.T) {
	cli, cmd := setupSTSCmd(t)
	cmd.SetArgs([]string{"version", "--verbose"})
	execute(cli.Context, &cli.Deps, cmd)
	assert.True(t, cli.IsVerBose)
}

func TestErrorOnUnknownSubCommand(t *testing.T) {
	cli, cmd := setupSTSCmd(t)
	cmd.SetArgs([]string{"context", "blaat"})
	exitCode := execute(cli.Context, &cli.Deps, cmd)
	assert.Equal(t, common.CommandFailedRequirementExitCode, exitCode)
	assert.Contains(t, cli.StdErr.String(), "Unknown sub-command \"blaat\" for \"context\"")
}

func TestHelpOnMissingSubCommand(t *testing.T) {
	cli, cmd := setupSTSCmd(t)
	cmd.SetArgs([]string{"context"})
	exitCode := execute(cli.Context, &cli.Deps, cmd)
	assert.Equal(t, common.OkExitCode, exitCode)
	assert.Equal(t, "", cli.StdErr.String())
	assert.Contains(t, cli.StdOut.String(), "Usage:")
	assert.NotContains(t, cli.StdOut.String(), "flags")
}

// thanks https://stackoverflow.com/a/33404435/1860591
func TestExitCodeOnError(t *testing.T) {
	if os.Getenv("BE_CRASHER") == "1" {
		os.Args = []string{"unknown_cmd", "--unknown-flag"}
		main()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestExitCodeOnError") //nolint:gosec,noctx
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
