package errors

type ValidationError struct {
	FieldNames  []string
	FieldErrors map[string]interface{}
	Err         error
}

func (err ValidationError) Error() string {
	return err.Err.Error()
}

func (err ValidationError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m["fieldNames"] = err.FieldNames
	m["fieldErrors"] = err.FieldErrors

	return m
}

func (err ValidationError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "validation",
		Message:    err.Error(),
		Detail:     err.Err.Error(),
		Code:       12990,
		StatusCode: 412,
		Data:       err.GetData(),
	}
}
