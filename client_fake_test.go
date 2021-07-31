package modzy

import (
	"testing"
)

func TestClientFake(t *testing.T) {
	calls := 0
	fake := &ClientFake{
		WithAPIKeyFunc: func(apiKey string) Client {
			if apiKey != "apiKey" {
				t.Error("apiKey not passed through")
			}
			calls++
			return nil
		},
		WithTeamKeyFunc: func(teamID string, token string) Client {
			if teamID != "teamID" {
				t.Error("teamID not passed through")
			}

			if token != "token" {
				t.Error("token not passed through")
			}
			calls++
			return nil
		},
		WithOptionsFunc: func(opts ...ClientOption) Client {
			if len(opts) != 2 {
				t.Errorf("did not pass through options")
			}
			calls++
			return nil
		},
		AccountingFunc: func() AccountingClient {
			calls++
			return nil
		},
		JobsFunc: func() JobsClient {
			calls++
			return nil
		},
		ModelsFunc: func() ModelsClient {
			calls++
			return nil
		},
	}
	fake.WithAPIKey("apiKey")
	fake.WithTeamKey("teamID", "token")
	fake.WithOptions(nil, nil)
	fake.Accounting()
	fake.Jobs()
	fake.Models()

	if calls != 6 {
		t.Errorf("Did not call all of the funcs: %d", calls)
	}
}
