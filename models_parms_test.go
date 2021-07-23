package modzy

import (
	"strings"
	"testing"
)

func TestListModelsInputWithPaging(t *testing.T) {
	i := &ListModelsInput{}
	i.WithPaging(4, 5)

	if i.Paging.PerPage != 4 {
		t.Errorf("expected perPage to be 4, got %d", i.Paging.PerPage)
	}
	if i.Paging.Page != 5 {
		t.Errorf("expected page to be 5, got %d", i.Paging.Page)
	}
}

func TestListModelsInputWithFilterAnd(t *testing.T) {
	i := &ListModelsInput{}
	i.WithFilterAnd(ListModelsFilterFieldAuthor, "a", "b")

	if i.Paging.Filters[0].Field != "author" {
		t.Errorf("expected filter field to be author, got %s", i.Paging.Filters[0].Field)
	}
	if i.Paging.Filters[0].Type != "AND" {
		t.Errorf("expected filter type to be AND, got %s", i.Paging.Filters[0].Type)
	}
	if strings.Join(i.Paging.Filters[0].Values, ",") != "a,b" {
		t.Errorf("expected filter values to be [a, b], got %+v", i.Paging.Filters[0].Values)
	}
}

func TestListModelsInputWithFilterOr(t *testing.T) {
	i := &ListModelsInput{}
	i.WithFilterOr(ListModelsFilterFieldAuthor, "c", "d")

	if i.Paging.Filters[0].Field != "author" {
		t.Errorf("expected filter field to be author, got %s", i.Paging.Filters[0].Field)
	}
	if i.Paging.Filters[0].Type != "OR" {
		t.Errorf("expected filter type to be OR, got %s", i.Paging.Filters[0].Type)
	}
	if strings.Join(i.Paging.Filters[0].Values, ",") != "c,d" {
		t.Errorf("expected filter values to be [c, d], got %+v", i.Paging.Filters[0].Values)
	}
}

func TestListModelVersionsInputWithPaging(t *testing.T) {
	i := &ListModelVersionsInput{}
	i.WithPaging(4, 5)

	if i.Paging.PerPage != 4 {
		t.Errorf("expected perPage to be 4, got %d", i.Paging.PerPage)
	}
	if i.Paging.Page != 5 {
		t.Errorf("expected page to be 5, got %d", i.Paging.Page)
	}
}

func TestListModelVersionsInputWithFilterAnd(t *testing.T) {
	i := &ListModelVersionsInput{}
	i.WithFilterAnd(ListModelVersionsFilterFieldCreatedAt, "a", "b")

	if i.Paging.Filters[0].Field != "createdAt" {
		t.Errorf("expected filter field to be author, got %s", i.Paging.Filters[0].Field)
	}
	if i.Paging.Filters[0].Type != "AND" {
		t.Errorf("expected filter type to be AND, got %s", i.Paging.Filters[0].Type)
	}
	if strings.Join(i.Paging.Filters[0].Values, ",") != "a,b" {
		t.Errorf("expected filter values to be [a, b], got %+v", i.Paging.Filters[0].Values)
	}
}

func TestListModelVersionsInputWithFilterOr(t *testing.T) {
	i := &ListModelVersionsInput{}
	i.WithFilterOr(ListModelVersionsFilterFieldCreatedAt, "c", "d")

	if i.Paging.Filters[0].Field != "createdAt" {
		t.Errorf("expected filter field to be author, got %s", i.Paging.Filters[0].Field)
	}
	if i.Paging.Filters[0].Type != "OR" {
		t.Errorf("expected filter type to be OR, got %s", i.Paging.Filters[0].Type)
	}
	if strings.Join(i.Paging.Filters[0].Values, ",") != "c,d" {
		t.Errorf("expected filter values to be [c, d], got %+v", i.Paging.Filters[0].Values)
	}
}
