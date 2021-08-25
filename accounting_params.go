package modzy

import "github.com/modzy/sdk-go/model"

// GetEntitlementsOutput -
type GetEntitlementsOutput struct {
	Entitlements []model.Entitlement `json:"entitlements"`
}

type ListAccountingUsersInput struct {
	Paging PagingInput
}

// ListAccountingUsersFilterField are known field names that can be used when filtering the jobs history
type ListAccountingUsersFilterField string

const (
	ListAccountingUsersFilterFieldFirstName  ListAccountingUsersFilterField = "firstName"
	ListAccountingUsersFilterFieldLastName   ListAccountingUsersFilterField = "lastName"
	ListAccountingUsersFilterFieldEmail      ListAccountingUsersFilterField = "email"
	ListAccountingUsersFilterFieldSearch     ListAccountingUsersFilterField = "search"
	ListAccountingUsersFilterFieldStatus     ListAccountingUsersFilterField = "status"
	ListAccountingUsersFilterFieldAccessKey  ListAccountingUsersFilterField = "accessKey"
	ListAccountingUsersFilterFieldStartDate  ListAccountingUsersFilterField = "startDate"
	ListAccountingUsersFilterFieldEndDate    ListAccountingUsersFilterField = "endDate"
	ListAccountingUsersFilterFieldSearchDate ListAccountingUsersFilterField = "searchDate"
)

func (i *ListAccountingUsersInput) WithPaging(perPage int, page int) *ListAccountingUsersInput {
	i.Paging = NewPaging(perPage, page)
	return i
}

func (i *ListAccountingUsersInput) WithFilter(field ListAccountingUsersFilterField, value string) *ListAccountingUsersInput {
	i.Paging = i.Paging.WithFilterAnd(string(field), value)
	return i
}

func (i *ListAccountingUsersInput) WithFilterAnd(field ListAccountingUsersFilterField, values ...string) *ListAccountingUsersInput {
	i.Paging = i.Paging.WithFilterAnd(string(field), values...)
	return i
}

func (i *ListAccountingUsersInput) WithFilterOr(field ListAccountingUsersFilterField, values ...string) *ListAccountingUsersInput {
	i.Paging = i.Paging.WithFilterOr(string(field), values...)
	return i
}

type ListAccountingUsersOutput struct {
	Users    []model.AccountingUser    `json:"users"`
	NextPage *ListAccountingUsersInput `json:"nextPage"`
}

type GetLicenseOutput struct {
	License model.License `json:"license"`
}
