package Modzy

import (
	"github.com/dghubble/sling"
	"net/http"
	"time"
)

type Client struct {
	sling  *sling.Sling
	Models *ModelService
	User   *UserService
	Jobs   *JobService
}

func NewClient(base string, apiKey string) *Client {

	httpClient := http.Client{
		Timeout: 120 * time.Second,
	}

	modzyBaseSling := sling.New().Client(&httpClient).Base(base)
	modzyBaseSling.Add("Authorization", "ApiKey " + apiKey)

	return &Client{
		sling:  modzyBaseSling,
		Models: newModelService(modzyBaseSling),
		User:   newUserService(modzyBaseSling, apiKey),
		Jobs:   newJobService(modzyBaseSling),
	}
}
