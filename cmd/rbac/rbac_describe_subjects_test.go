package rbac

import (
	"strings"
	"testing"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stackvista/stackstate-cli/internal/printer"
	"github.com/stretchr/testify/assert"
)

var (
	SomeScopeVar   = SomeScope
	SubjectConfig1 = stackstate_api.SubjectConfig{
		Handle: SomeSubject,
		Source: stackstate_api.SUBJECTSOURCE_OBSERVABILITY,
	}

	SomeOtherSubject = "handle"

	SubjectConfig2 = stackstate_api.SubjectConfig{
		Handle: SomeOtherSubject,
		Source: stackstate_api.SUBJECTSOURCE_STATIC,
	}

	SubjectConfig3 = stackstate_api.SubjectConfig{
		Handle: SubjectHandle,
		Source: stackstate_api.SUBJECTSOURCE_STATIC,
	}
)

func TestDescribeSubjectsTable(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeSubjectsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.SubjectApi.ListSubjectsResponse.Result = []stackstate_api.SubjectConfig{
		SubjectConfig1,
		SubjectConfig2,
		SubjectConfig3,
	}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	calls := *cli.MockClient.ApiMocks.SubjectApi.ListSubjectsCalls
	assert.Len(t, calls, 1)

	expected := []printer.TableData{
		{
			Header: []string{"Subject", "Source"},
			Data: [][]interface{}{
				{SubjectConfig1.Handle, SubjectConfig1.Source},
				{SubjectConfig2.Handle, SubjectConfig2.Source},
				{SubjectConfig3.Handle, SubjectConfig3.Source},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "subjects"},
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.TableCalls)
}

func TestDescribeSubjectsJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeSubjectsCommand(&cli.Deps)

	result := []stackstate_api.SubjectConfig{
		SubjectConfig1,
		SubjectConfig2,
		SubjectConfig3,
	}

	cli.MockClient.ApiMocks.SubjectApi.ListSubjectsResponse.Result = result

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "-o", "json")

	calls := *cli.MockClient.ApiMocks.SubjectApi.ListSubjectsCalls
	assert.Len(t, calls, 1)

	expectedJson := []map[string]interface{}{
		{
			"subjects": result,
		},
	}

	assert.Equal(t, expectedJson, *cli.MockPrinter.PrintJsonCalls)
}

func TestDescribeSubjectsTableWithFilter(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeSubjectsCommand(&cli.Deps)

	result := []stackstate_api.SubjectConfig{
		SubjectConfig1,
		SubjectConfig2,
		SubjectConfig3,
	}

	cli.MockClient.ApiMocks.SubjectApi.ListSubjectsResponse.Result = result

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SomeSubject)

	calls := *cli.MockClient.ApiMocks.SubjectApi.ListSubjectsCalls
	assert.Len(t, calls, 1)

	expected := []printer.TableData{
		{
			Header: []string{"Subject", "Source"},
			Data: [][]interface{}{
				{SubjectConfig1.Handle, SubjectConfig1.Source},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "subjects"},
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.TableCalls)
}

func TestDescribeSubjectsJsonWithFilter(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeSubjectsCommand(&cli.Deps)

	result := []stackstate_api.SubjectConfig{
		SubjectConfig1,
		SubjectConfig2,
		SubjectConfig3,
	}

	cli.MockClient.ApiMocks.SubjectApi.ListSubjectsResponse.Result = result

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SubjectConfig1.Handle, "-o", "json")

	calls := *cli.MockClient.ApiMocks.SubjectApi.ListSubjectsCalls
	assert.Len(t, calls, 1)

	expectedJson := []map[string]interface{}{
		{
			"subjects": []stackstate_api.SubjectConfig{SubjectConfig1},
		},
	}

	assert.Equal(t, expectedJson, *cli.MockPrinter.PrintJsonCalls)
}

func TestDescribeSubjectsTableWithSourceFilter(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeSubjectsCommand(&cli.Deps)

	result := []stackstate_api.SubjectConfig{
		SubjectConfig1,
		SubjectConfig2,
		SubjectConfig3,
	}

	cli.MockClient.ApiMocks.SubjectApi.ListSubjectsResponse.Result = result

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--source", strings.ToLower(string(SubjectConfig1.Source)))

	calls := *cli.MockClient.ApiMocks.SubjectApi.ListSubjectsCalls
	assert.Len(t, calls, 1)

	expected := []printer.TableData{
		{
			Header: []string{"Subject", "Source"},
			Data: [][]interface{}{
				{SubjectConfig1.Handle, SubjectConfig1.Source},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "subjects"},
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.TableCalls)
}

func TestProduceEmptyWhenNoSubjectFilter(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeSubjectsCommand(&cli.Deps)

	var result []stackstate_api.SubjectConfig

	cli.MockClient.ApiMocks.SubjectApi.ListSubjectsResponse.Result = result

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd)

	calls := *cli.MockClient.ApiMocks.SubjectApi.ListSubjectsCalls
	assert.Len(t, calls, 1)

	expected := []printer.TableData{
		{
			Header:              []string{"Subject", "Source"},
			Data:                [][]interface{}{},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "subjects"},
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.TableCalls)
}

func TestProduceNotFoundWithSubjectFilter(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeSubjectsCommand(&cli.Deps)

	var result []stackstate_api.SubjectConfig

	cli.MockClient.ApiMocks.SubjectApi.ListSubjectsResponse.Result = result

	_, err := di.ExecuteCommandWithContext(&cli.Deps, cmd, "--subject", "bladibla")

	assert.EqualError(t, err, "could not find subject: 'bladibla'")
	calls := *cli.MockClient.ApiMocks.SubjectApi.ListSubjectsCalls
	assert.Len(t, calls, 1)

}
