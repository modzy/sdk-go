package modzy

import (
	"net/http"
)

type ClientOption func(*standardClient)

// WithHTTPClient allows providing a custom underlying http client.  It is good practice to _not_ use the default http client
// that Go provides as it has no timeouts.  If you do not provide your own default client, a reasonable one will be created for you.
func WithHTTPClient(httpClient *http.Client) ClientOption {
	return func(c *standardClient) {
		c.requestor.(*stdRequestor).httpClient = httpClient
	}
}

// WithHTTPDebugging will trigger logrus debug messages to be emitted with the raw request and response information.
// This should only be used for debugging purposes as it can decode entire messages into memory.
func WithHTTPDebugging(request bool, response bool) ClientOption {
	return func(c *standardClient) {
		c.requestor.(*stdRequestor).requestDebugging = request
		c.requestor.(*stdRequestor).responseDebugging = response
	}
}

// withRequestor is helpful for mock testing; do not use this outside of unit testing, and do not combine this with other Options
func withRequestor(requestor requestor) ClientOption {
	return func(c *standardClient) {
		c.requestor = requestor
	}
}
