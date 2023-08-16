package errors

type TenantMismatchError struct {
	ExpectedTenant string
	ActualTenant   string
}

func (err TenantMismatchError) Error() string {
	return "Tenant mismatch error (expected: " + err.ExpectedTenant + ", actual: " + err.ActualTenant + ")"
}

func (err TenantMismatchError) GetData() map[string]interface{} {
	m := map[string]interface{}{}
	m["expectedTenant"] = err.ExpectedTenant
	m["actualTenant"] = err.ActualTenant
	return m
}

func (err TenantMismatchError) JsonResponse() JsonErrorResponse {
	return JsonErrorResponse{
		Status:     "error",
		ErrorType:  "tenant_mismatch",
		Message:    err.Error(),
		Code:       12400,
		StatusCode: 400,
		Data:       err.GetData(),
	}
}
