package apperror

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHTTPStatus(t *testing.T) {
	type args struct {
		err error
	}
	tests := []struct {
		args args
		name string
		want int
	}{
		{
			name: "nil error",
			args: args{
				err: nil,
			},
			want: 0,
		},
		{
			name: "not found error",
			args: args{
				err: StatusError{
					error:  fmt.Errorf("not found"),
					status: 404,
				},
			},
			want: http.StatusNotFound,
		},
		{
			name: "internal server error",
			args: args{
				err: StatusError{
					error:  fmt.Errorf("internal server error"),
					status: 500,
				},
			},
			want: http.StatusInternalServerError,
		},
		{
			name: "bad request",
			args: args{
				err: StatusError{
					error:  fmt.Errorf("bad request"),
					status: 400,
				},
			},
			want: http.StatusBadRequest,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, HTTPStatus(tt.args.err), tt.want)
		})
	}
}

func TestStatusError_HTTPStatus(t *testing.T) {
	e := StatusError{
		error:  fmt.Errorf("internal server error"),
		status: 500,
	}

	assert.Equal(t, e.Unwrap(), fmt.Errorf("internal server error"))
	assert.Equal(t, e.HTTPStatus(), 500)
}
