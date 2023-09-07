package errors

type UnsupportedValueError struct {
	Type  string `json:"type,omitempty"`
	Value string `json:"value,omitempty"`
}

func (err UnsupportedValueError) Error() string {
	return "Unsupported value '" + err.Value + "' for " + err.Type
}

func (err UnsupportedValueError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m["type"] = err.Type
	m["value"] = err.Value
	return m
}

func (err UnsupportedValueError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "unsupported_value",
		Message:    err.Error(),
		Code:       12150,
		StatusCode: 400,
		Data:       err.GetData(),
	}
}
