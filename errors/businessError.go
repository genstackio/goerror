package errors

import "strconv"

type BusinessError struct {
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Value   string `json:"value,omitempty"`
}

func (err BusinessError) Error() string {
	return "Business error #" + strconv.Itoa(err.Code) + ": " + err.Message
}

func (err BusinessError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m["code"] = err.Code
	m["message"] = err.Message
	m["value"] = err.Value
	return m
}

func (err BusinessError) JsonResponse() JsonErrorResponse {
	code := err.Code
	if 0 == code {
		code = 12997
	}
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "business",
		Message:    err.Error(),
		Code:       code,
		StatusCode: 403,
		Data:       err.GetData(),
	}
}
