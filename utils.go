package goerror

import "github.com/genstackio/goerror/errors"

//goland:noinspection GoUnusedExportedFunction
func FormatJsonErrorResponse(err interface{}) errors.JsonErrorResponse {
	_, ok := err.(errors.JsonResponseCompatibleError)
	if ok {
		return (err.(errors.JsonResponseCompatibleError)).JsonResponse()
	}

	return FormatGenericError(err)
}

func FormatGenericError(err interface{}) errors.JsonErrorResponse {
	return errors.JsonErrorResponse{
		Status:     "error",
		Message:    (err.(error)).Error(),
		Code:       10000,
		StatusCode: 500,
	}
}
