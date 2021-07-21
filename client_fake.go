package modzy

// ClientFake is meant to help in mocking the Client interface easily for unit testing.
type ClientFake struct {
	WithAPIKeyFunc  func(apiKey string) Client
	WithTeamKeyFunc func(teamID string, token string) Client
	WithOptionsFunc func(opts ...ClientOption) Client
	JobsFunc        func() JobsClient
}

var _ Client = &ClientFake{}

func (c *ClientFake) WithAPIKey(apiKey string) Client {
	return c.WithAPIKeyFunc(apiKey)
}

func (c *ClientFake) WithTeamKey(teamID string, token string) Client {
	return c.WithTeamKeyFunc(teamID, token)
}

func (c *ClientFake) WithOptions(opts ...ClientOption) Client {
	return c.WithOptionsFunc(opts...)
}

func (c *ClientFake) Jobs() JobsClient {
	return c.JobsFunc()
}
