package rbac

import (
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
	}

	SomeOtherSubject = "handle"

	SubjectConfig2 = stackstate_api.SubjectConfig{
		Handle: SomeOtherSubject,
	}

	SubjectConfig3 = stackstate_api.SubjectConfig{
		Handle: SubjectHandle,
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
			Header: []string{"Subject"},
			Data: [][]interface{}{
				{SubjectConfig1.Handle},
				{SubjectConfig2.Handle},
				{SubjectConfig3.Handle},
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

	cli.MockClient.ApiMocks.SubjectApi.GetSubjectResponse.Result = SubjectConfig1

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SomeSubject)

	calls := *cli.MockClient.ApiMocks.SubjectApi.GetSubjectCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, SomeSubject, calls[0].Psubject)

	expected := []printer.TableData{
		{
			Header: []string{"Subject"},
			Data: [][]interface{}{
				{SubjectConfig1.Handle},
			},
			MissingTableDataMsg: printer.NotFoundMsg{Types: "matching subjects"},
		},
	}

	assert.Equal(t, expected, *cli.MockPrinter.TableCalls)
}

func TestDescribeSubjectsJsonWithFilter(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DescribeSubjectsCommand(&cli.Deps)

	cli.MockClient.ApiMocks.SubjectApi.GetSubjectResponse.Result = SubjectConfig1

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SomeSubject, "-o", "json")

	cli.MockClient.ApiMocks.SubjectApi.GetSubjectResponse.Result = SubjectConfig3

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SomeOtherSubject, "-o", "json")

	calls := *cli.MockClient.ApiMocks.SubjectApi.GetSubjectCalls
	assert.Len(t, calls, 2)
	assert.Equal(t, SomeSubject, calls[0].Psubject)
	assert.Equal(t, SomeOtherSubject, calls[1].Psubject)

	expectedJson := []map[string]interface{}{
		{
			"handle": SubjectConfig1.Handle,
		},
		{
			"handle": SubjectConfig3.Handle,
		},
	}

	assert.Equal(t, expectedJson, *cli.MockPrinter.PrintJsonCalls)
}
