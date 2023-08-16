package errors

type OwnerMismatchError struct {
	ExpectedOwner string
	ActualOwner   string
}

func (err OwnerMismatchError) Error() string {
	return "Owner mismatch error (expected: " + err.ExpectedOwner + ", actual: " + err.ActualOwner + ")"
}

func (err OwnerMismatchError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m["expectedOwner"] = err.ExpectedOwner
	m["actualOwner"] = err.ActualOwner
	return m
}

func (err OwnerMismatchError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "owner_mismatch",
		Message:    err.Error(),
		Code:       12500,
		StatusCode: 400,
		Data:       err.GetData(),
	}
}
