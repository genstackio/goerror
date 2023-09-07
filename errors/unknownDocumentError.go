package errors

type UnknownDocumentError struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"`
	Type  string `json:"type,omitempty"`
	Err   error  `json:"err,omitempty"`
}

func (err UnknownDocumentError) Error() string {
	return "Unknown " + err.Type + " with " + err.Key + " '" + err.Value + "'"
}

func (err UnknownDocumentError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m[err.Key] = err.Value
	m["type"] = err.Type
	m["key"] = err.Key
	m["value"] = err.Value
	return m
}

func (err UnknownDocumentError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "unknown_document",
		Message:    err.Error(),
		Detail:     err.Err.Error(),
		Code:       12300,
		StatusCode: 404,
		Data:       err.GetData(),
	}
}
