package modzy

// ClientFake is meant to help in mocking the Client interface easily for unit testing.
type ClientFake struct {
	WithAPIKeyFunc  func(apiKey string) Client
	WithTeamKeyFunc func(teamID string, token string) Client
	WithOptionsFunc func(opts ...ClientOption) Client
	AccountingFunc  func() AccountingClient
	JobsFunc        func() JobsClient
	ModelsFunc      func() ModelsClient
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

func (c *ClientFake) Accounting() AccountingClient {
	return c.AccountingFunc()
}

func (c *ClientFake) Jobs() JobsClient {
	return c.JobsFunc()
}

func (c *ClientFake) Models() ModelsClient {
	return c.ModelsFunc()
}
