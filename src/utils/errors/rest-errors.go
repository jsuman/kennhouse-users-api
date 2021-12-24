package errors

import "net/http"

type RestErr struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
	Error   string `json:"error"`
}

func BadRequestError(err string) *RestErr {
	return &RestErr{
		Message: err,
		Status:  http.StatusBadRequest,
		Error:   "bad_request",
	}
}

func NotFoundError(err string) *RestErr {
	return &RestErr{
		Message: err,
		Status:  http.StatusNotFound,
		Error:   "not_found",
	}
}

func InternalServerError(err string) *RestErr {
	return &RestErr{
		Message: err,
		Status:  http.StatusInternalServerError,
		Error:   "internal_server_errpr",
	}
}
