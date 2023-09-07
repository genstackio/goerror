package errors

type DocumentCreationError struct {
	Params map[string]interface{} `json:"params,omitempty"`
	Type   string                 `json:"type,omitempty"`
	Err    error                  `json:"err,omitempty"`
}

func (err DocumentCreationError) Error() string {
	return "Unable to create " + err.Type + ""
}

func (err DocumentCreationError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m["type"] = err.Type
	params := map[string]interface{}{}
	for k, v := range err.Params {
		params[k] = v
	}
	m["params"] = params
	return m
}

func (err DocumentCreationError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "document_creation",
		Message:    err.Error(),
		Detail:     err.Err.Error(),
		Code:       12800,
		StatusCode: 500,
		Data:       err.GetData(),
	}
}
