package errors

import "strconv"

type BusinessError struct {
	Code    int
	Message string
	Value   string
}

func (err BusinessError) Error() string {
	return "Business error #" + strconv.Itoa(err.Code) + ": " + err.Message
}
