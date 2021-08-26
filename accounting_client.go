package modzy

import (
	"context"
	"fmt"
	"sync"

	"github.com/modzy/sdk-go/model"
)

type AccountingClient interface {
	// GetEntitlements will get all of your entitlements
	GetEntitlements(ctx context.Context) (*GetEntitlementsOutput, error)
	// HasEntitlement will return true if you have the provided entitlement is
	HasEntitlement(ctx context.Context, entitlement string) (bool, error)
	// GetLicense returns a truncated view of your license information
	GetLicense(ctx context.Context) (*GetLicenseOutput, error)
	// ListAccountingUsers returns account user information
	ListAccountingUsers(ctx context.Context, input *ListAccountingUsersInput) (*ListAccountingUsersOutput, error)
	// ListProjects will list your projects
	ListProjects(ctx context.Context, input *ListProjectsInput) (*ListProjectsOutput, error)
	// GetProjectDetails will read the deatils about a project
	GetProjectDetails(ctx context.Context, input *GetProjectDetailsInput) (*GetProjectDetailsOutput, error)
}

type standardAccountingClient struct {
	sync.Mutex
	baseClient       *standardClient
	entitlementCache []model.Entitlement
}

var _ AccountingClient = &standardAccountingClient{}

func (c *standardAccountingClient) GetEntitlements(ctx context.Context) (*GetEntitlementsOutput, error) {
	var out []model.Entitlement
	url := "/api/accounting/entitlements"
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetEntitlementsOutput{
		Entitlements: out,
	}, nil
}

func (c *standardAccountingClient) HasEntitlement(ctx context.Context, entitlement string) (bool, error) {
	c.Lock()
	defer c.Unlock()
	if c.entitlementCache == nil {
		out, err := c.GetEntitlements(ctx)
		if err != nil {
			return false, err
		}
		c.entitlementCache = out.Entitlements
	}

	for _, e := range c.entitlementCache {
		if e.Identifier == entitlement {
			return true, nil
		}
	}
	return false, nil
}
func (c *standardAccountingClient) GetLicense(ctx context.Context) (*GetLicenseOutput, error) {
	var out model.License
	url := "/api/license"
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetLicenseOutput{
		License: out,
	}, nil
}

func (c *standardAccountingClient) ListAccountingUsers(ctx context.Context, input *ListAccountingUsersInput) (*ListAccountingUsersOutput, error) {
	input.Paging = input.Paging.withDefaults()

	var items []model.AccountingUser
	url := "/api/accounting/users"
	_, links, err := c.baseClient.requestor.List(ctx, url, input.Paging, &items)
	if err != nil {
		return nil, err
	}

	// decide if we have a next page (the next link is not always accurate?)
	var nextPage *ListAccountingUsersInput
	if _, hasNextLink := links["next"]; len(items) == input.Paging.PerPage && hasNextLink {
		nextPage = &ListAccountingUsersInput{
			Paging: input.Paging.Next(),
		}
	}

	return &ListAccountingUsersOutput{
		Users:    items,
		NextPage: nextPage,
	}, nil
}

func (c *standardAccountingClient) ListProjects(ctx context.Context, input *ListProjectsInput) (*ListProjectsOutput, error) {
	input.Paging = input.Paging.withDefaults()

	var items []model.AccountingProject
	url := "/api/accounting/projects"
	_, links, err := c.baseClient.requestor.List(ctx, url, input.Paging, &items)
	if err != nil {
		return nil, err
	}

	// decide if we have a next page (the next link is not always accurate?)
	var nextPage *ListProjectsInput
	if _, hasNextLink := links["next"]; len(items) == input.Paging.PerPage && hasNextLink {
		nextPage = &ListProjectsInput{
			Paging: input.Paging.Next(),
		}
	}

	return &ListProjectsOutput{
		Projects: items,
		NextPage: nextPage,
	}, nil
}

func (c *standardAccountingClient) GetProjectDetails(ctx context.Context, input *GetProjectDetailsInput) (*GetProjectDetailsOutput, error) {
	var out model.AccountingProject
	url := fmt.Sprintf("/api/accounting/projects/%s", input.ProjectID)
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetProjectDetailsOutput{
		Project: out,
	}, nil
}
