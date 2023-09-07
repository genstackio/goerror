package errors

type WriteError struct {
	Err error `json:"err,omitempty"`
}

func (err WriteError) Error() string {
	return "Write failure"
}

func (err WriteError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "write",
		Message:    err.Error(),
		Detail:     err.Err.Error(),
		Code:       12900,
		StatusCode: 500,
	}
}
