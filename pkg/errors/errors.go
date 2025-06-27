package errors

import (
	"github.com/emicklei/go-restful"
)

type RequestError struct {
	StatusCode int    `json:"statusCode" validate:"required"`
	ErrMessage string `json:"errMessage" validate:"required"`
}

func (r *RequestError) Error() string {
	return r.ErrMessage
}

func (r *RequestError) Code() int {
	return r.StatusCode
}

func (r *RequestError) APIErrorHandler(request *restful.Request, response *restful.Response) {
	_ = response.WriteError(r.StatusCode, r)
}
