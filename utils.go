package goerror

import (
	"encoding/json"
	"github.com/genstackio/goerror/errors"
	"net/http"
)

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

//goland:noinspection GoUnusedExportedFunction
func WriteError(w http.ResponseWriter, err interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("X-Content-Type-Options", "nosniff")

	response := FormatJsonErrorResponse(err)
	w.WriteHeader(response.StatusCode)

	erro, _ := json.Marshal(response)
	_, _ = w.Write(erro)
}
