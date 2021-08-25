package modzy

import (
	"fmt"
	"net"
	"net/http"
	"time"
)

type Client interface {
	// WithAPIKey allows you to provided your ApiKey.  Use this or WithTeamKey.
	WithAPIKey(apiKey string) Client
	// WithTeamKey allows you to provided your team credentials.  Use this or WithAPIKey.
	WithTeamKey(teamID string, token string) Client
	// WithOptions allows providing additional client options such as WithHTTPDebugging. These are not commonly needed.
	WithOptions(opts ...ClientOption) Client
	// Accounting returns a client for access to all accounting related API functions
	Accounting() AccountingClient
	// Jobs returns a client for access to all job related API functions
	Jobs() JobsClient
	// Models returns a client for access to all model related API functions
	Models() ModelsClient
	// Dashboard returns a client for access to dashboard API functions
	Dashboard() DashboardClient
	// Resources returns a client for access to resource information
	Resources() ResourcesClient
}

type standardClient struct {
	accountingClient *standardAccountingClient
	jobsClient       *standardJobsClient
	modelsClient     *standardModelsClient
	dashboardClient  *standardDashboardClient
	resourcesClient  *standardResourcesClient
	requestor        *requestor
}

var defaultHTTPClient = &http.Client{
	Timeout: time.Second * 30,
	Transport: &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 10 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 10 * time.Second,
	},
}

// NewClient will create a standard client for the given baseURL.
// You need to provide your authentication key to the client through one of two methods:
// 	client.WithAPIKey(apiKey) or client.WithTeamKey(teamID, token)
func NewClient(baseURL string, opts ...ClientOption) Client {
	var client = &standardClient{
		requestor: &requestor{
			baseURL:    baseURL,
			httpClient: defaultHTTPClient,
		},
	}
	client.WithOptions(opts...)

	// setup our namespaced groupings that all share this as their base client
	client.accountingClient = &standardAccountingClient{
		baseClient: client,
	}
	client.jobsClient = &standardJobsClient{
		baseClient: client,
	}
	client.modelsClient = &standardModelsClient{
		baseClient: client,
	}
	client.dashboardClient = &standardDashboardClient{
		baseClient: client,
	}
	client.resourcesClient = &standardResourcesClient{
		baseClient: client,
	}

	return client
}

func (c *standardClient) WithAPIKey(apiKey string) Client {
	c.requestor.authorizationDecorator = func(req *http.Request) *http.Request {
		req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", apiKey))
		return req
	}
	return c
}

func (c *standardClient) WithTeamKey(teamID string, token string) Client {
	c.requestor.authorizationDecorator = func(req *http.Request) *http.Request {
		req.Header.Add("Modzy-Team-Id", teamID)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		return req
	}
	return c
}

func (c *standardClient) WithOptions(opts ...ClientOption) Client {
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func (c *standardClient) Accounting() AccountingClient {
	return c.accountingClient
}

func (c *standardClient) Jobs() JobsClient {
	return c.jobsClient
}

func (c *standardClient) Models() ModelsClient {
	return c.modelsClient
}

func (c *standardClient) Dashboard() DashboardClient {
	return c.dashboardClient
}

func (c *standardClient) Resources() ResourcesClient {
	return c.resourcesClient
}
