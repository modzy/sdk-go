package modzy

import (
	"context"
	"testing"
)

func TestAccountingClientFake(t *testing.T) {
	expectedCtx := context.WithValue(context.TODO(), testContextKey("a"), "b")

	calls := 0
	fake := &AccountingClientFake{
		GetEntitlementsFunc: func(ctx context.Context) (*GetEntitlementsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			return nil, nil
		},
		HasEntitlementFunc: func(ctx context.Context, entitlement string) (bool, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if entitlement != "entitlement" {
				t.Errorf("input was not passed through")
			}
			return false, nil
		},
		GetLicenseFunc: func(ctx context.Context) (*GetLicenseOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			return nil, nil
		},
		ListAccountingUsersFunc: func(ctx context.Context, input *ListAccountingUsersInput) (*ListAccountingUsersOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		ListProjectsFunc: func(ctx context.Context, input *ListProjectsInput) (*ListProjectsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetProjectDetailsFunc: func(ctx context.Context, input *GetProjectDetailsInput) (*GetProjectDetailsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
	}

	fake.GetEntitlements(expectedCtx)
	fake.HasEntitlement(expectedCtx, "entitlement")
	fake.GetLicense(expectedCtx)
	fake.ListAccountingUsers(expectedCtx, &ListAccountingUsersInput{})
	fake.ListProjects(expectedCtx, &ListProjectsInput{})
	fake.GetProjectDetails(expectedCtx, &GetProjectDetailsInput{})

	if calls != 6 {
		t.Errorf("Did not call all of the funcs: %d", calls)
	}
}
