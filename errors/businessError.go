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
