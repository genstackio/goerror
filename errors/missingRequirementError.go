package errors

type MissingRequirementError struct {
	Name string
}

func (err MissingRequirementError) Error() string {
	return "No " + err.Name + " specified"
}

func (err MissingRequirementError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m["name"] = err.Name
	return m
}

func (err MissingRequirementError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "missing_requirement",
		Message:    err.Error(),
		Code:       12910,
		StatusCode: 400,
		Data:       err.GetData(),
	}
}
