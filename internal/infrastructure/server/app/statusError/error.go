package statusError

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
