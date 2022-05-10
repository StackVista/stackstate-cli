package anomaly

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/stretchr/testify/assert"
	"gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func setupCommandCollect() (di.MockDeps, *cobra.Command) {
	mockCli := di.NewMockDeps()
	cmd := AnomalyCollect(&mockCli.Deps)

	return mockCli, cmd
}

type TestClock struct {
	now int64
}

func newTestClock() Clock {
	return &TestClock{
		now: 1652108645000,
	}
}

func (clock *TestClock) Now() time.Time {
	return time.UnixMilli(clock.now)
}

func TestTimeParsing(t *testing.T) {
	clock := newTestClock()
	parsedTime, err := parseTime(clock, "-2h")
	assert.Nil(t, err)
	assert.Equal(t, int64(1652108645000-7200000), parsedTime.UnixMilli())

	parsedTime, err = parseTime(clock, "2022-05-09T15:04:05Z")
	assert.Nil(t, err)
	assert.Equal(t, int64(1652108645000), parsedTime.UnixMilli())

	parsedTime, err = parseTime(clock, "1652108645000")
	assert.Nil(t, err)
	assert.Equal(t, int64(1652108645000), parsedTime.UnixMilli())

	_, err = parseTime(clock, "-2d")
	assert.NotNil(t, err)
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
	var feedback []stackstate_api.FeedbackWithContext
	err = json.Unmarshal([]byte(expectedStr), &feedback)
	assert.Nil(t, err)

	cli, cmd := setupCommandCollect()
	cli.MockClient.ApiMocks.AnomalyFeedbackApi.CollectAnomalyFeedbackResponse.Result = feedback

	_, err = di.ExecuteCommandWithContext(&cli.Deps, cmd, "--start-time", "1652108645000", "--end-time", "1652108845000", "--file", filePath)
	assert.Nil(t, err)
	assert.Equal(t, []string{"Downloaded 2 anomalies with feedback."}, *cli.MockPrinter.SuccessCalls)

	calls := *cli.MockClient.ApiMocks.AnomalyFeedbackApi.CollectAnomalyFeedbackCalls
	assert.Equal(t, 1, len(calls))
	assert.Equal(t, int64(1652108645000), *calls[0].PstartTime)
	assert.Equal(t, int64(1652108845000), *calls[0].PendTime)
	assert.Equal(t, int64(86400000), *calls[0].Phistory)

	body, err := ioutil.ReadFile(filePath)
	assert.Nil(t, err)
	var writtenFeedback []stackstate_api.FeedbackWithContext
	err = json.Unmarshal(body, &writtenFeedback)
	assert.Nil(t, err)
	assert.Equal(t, feedback, writtenFeedback)
}
