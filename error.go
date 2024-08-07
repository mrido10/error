package errs

import "errors"

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
