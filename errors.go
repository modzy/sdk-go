package modzy

import (
	"fmt"
)

// Known errors
var (
	ErrNotImplemented = fmt.Errorf("Method not implemented")
	ErrBadRequest     = fmt.Errorf("The API doesn’t understand the request. Something is missing")
	ErrUnauthorized   = fmt.Errorf("The API key is missing or misspelled")
	ErrForbidden      = fmt.Errorf("The API key doesn’t have the roles required to perform the request")
	ErrNotFound       = fmt.Errorf("The API understands the request but a parameter is missing or misspelled")
	ErrInternalServer = fmt.Errorf("Something went wrong on the server’s side")
	ErrUnknown        = fmt.Errorf("An unknown error was returned")
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
