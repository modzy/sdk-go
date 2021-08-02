package modzy

import (
	"strings"
	"testing"
)

func TestListJobsHistoryInputWithPaging(t *testing.T) {
	i := &ListJobsHistoryInput{}
	i.WithPaging(4, 5)

	if i.Paging.PerPage != 4 {
		t.Errorf("expected perPage to be 4, got %d", i.Paging.PerPage)
	}
	if i.Paging.Page != 5 {
		t.Errorf("expected page to be 5, got %d", i.Paging.Page)
	}
}

func TestListJobsHistoryInputWithFilter(t *testing.T) {
	i := &ListJobsHistoryInput{}
	i.WithFilter(ListJobsHistoryFilterFieldAccessKey, "a")

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

func TestListJobsHistoryInputWithFilterAnd(t *testing.T) {
	i := &ListJobsHistoryInput{}
	i.WithFilterAnd(ListJobsHistoryFilterFieldAccessKey, "a", "b")

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

func TestListJobsHistoryInputWithFilterOr(t *testing.T) {
	i := &ListJobsHistoryInput{}
	i.WithFilterOr(ListJobsHistoryFilterFieldAccessKey, "c", "d")

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

func TestListJobsHistoryInputWitSort(t *testing.T) {
	i := &ListJobsHistoryInput{}
	i.WithSort(SortDirectionDescending, ListJobsHistorySortFieldIdentifier, "c")

	if i.Paging.SortDirection != "DESC" {
		t.Errorf("expected sort-order to be DESC, got %s", i.Paging.SortDirection)
	}
	if strings.Join(i.Paging.SortBy, ",") != "identifier,c" {
		t.Errorf("expected filter values to be [identifier,c], got %+v", i.Paging.SortBy)
	}
}
