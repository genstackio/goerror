package errors

type DocumentUpdateError struct {
	Key    string                 `json:"key,omitempty"`
	Value  string                 `json:"value,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
	Type   string                 `json:"type,omitempty"`
	Err    error                  `json:"err,omitempty"`
}

func (err DocumentUpdateError) Error() string {
	return "Unable to update " + err.Type + " with " + err.Key + " '" + err.Value + "'"
}

func (err DocumentUpdateError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m["type"] = err.Type
	m[err.Key] = err.Value
	m["key"] = err.Key
	m["value"] = err.Value
	return m
}

func (err DocumentUpdateError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "document_update",
		Message:    err.Error(),
		Detail:     err.Err.Error(),
		Code:       12810,
		StatusCode: 500,
		Data:       err.GetData(),
	}
}
