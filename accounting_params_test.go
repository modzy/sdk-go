package modzy

import (
	"strings"
	"testing"
)

func TestListAccountingUsersInputWithPaging(t *testing.T) {
	i := &ListAccountingUsersInput{}
	i.WithPaging(4, 5)

	if i.Paging.PerPage != 4 {
		t.Errorf("expected perPage to be 4, got %d", i.Paging.PerPage)
	}
	if i.Paging.Page != 5 {
		t.Errorf("expected page to be 5, got %d", i.Paging.Page)
	}
}

func TestListAccountingUsersInputWithFilter(t *testing.T) {
	i := &ListAccountingUsersInput{}
	i.WithFilter(ListAccountingUsersFilterFieldAccessKey, "a")

	if i.Paging.Filters[0].Field != "accessKey" {
		t.Errorf("expected filter field to be accessKey, got %s", i.Paging.Filters[0].Field)
	}
	if i.Paging.Filters[0].Type != "AND" {
		t.Errorf("expected filter type to be AND, got %s", i.Paging.Filters[0].Type)
	}
	if strings.Join(i.Paging.Filters[0].Values, ",") != "a" {
		t.Errorf("expected filter values to be [a], got %+v", i.Paging.Filters[0].Values)
	}
}

func TestListAccountingUsersInputWithFilterAnd(t *testing.T) {
	i := &ListAccountingUsersInput{}
	i.WithFilterAnd(ListAccountingUsersFilterFieldAccessKey, "a", "b")

	if i.Paging.Filters[0].Field != "accessKey" {
		t.Errorf("expected filter field to be accessKey, got %s", i.Paging.Filters[0].Field)
	}
	if i.Paging.Filters[0].Type != "AND" {
		t.Errorf("expected filter type to be AND, got %s", i.Paging.Filters[0].Type)
	}
	if strings.Join(i.Paging.Filters[0].Values, ",") != "a,b" {
		t.Errorf("expected filter values to be [a, b], got %+v", i.Paging.Filters[0].Values)
	}
}

func TestListAccountingUsersInputWithFilterOr(t *testing.T) {
	i := &ListAccountingUsersInput{}
	i.WithFilterOr(ListAccountingUsersFilterFieldAccessKey, "c", "d")

	if i.Paging.Filters[0].Field != "accessKey" {
		t.Errorf("expected filter field to be accessKey, got %s", i.Paging.Filters[0].Field)
	}
	if i.Paging.Filters[0].Type != "OR" {
		t.Errorf("expected filter type to be OR, got %s", i.Paging.Filters[0].Type)
	}
	if strings.Join(i.Paging.Filters[0].Values, ",") != "c,d" {
		t.Errorf("expected filter values to be [c, d], got %+v", i.Paging.Filters[0].Values)
	}
}

func TestListProjectsInputWithPaging(t *testing.T) {
	i := &ListProjectsInput{}
	i.WithPaging(4, 5)

	if i.Paging.PerPage != 4 {
		t.Errorf("expected perPage to be 4, got %d", i.Paging.PerPage)
	}
	if i.Paging.Page != 5 {
		t.Errorf("expected page to be 5, got %d", i.Paging.Page)
	}
}

func TestListProjectsInputWithFilter(t *testing.T) {
	i := &ListProjectsInput{}
	i.WithFilter(ListProjectsFilterFieldSearch, "a")

	if i.Paging.Filters[0].Field != "search" {
		t.Errorf("expected filter field to be search, got %s", i.Paging.Filters[0].Field)
	}
	if i.Paging.Filters[0].Type != "AND" {
		t.Errorf("expected filter type to be AND, got %s", i.Paging.Filters[0].Type)
	}
	if strings.Join(i.Paging.Filters[0].Values, ",") != "a" {
		t.Errorf("expected filter values to be [a], got %+v", i.Paging.Filters[0].Values)
	}
}

func TestListProjectsInputWithFilterAnd(t *testing.T) {
	i := &ListProjectsInput{}
	i.WithFilterAnd(ListProjectsFilterFieldSearch, "a", "b")

	if i.Paging.Filters[0].Field != "search" {
		t.Errorf("expected filter field to be search, got %s", i.Paging.Filters[0].Field)
	}
	if i.Paging.Filters[0].Type != "AND" {
		t.Errorf("expected filter type to be AND, got %s", i.Paging.Filters[0].Type)
	}
	if strings.Join(i.Paging.Filters[0].Values, ",") != "a,b" {
		t.Errorf("expected filter values to be [a, b], got %+v", i.Paging.Filters[0].Values)
	}
}

func TestListProjectsInputWithFilterOr(t *testing.T) {
	i := &ListProjectsInput{}
	i.WithFilterOr(ListProjectsFilterFieldStatus, "c", "d")

	if i.Paging.Filters[0].Field != "status" {
		t.Errorf("expected filter field to be status, got %s", i.Paging.Filters[0].Field)
	}
	if i.Paging.Filters[0].Type != "OR" {
		t.Errorf("expected filter type to be OR, got %s", i.Paging.Filters[0].Type)
	}
	if strings.Join(i.Paging.Filters[0].Values, ",") != "c,d" {
		t.Errorf("expected filter values to be [c, d], got %+v", i.Paging.Filters[0].Values)
	}
}
