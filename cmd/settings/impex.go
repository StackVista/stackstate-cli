package settings

import (
	"context"
	"net/http"

	"github.com/stackvista/stackstate-cli/generated/stackstate_api"
)

func DoExport(ctx context.Context, api *stackstate_api.APIClient, ids []int64, namespace string, nodeTypes []string, allowReferences []string) (string, *http.Response, error) {
	exportArgs := stackstate_api.NewExport()

	if len(ids) != 0 {
		exportArgs.NodesWithIds = ids
	}

	if len(namespace) != 0 {
		exportArgs.Namespace = &namespace
	}

	if len(nodeTypes) != 0 {
		exportArgs.AllNodesOfTypes = nodeTypes
	}

	if len(allowReferences) != 0 {
		exportArgs.AllowReferences = allowReferences
	}

	return api.ExportAPI.ExportSettings(ctx).Export(*exportArgs).Execute()
}

func doImport(ctx context.Context, api *stackstate_api.APIClient, body string, namespace string, unlockedStrategy string, timeout int64) ([]map[string]interface{}, *http.Response, error) {
	request := api.ImportAPI.ImportSettings(ctx).Body(body)
	if namespace != "" {
		request = request.Namespace(namespace)
	}
	if unlockedStrategy != "" {
		request = request.Unlocked(unlockedStrategy)
	}
	if timeout > 0 {
		request = request.TimeoutSeconds(timeout)
	}
	return request.Execute()
}
