package stackpack

import (
	"fmt"
	"testing"
	"time"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

var (
	testDisplayName   = "Test StackPack"
	testCurrVersion   = "1.2.3"
	testLatestVersion = "3.2.1"
	testCat1          = "Foo"
	testCat2          = "Bar"
	testCategories    = []string{testCat1, testCat2}
	testInt1          = stackstate_api.StackPackIntegration{
		DisplayName: testCat1,
	}
	testInt2 = stackstate_api.StackPackIntegration{
		DisplayName: testCat2,
	}
	testIntegrations = []stackstate_api.StackPackIntegration{testInt1, testInt2}
	testStepType     = "text"
	testStep1Name    = "step1"
	testStep1        = stackstate_api.StackPackStep{
		Name:    &testStep1Name,
		Display: &testStep1Name,
		Value: &stackstate_api.StackPackStepValue{
			Type: &testStepType,
		},
	}
	testStep2Name = "step2"
	testStep2     = stackstate_api.StackPackStep{
		Name:    &testStep2Name,
		Display: &testStep2Name,
		Value: &stackstate_api.StackPackStepValue{
			Type: &testStepType,
		},
	}
	testSteps     = []stackstate_api.StackPackStep{testStep1, testStep2}
	testId        = int64(23)
	testLut       = int64(5)
	testTimestamp = time.UnixMilli(5)
	testStatus    = "Tested with Adam Savage"
	testStackPack = stackstate_api.FullStackPack{
		Name:         testName,
		DisplayName:  testDisplayName,
		Version:      testCurrVersion,
		Categories:   testCategories,
		Integrations: testIntegrations,
		Steps:        testSteps,
		Configurations: []stackstate_api.StackPackConfiguration{
			{
				Id:                  &testId,
				LastUpdateTimestamp: &testLut,
				Status:              testStatus,
				StackPackVersion:    testCurrVersion,
			},
		},
		NextVersion: &stackstate_api.FullStackPack{
			Version: testLatestVersion,
		},
		LatestVersion: &stackstate_api.FullStackPack{
			Version: testLatestVersion,
		},
	}
)

func TestStackpackDescribeNotFound(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDescribeCommand(&cli.Deps)

	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = mockResponse

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "describe", "--name", "foobar")
	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "describe", "--name", "foobar", "-o", "json")

	expected := []error{
		fmt.Errorf("StackPack 'foobar' not found."),
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintErrCalls)
	assert.Equal(t, expected, *cli.MockPrinter.PrintErrJsonCalls)
}

func TestStackpackDescribeJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDescribeCommand(&cli.Deps)

	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = mockResponse

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "describe", "--name", awsName, "-o", "json")

	expected := []map[string]interface{}{
		{
			"stackpack": awsStackPack,
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.PrintJsonCalls)
}

func TestStackpackDescribeText(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDescribeCommand(&cli.Deps)

	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = []stackstate_api.FullStackPack{testStackPack}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "describe", "--name", testName)

	expectedLines := []string{
		fmt.Sprintf("StackPack %s v%s (latest available is v%s)", testDisplayName, testCurrVersion, testLatestVersion),
		fmt.Sprintf("Categories: %s, %s", testCat1, testCat2),
		fmt.Sprintf("Provided integrations: %s, %s", testCat1, testCat2),
		"\nInstance configuration parameters:",
		"\nInstalled instances:",
	}

	expectedTables := []printer.TableData{
		{
			Header: []string{"name", "display name", "type"},
			Data: [][]interface{}{
				{&testStep1Name, &testStep1Name, &testStepType},
				{&testStep2Name, &testStep2Name, &testStepType},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "instance parameters"},
		},
		{
			Header: []string{"id", "status", "version", "last updated"},
			Data: [][]interface{}{
				{&testId, testStatus, testCurrVersion, testTimestamp},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "installed StackPack instances"},
		},
	}

	assert.Equal(t, expectedLines, *cli.MockPrinter.PrintLnCalls)
	assert.Equal(t, expectedTables, *cli.MockPrinter.TableCalls)
}

func TestStackpackDescribeTextNoCategories(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDescribeCommand(&cli.Deps)

	newTest := testStackPack
	newTest.Categories = []string{}

	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = []stackstate_api.FullStackPack{newTest}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "describe", "--name", testName)

	expectedLines := []string{
		fmt.Sprintf("StackPack %s v%s (latest available is v%s)", testDisplayName, testCurrVersion, testLatestVersion),
		fmt.Sprintf("Provided integrations: %s, %s", testCat1, testCat2),
		"\nInstance configuration parameters:",
		"\nInstalled instances:",
	}

	assert.Equal(t, expectedLines, *cli.MockPrinter.PrintLnCalls)
}

func TestStackpackDescribeTextNoSteps(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDescribeCommand(&cli.Deps)

	newTest := testStackPack
	newTest.Steps = []stackstate_api.StackPackStep{}

	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = []stackstate_api.FullStackPack{newTest}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "describe", "--name", testName)

	expectedTables := []printer.TableData{
		{
			Header: []string{"id", "status", "version", "last updated"},
			Data: [][]interface{}{
				{&testId, testStatus, testCurrVersion, testTimestamp},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "installed StackPack instances"},
		},
	}

	assert.Equal(t, expectedTables, *cli.MockPrinter.TableCalls)
}

func TestStackpackDescribeTextNoInstances(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := StackpackDescribeCommand(&cli.Deps)

	newTest := testStackPack
	newTest.Configurations = []stackstate_api.StackPackConfiguration{}

	cli.MockClient.ApiMocks.StackpackApi.StackPackListResponse.Result = []stackstate_api.FullStackPack{newTest}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "describe", "--name", testName)

	expectedTables := []printer.TableData{
		{
			Header: []string{"name", "display name", "type"},
			Data: [][]interface{}{
				{&testStep1Name, &testStep1Name, &testStepType},
				{&testStep2Name, &testStep2Name, &testStepType},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "instance parameters"},
		},
		// NOTE Still printed.
		{
			Header:              []string{"id", "status", "version", "last updated"},
			Data:                [][]interface{}{},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "installed StackPack instances"},
		},
	}
	assert.Equal(t, expectedTables, *cli.MockPrinter.TableCalls)
}
