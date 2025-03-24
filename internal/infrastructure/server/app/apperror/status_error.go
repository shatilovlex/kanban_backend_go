package apperror

import (
	"errors"
	"net/http"
)

type StatusError struct {
	error
	status int
}

func (e StatusError) Unwrap() error   { return e.error }
func (e StatusError) HTTPStatus() int { return e.status }

func WithHTTPStatus(err error, status int) error {
	return StatusError{
		error:  err,
		status: status,
	}
}

func HTTPStatus(err error) int {
	if err == nil {
		return 0
	}
	var statusErr interface {
		error
		HTTPStatus() int
	}
	if errors.As(err, &statusErr) {
		return statusErr.HTTPStatus()
	}
	return http.StatusInternalServerError
}
