package modzy

import (
	"fmt"
	"net/http"
)

type Client interface {
	WithAPIKey(apiKey string) Client
	WithTeamKey(teamID string, token string) Client
	WithOptions(opts ...ClientOption) Client
	Jobs() JobsClient
	Models() ModelsClient
	// Tags() TagsClient
}

type standardClient struct {
	jobsClient   *standardJobsClient
	modelsClient *standardModelsClient
	requestor    *requestor
}

func NewClient(baseURL string, opts ...ClientOption) Client {
	var client = &standardClient{
		requestor: &requestor{
			baseURL:    baseURL,
			httpClient: defaultHTTPClient,
		},
	}
	client.WithOptions(opts...)

	// setup our namespaced groupings that all share this as their base client
	client.jobsClient = &standardJobsClient{
		baseClient: client,
	}
	client.modelsClient = &standardModelsClient{
		baseClient: client,
	}

	return client
}

// WithAPIKey -
func (c *standardClient) WithAPIKey(apiKey string) Client {
	c.requestor.authorizationDecorator = func(req *http.Request) *http.Request {
		req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", apiKey))
		return req
	}
	return c
}

// WithTeamKey -
func (c *standardClient) WithTeamKey(teamID string, token string) Client {
	c.requestor.authorizationDecorator = func(req *http.Request) *http.Request {
		req.Header.Add("Modzy-Team-Id", teamID)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		return req
	}
	return c
}

// WithOptions -
func (c *standardClient) WithOptions(opts ...ClientOption) Client {
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Jobs returns a client for access to all job related API functions
func (c *standardClient) Jobs() JobsClient {
	return c.jobsClient
}

// Models returns a client for access to all model related API functions
func (c *standardClient) Models() ModelsClient {
	return c.modelsClient
}
