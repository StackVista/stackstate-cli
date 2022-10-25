package anomaly

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
	"github.com/stretchr/testify/assert"
)

func setupCommandCollect(t *testing.T) (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps(t)
	cmd := AnomalyCollectFeedback(&mockCli.Deps)

	return mockCli, cmd
}

func TestAnomalyCollectEmpty(t *testing.T) {
	file, err := os.CreateTemp(os.TempDir(), "test_")
	if err != nil {
		panic(err)
	}
	filePath := file.Name()
	file.Close()

	referenceFeedback, err := os.ReadFile("test_feedback.json")
	assert.Nil(t, err)
	expectedStr := string(referenceFeedback)
	var feedback []stackstate_api.AnomalyWithContext
	err = json.Unmarshal([]byte(expectedStr), &feedback)
	assert.Nil(t, err)

	cli, cmd := setupCommandCollect(t)
	cli.MockClient.ApiMocks.AnomalyFeedbackApi.ExportAnomalyResponse.Result = feedback

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--start-time", "1652108645000", "--end-time", "1652108845000", "--file", filePath)
	assert.Nil(t, err)
	assert.Equal(t, []string{fmt.Sprintf("Collected 2 anomalies to %s.", filePath)}, *cli.MockPrinter.SuccessCalls)

	calls := *cli.MockClient.ApiMocks.AnomalyFeedbackApi.ExportAnomalyCalls
	assert.Equal(t, 1, len(calls))
	assert.Equal(t, "present", *calls[0].Pfeedback)
	assert.Equal(t, int64(1652108645000), *calls[0].PstartTime)
	assert.Equal(t, int64(1652108845000), *calls[0].PendTime)
	assert.Equal(t, int64(86400000), *calls[0].Phistory)

	body, err := os.ReadFile(filePath)
	assert.Nil(t, err)
	var writtenFeedback []stackstate_api.AnomalyWithContext
	err = json.Unmarshal(body, &writtenFeedback)
	assert.Nil(t, err)
	assert.Equal(t, feedback, writtenFeedback)
}
