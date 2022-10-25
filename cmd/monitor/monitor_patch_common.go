package monitor

import (
	"net/http"

	sts "github.com/stackvista/stackstate-cli/generated/stackstate_api"
	"github.com/stackvista/stackstate-cli/internal/di"
)

func RunMonitorStatusPatch(cli *di.Deps, api *sts.APIClient, status *sts.MonitorStatusValue, identifier string) (*sts.Monitor, *http.Response, error) {
	monitorPatch := sts.MonitorPatch{
		Status: status,
	}

	return api.MonitorApi.PatchMonitor(cli.Context, identifier).MonitorPatch(monitorPatch).Execute()
}
