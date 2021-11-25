package openapi_util

import (
	"fmt"
	"net/http"

	sts "gitlab.com/stackvista/stackstate-cli2/internal/stackstate_client"
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
		bodyStr := string(v.Body())
		if bodyStr != "" {
			bodyStr = fmt.Sprintf("Server response: %s", bodyStr)
		}
		return fmt.Errorf("%s%s%s", status, bodyStr, suggestion)
	default:
		return fmt.Errorf("%v%v%s", status, v, suggestion)
	}
}
