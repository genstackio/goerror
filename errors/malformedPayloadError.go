package errors

type MalformedPayloadError struct {
	Err error `json:"err,omitempty"`
}

func (err MalformedPayloadError) Error() string {
	return "Malformed payload"
}

func (err MalformedPayloadError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "malformed_payload",
		Message:    err.Error(),
		Detail:     err.Err.Error(),
		Code:       12100,
		StatusCode: 400,
	}
}
