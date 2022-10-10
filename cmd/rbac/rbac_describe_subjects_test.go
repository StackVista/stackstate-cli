package rbac

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/printer"
)

var (
	SomeScopeVar   = SomeScope
	SubjectConfig1 = stackstate_api.SubjectConfig{
		Handle:     SomeSubject,
		ScopeQuery: &SomeScopeVar,
	}

	SomeOtherSubject = "handle"
	SomeOtherScope   = "meaningOfLife = 23"

	SubjectConfig2 = stackstate_api.SubjectConfig{
		Handle:     SomeOtherSubject,
		ScopeQuery: &SomeOtherScope,
	}

	SubjectConfig3 = stackstate_api.SubjectConfig{
		Handle:     SubjectHandle,
		ScopeQuery: nil,
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
			Header: []string{"Subject", "Scope Query"},
			Data: [][]interface{}{
				{SubjectConfig1.Handle, *SubjectConfig1.ScopeQuery},
				{SubjectConfig2.Handle, *SubjectConfig2.ScopeQuery},
				{SubjectConfig3.Handle, ""},
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
			Header: []string{"Subject", "Scope Query"},
			Data: [][]interface{}{
				{SubjectConfig1.Handle, *SubjectConfig1.ScopeQuery},
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
			"handle":     SubjectConfig1.Handle,
			"scopeQuery": *SubjectConfig1.ScopeQuery,
		},
		{
			"handle":     SubjectConfig3.Handle,
			"scopeQuery": "",
		},
	}

	assert.Equal(t, expectedJson, *cli.MockPrinter.PrintJsonCalls)
}
