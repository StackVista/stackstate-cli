package openapi_util

import (
	"encoding/json"
	"fmt"
	"net/http"

	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
	"gopkg.in/yaml.v3"
)

func MakeErrorFromResponse(err error, resp *http.Response) error {
	var status string
	if resp != nil && resp.Status != "" {
		status = resp.Status + ". "
	}

	suggestion := ""
	if resp != nil && resp.StatusCode == 401 {
		suggestion = "\nPlease check your configured API token."
	}

	switch v := err.(type) {
	case sts.GenericOpenAPIError:
		var bodyStr string
		if resp.Header.Get("Content-Type") == "application/json" {
			var bodyStruct interface{}
			json.Unmarshal(v.Body(), &bodyStruct)
			yaml, err := yaml.Marshal(bodyStruct)
			if err == nil && yaml != nil && bodyStruct != nil {
				yamlResp := string(yaml)
				bodyStr = fmt.Sprintf(
					"\n\n----------------\n"+
						"Server response:\n"+
						"----------------\n"+
						"%s", yamlResp)
			}
		}
		return fmt.Errorf("%s%s%s", status, bodyStr, suggestion)
	default:
		return fmt.Errorf("%v%v%s", status, v, suggestion)
	}
}
