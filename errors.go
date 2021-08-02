package modzy

import (
	"fmt"
)

// Known errors
var (
	ErrNotImplemented = fmt.Errorf("method not implemented")
	ErrBadRequest     = fmt.Errorf("the API doesn’t understand the request. Something is missing")
	ErrUnauthorized   = fmt.Errorf("the API key is missing or misspelled")
	ErrForbidden      = fmt.Errorf("the API key doesn’t have the roles required to perform the request")
	ErrNotFound       = fmt.Errorf("the API understands the request but a parameter is missing or misspelled")
	ErrInternalServer = fmt.Errorf("something went wrong on the server’s side")
	ErrUnknown        = fmt.Errorf("an unknown error was returned")
)

// ModzyHTTPError contains additional error information as returned by the http API
type ModzyHTTPError struct {
	StatusCode     int    `json:"statusCode"`
	Status         string `json:"status"`
	Message        string `json:"message"`
	ReportErrorURL string `json:"reportErrorUrl"`
}

func (m *ModzyHTTPError) Error() string {
	return m.Message
}

func (m *ModzyHTTPError) Cause() error {
	switch m.StatusCode {
	case 400:
		return ErrBadRequest
	case 401:
		return ErrUnauthorized
	case 403:
		return ErrForbidden
	case 404:
		return ErrNotFound
	case 500:
		return ErrInternalServer
	}
	return ErrUnknown
}
