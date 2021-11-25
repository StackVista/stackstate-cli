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

	switch v := err.(type) {
	case sts.GenericOpenAPIError:
		return fmt.Errorf("%vError response: %+v", status, string(v.Body()))
	default:
		return fmt.Errorf("%v%+v", status, v)
	}
}
