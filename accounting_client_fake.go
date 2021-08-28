package modzy

import (
	"context"
)

// AccountingClientFake is meant to help in mocking the AccountingClient interface easily for unit testing.
type AccountingClientFake struct {
	GetEntitlementsFunc     func(ctx context.Context) (*GetEntitlementsOutput, error)
	HasEntitlementFunc      func(ctx context.Context, entitlement string) (bool, error)
	GetLicenseFunc          func(ctx context.Context) (*GetLicenseOutput, error)
	ListAccountingUsersFunc func(ctx context.Context, input *ListAccountingUsersInput) (*ListAccountingUsersOutput, error)
	ListProjectsFunc        func(ctx context.Context, input *ListProjectsInput) (*ListProjectsOutput, error)
	GetProjectDetailsFunc   func(ctx context.Context, input *GetProjectDetailsInput) (*GetProjectDetailsOutput, error)
}

var _ AccountingClient = &AccountingClientFake{}

func (c *AccountingClientFake) GetEntitlements(ctx context.Context) (*GetEntitlementsOutput, error) {
	return c.GetEntitlementsFunc(ctx)
}

func (c *AccountingClientFake) HasEntitlement(ctx context.Context, entitlement string) (bool, error) {
	return c.HasEntitlementFunc(ctx, entitlement)
}

func (c *AccountingClientFake) GetLicense(ctx context.Context) (*GetLicenseOutput, error) {
	return c.GetLicenseFunc(ctx)
}

func (c *AccountingClientFake) ListAccountingUsers(ctx context.Context, input *ListAccountingUsersInput) (*ListAccountingUsersOutput, error) {
	return c.ListAccountingUsersFunc(ctx, input)
}

func (c *AccountingClientFake) ListProjects(ctx context.Context, input *ListProjectsInput) (*ListProjectsOutput, error) {
	return c.ListProjectsFunc(ctx, input)
}

func (c *AccountingClientFake) GetProjectDetails(ctx context.Context, input *GetProjectDetailsInput) (*GetProjectDetailsOutput, error) {
	return c.GetProjectDetailsFunc(ctx, input)
}
