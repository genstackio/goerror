package errors

type AccessDeniedError struct {
	Err error `json:"err,omitempty"`
}

func (err AccessDeniedError) Error() string {
	return err.Err.Error()
}

func (err AccessDeniedError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "denied",
		Message:    err.Error(),
		Detail:     err.Err.Error(),
		Code:       12993,
		StatusCode: 403,
	}
}
