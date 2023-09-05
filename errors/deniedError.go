package errors

type DeniedError struct {
	Reason string
}

func (err DeniedError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "denied",
		Message:    err.Reason,
		Detail:     "",
		Code:       12997,
		StatusCode: 403,
	}
}
