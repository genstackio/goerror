package errors

type GenericError struct {
	Params map[string]interface{} `json:"params,omitempty"`
	Err    error                  `json:"err,omitempty"`
}

func (err GenericError) Error() string {
	return "Unexpected error"
}

func (err GenericError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	for k, v := range err.Params {
		m[k] = v
	}
	return m
}

func (err GenericError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "generic",
		Message:    err.Error(),
		Detail:     err.Err.Error(),
		Code:       12200,
		StatusCode: 500,
		Data:       err.GetData(),
	}
}
