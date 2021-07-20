package modzy

import (
	"net"
	"net/http"
	"time"
)

var defaultHTTPClient = &http.Client{
	Timeout: time.Second * 30,
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	},
}

const (
	JobStatusSubmitted  = "SUBMITTED"
	JobStatusInProgress = "IN_PROGRESS"
	JobStatusCompleted  = "COMPLETED"
	JobStatusCanceled   = "CANCELED"
	JobStatusTimedOut   = "TIMEDOUT"
)

const (
	MinimumWaitForJobInterval = time.Second * 5
)
