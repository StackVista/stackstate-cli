package anomaly

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
	"gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gitlab.com/stackvista/stackstate-cli2/internal/util"
)

func setupCommandCollect() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := AnomalyCollect(&mockCli.Deps)

	return mockCli, cmd
}

func TestAnomalyCollectEmpty(t *testing.T) {
	file, err := ioutil.TempFile(os.TempDir(), "test_")
	if err != nil {
		panic(err)
	}
	filePath := file.Name()
	file.Close()

	referenceFeedback, err := ioutil.ReadFile("test_feedback.json")
	assert.Nil(t, err)
	expectedStr := string(referenceFeedback)
	var feedback []stackstate_client.FeedbackWithContext
	err = json.Unmarshal([]byte(expectedStr), &feedback)
	assert.Nil(t, err)

	cli, cmd := setupCommandCollect()
	cli.MockClient.ApiMocks.AnomalyFeedbackApi.CollectAnomalyFeedbackResponse.Result = feedback

	_, err = util.ExecuteCommandWithContext(cli.Context, cmd, "--start-time=-1h", "--file", filePath)
	assert.Nil(t, err)
	assert.Equal(t, []string{}, *cli.MockPrinter.PrintLnCalls)

	body, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	var writtenFeedback []stackstate_client.FeedbackWithContext
	err = json.Unmarshal(body, &writtenFeedback)
	assert.Nil(t, err)
	assert.Equal(t, feedback, writtenFeedback)
}
