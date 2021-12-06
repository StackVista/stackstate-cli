package common

import (
	"net/http"
)

type ResponseError struct {
	Err  error
	Resp *http.Response
}

func (p ResponseError) Error() string {
	return p.Err.Error()
}

func NewResponseError(err error, resp *http.Response) ResponseError {
	return ResponseError{
		Err:  err,
		Resp: resp,
	}
}
