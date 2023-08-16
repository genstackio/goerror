package errors

type JsonResponseCompatibleError interface {
	JsonResponse() JsonErrorResponse
}
