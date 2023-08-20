package errors

import (
	"fmt"
)

type DocumentsFindError struct {
	Type   string
	Err    error
	Params map[string]interface{}
}

func (err DocumentsFindError) Error() string {
	return "Unable to find " + err.Type + " documents: " + err.Err.Error()
}

func (err DocumentsFindError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m["type"] = err.Type
	for k, v := range err.Params {
		m[k] = fmt.Sprintf("%s", v)
	}
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
