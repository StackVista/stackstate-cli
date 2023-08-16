package client

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/common"
	"github.com/stretchr/testify/assert"
)

func setupClient() (StackStateClient, *stackstate_api.ServerApiMock) {
	serverApi := stackstate_api.NewServerApiMock()
	api := stackstate_api.APIClient{
		ServerApi: &serverApi,
	}
	client := StdStackStateClient{
		client:  &api,
		Context: context.Background(),
		apiURL:  "https://test",
	}
	return client, &serverApi
}

func TestConnectError(t *testing.T) {
	client, serverApi := setupClient()
	err := fmt.Errorf("authentication error")
	resp := http.Response{StatusCode: 401}
	serverApi.ServerInfoResponse.Response = &resp
	serverApi.ServerInfoResponse.Error = err

	_, _, actualErr := client.Connect()

	assert.Equal(t, common.NewConnectError(err, "https://test", &resp), actualErr)
}

func TestConnectSuccess(t *testing.T) {
	client, serverApi := setupClient()
	serverApi.ServerInfoResponse.Result = stackstate_api.ServerInfo{DeploymentMode: "test"}

	_, serverInfo, _ := client.Connect()

	assert.Equal(t, "test", serverInfo.DeploymentMode)
}

func TestCombineURLAndApi(t *testing.T) {
	assert.Equal(t, combineURLandPath("https://bla", "api"), "https://bla/api")
	assert.Equal(t, combineURLandPath("https://bla///", "///api"), "https://bla/api")
	assert.Equal(t, combineURLandPath("https://bla", "/api"), "https://bla/api")
	assert.Equal(t, combineURLandPath("/https://bla", "api/"), "https://bla/api")
}

func TestVersionToString(t *testing.T) {
	assert.Equal(t, "1.2.3", VersionToString(stackstate_api.ServerVersion{Major: 1, Minor: 2, Patch: 3}))
	assert.Equal(t, "1.2.3+c0ffee23", VersionToString(stackstate_api.ServerVersion{Major: 1, Minor: 2, Patch: 3, Diff: "c0ffee23"}))
	assert.Equal(t, "1.2.3+c0ffee23-yes", VersionToString(stackstate_api.ServerVersion{Major: 1, Minor: 2, Patch: 3, Diff: "c0ffee23", Commit: "yes"}))
	assert.Equal(t, "1.2.3+c0ffee23-yes-dev", VersionToString(stackstate_api.ServerVersion{Major: 1, Minor: 2, Patch: 3, Diff: "c0ffee23", Commit: "yes", IsDev: true}))
	assert.Equal(t, "1.2.3-dev", VersionToString(stackstate_api.ServerVersion{Major: 1, Minor: 2, Patch: 3, IsDev: true}))
	assert.Equal(t, "1.2.3-yes-dev", VersionToString(stackstate_api.ServerVersion{Major: 1, Minor: 2, Patch: 3, Commit: "yes", IsDev: true}))
	assert.Equal(t, "1.2.3-yes", VersionToString(stackstate_api.ServerVersion{Major: 1, Minor: 2, Patch: 3, Commit: "yes"}))
}

type VersionTest struct {
	Server stackstate_api.ServerVersion
	Cmd    string
	Result common.CLIError
}

func V(major int, minor int, patch int) stackstate_api.ServerVersion {
	v := stackstate_api.NewServerVersion(int32(major), int32(minor), int32(patch), "", "", false)
	return *v
}

func TestVersionCompatibilityCheck(t *testing.T) {
	tests := []VersionTest{
		// Empty minimum version is fine.
		{Server: V(1, 2, 3), Cmd: "", Result: nil},
		// Compatible versions are fine.
		{Server: V(1, 2, 3), Cmd: "1.0.0", Result: nil},
		{Server: V(1, 2, 3), Cmd: "1.0.6", Result: nil},
		{Server: V(1, 2, 3), Cmd: "1.2.0", Result: nil},
		{Server: V(1, 2, 3), Cmd: "1.2.3", Result: nil},

		// Server version really only needs to be greater than or equal to the command minimum version, even newer major versions are fine
		{Server: V(2, 2, 3), Cmd: "1.2.4", Result: nil},

		// Weird versions are fine(?).
		{Server: V(1, 2, 3), Cmd: "1.2.2-alpha.preview+123", Result: nil},
		{
			Server: stackstate_api.ServerVersion{Major: 1, Minor: 2, Patch: 3, Diff: "diff", Commit: "deadbeef"},
			Cmd:    "1.2.3",
			// FIXME Should be NewAPIVersionMismatchError("1.2.3+diff-deadbeef", "1.2.3"),
			Result: nil,
		},
		// Incompatible versions are detected.
		{Server: V(1, 2, 3), Cmd: "1.2.3-alpha.preview+123", Result: nil},
		{Server: V(1, 2, 3), Cmd: "1.2.4", Result: NewAPIVersionMismatchError("1.2.3", "1.2.4")},
		{Server: V(1, 2, 3), Cmd: "1.3.3", Result: NewAPIVersionMismatchError("1.2.3", "1.3.3")},
		{Server: V(1, 2, 3), Cmd: "2.2.4", Result: NewAPIVersionMismatchError("1.2.3", "2.2.4")},
		// Pre-release versions are less than the released version, so they are not sufficient
		{
			Server: stackstate_api.ServerVersion{Major: 1, Minor: 2, Patch: 3, Commit: "deadbeef"},
			Cmd:    "1.2.3",
			Result: NewAPIVersionMismatchError("1.2.3-deadbeef", "1.2.3"),
		},
		{
			Server: stackstate_api.ServerVersion{Major: 1, Minor: 2, Patch: 3, IsDev: true},
			Cmd:    "1.2.3",
			Result: NewAPIVersionMismatchError("1.2.3-dev", "1.2.3"),
		},
		// Invalid versions are detected.
		{
			Server: V(1, 2, 3),
			Cmd:    "notAVersion",
			Result: common.NewAPIVersionError(fmt.Errorf("No Major.Minor.Patch elements found")),
		},
		{
			Server: V(-23, -5, 1337),
			Cmd:    "23.5.1337",
			Result: common.NewAPIVersionError(fmt.Errorf("Invalid character(s) found in major number \"-23\"")),
		},
	}

	for _, test := range tests {
		assert.Equal(t, test.Result, CheckVersionCompatibility(test.Server, test.Cmd))
	}
}
