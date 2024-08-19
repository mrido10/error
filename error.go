package errz

import (
	"errors"
	"net/http"
	"strings"
)

type Error struct {
	err        error
	statusCode int
	message    string
}

func New(msg string, statusCode int, err ...error) *Error {
	er := errors.New(msg)
	if len(err) > 1 {
		er = err[0]
	}
	return &Error{
		statusCode: statusCode,
		message:    msg,
		err:        er,
	}
}

func BadRequest(message ...string) *Error {
	code := http.StatusBadRequest
	return New(setErrorMessage(code, message...), code)
}

func Unauthorized(message ...string) *Error {
	code := http.StatusUnauthorized
	return New(setErrorMessage(code, message...), code)
}

func Forbiden(message ...string) *Error {
	code := http.StatusForbidden
	return New(setErrorMessage(code, message...), code)
}

func NotFound(message ...string) *Error {
	code := http.StatusNotFound
	return New(setErrorMessage(code, message...), code)
}

func MethodNotAllowed(message ...string) *Error {
	code := http.StatusMethodNotAllowed
	return New(setErrorMessage(code, message...), code)
}

func RequestTimeout(message ...string) *Error {
	code := http.StatusRequestTimeout
	return New(setErrorMessage(code, message...), code)
}

func InternalServerError(message ...string) *Error {
	code := http.StatusInternalServerError
	return New(setErrorMessage(code, message...), code)
}

func GatewayTimeout(message ...string) *Error {
	code := http.StatusGatewayTimeout
	return New(setErrorMessage(code, message...), code)
}

func setErrorMessage(code int, message ...string) string {
	msg := strings.Join(message, " ")
	if strings.TrimSpace(msg) == "" {
		msg = http.StatusText(code)
	}
	return msg
}

func (e *Error) Error() string {
	if e.err != nil {
		return e.err.Error()
	}
	return e.message
}

func (e *Error) GetStatusCode() int {
	return e.statusCode
}

func (e *Error) GetMessage() string {
	return e.message
}

func (e *Error) GetError() error {
	return e.err
}
