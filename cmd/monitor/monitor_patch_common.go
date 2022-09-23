package monitor

import (
	"net/http"

	sts "gitlab.com/stackvista/stackstate-cli2/generated/stackstate_api"
	"gitlab.com/stackvista/stackstate-cli2/internal/di"
)

func RunMonitorStatusPatch(cli *di.Deps, api *sts.APIClient, status *sts.MonitorStatusValue, identifier string) (*sts.Monitor, *http.Response, error) {
	monitorPatch := sts.MonitorPatch{
		Status: status,
	}

	return api.MonitorApi.PatchMonitor(cli.Context, identifier).MonitorPatch(monitorPatch).Execute()
}
