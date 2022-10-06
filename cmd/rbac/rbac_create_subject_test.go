package rbac

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

const (
	SomeSubject = "subject"
	SomeScope   = "withNeighboursOf(components = (id = '*'), levels = '14534564576475675671', direction = 'both')"
)

func TestCreateSubject(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := CreateSubjectCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SomeSubject, "-o", "json")

	calls := *cli.MockClient.ApiMocks.SubjectApi.CreateSubjectCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, SomeSubject, calls[0].Psubject)

	expectedSubject := stackstate_api.NewCreateSubject(DefaultScope, DefaultSTQLVersion)

	assert.Equal(t, expectedSubject, calls[0].PcreateSubject)

	expectedJson := []map[string]interface{}{
		{
			"create-subject": SomeSubject,
		},
	}

	assert.Equal(t, expectedJson, *cli.MockPrinter.PrintJsonCalls)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SomeSubject, "--scope", SomeScope)

	calls = *cli.MockClient.ApiMocks.SubjectApi.CreateSubjectCalls
	assert.Len(t, calls, 2)
	assert.Equal(t, SomeSubject, calls[0].Psubject)

	otherExpectedSubject := stackstate_api.NewCreateSubject(SomeScope, DefaultSTQLVersion)

	assert.Equal(t, otherExpectedSubject, calls[1].PcreateSubject)
}
