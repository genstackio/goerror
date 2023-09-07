package errors

type MarshallError struct {
	Err error `json:"err,omitempty"`
}

func (err MarshallError) Error() string {
	return "Marshall failure"
}

func (err MarshallError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "marshall",
		Message:    err.Error(),
		Detail:     err.Err.Error(),
		Code:       12200,
		StatusCode: 500,
	}
}
