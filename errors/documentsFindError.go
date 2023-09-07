package errors

type DocumentsFindError struct {
	Type   string                 `json:"type,omitempty"`
	Err    error                  `json:"err,omitempty"`
	Params map[string]interface{} `json:"params,omitempty"`
}

func (err DocumentsFindError) Error() string {
	return "Unable to find " + err.Type + " documents: " + err.Err.Error()
}

func (err DocumentsFindError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m["type"] = err.Type
	params := map[string]interface{}{}
	for k, v := range err.Params {
		params[k] = v
	}
	m["params"] = params
	return m
}

func (err DocumentsFindError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "document_retrieve",
		Message:    err.Error(),
		Detail:     err.Err.Error(),
		Code:       12800,
		StatusCode: 500,
		Data:       err.GetData(),
	}
}
