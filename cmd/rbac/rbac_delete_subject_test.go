package rbac

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func TestDeleteSubjectJson(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DeleteSubjectCommand(&cli.Deps)

	cli.MockClient.ApiMocks.SubjectAPI.DeleteSubjectResponse.Response = &http.Response{Body: io.NopCloser(strings.NewReader(""))}

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SomeSubject, "-o", "json")

	calls := *cli.MockClient.ApiMocks.SubjectAPI.DeleteSubjectCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, SomeSubject, calls[0].Psubject)

	expectedJson := []map[string]interface{}{
		{
			"deleted-subject": SomeSubject,
		},
	}

	assert.Equal(t, expectedJson, *cli.MockPrinter.PrintJsonCalls)
}

func TestDeleteSubject(t *testing.T) {
	cli := di.NewMockDeps(t)
	cmd := DeleteSubjectCommand(&cli.Deps)

	di.ExecuteCommandWithContextUnsafe(&cli.Deps, cmd, "--subject", SomeOtherSubject)

	calls := *cli.MockClient.ApiMocks.SubjectAPI.DeleteSubjectCalls
	assert.Len(t, calls, 1)
	assert.Equal(t, SomeOtherSubject, calls[0].Psubject)

	expectedStrings := []string{
		fmt.Sprintf("Deleted subject '%s'", SomeOtherSubject),
	}

	assert.Equal(t, expectedStrings, *cli.MockPrinter.SuccessCalls)
}
