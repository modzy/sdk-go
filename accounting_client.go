package modzy

import (
	"context"
	"sync"

	"github.com/modzy/go-sdk/model"
)

type AccountingClient interface {
	GetEntitlements(ctx context.Context) (*GetEntitlementsOutput, error)
	HasEntitlement(ctx context.Context, entitlement string) (bool, error)
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
