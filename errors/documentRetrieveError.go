package errors

type DocumentRetrieveError struct {
	Type  string
	Key   string
	Value string
	Err   error
}

func (err DocumentRetrieveError) Error() string {
	return "Unable to retrieve " + err.Type + " with " + err.Key + " is '" + err.Value + "'"
}

func (err DocumentRetrieveError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m["type"] = err.Type
	m["value"] = err.Value
	m["key"] = err.Key
	m[err.Key] = err.Value
	return m
}

func (err DocumentRetrieveError) JsonResponse() JsonErrorResponse {
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
