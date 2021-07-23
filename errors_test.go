package modzy

import (
	"testing"

	"github.com/pkg/errors"
)

func TestModzyHttpError(t *testing.T) {
	err := &ModzyHTTPError{
		Message: "msg",
	}

	got := err.Error()
	want := "msg"

	if err.Error() != "msg" {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestModzyHttpErrorCause(t *testing.T) {
	causes := map[int]error{
		400: ErrBadRequest,
		401: ErrUnauthorized,
		403: ErrForbidden,
		404: ErrNotFound,
		500: ErrInternalServer,
		418: ErrUnknown,
	}
	for code, expectedErr := range causes {
		modzyErr := &ModzyHTTPError{
			StatusCode: code,
		}
		if errors.Cause(modzyErr) != expectedErr {
			t.Errorf("%d modzy error was not expected cause %v", code, expectedErr)
		}
	}
}
