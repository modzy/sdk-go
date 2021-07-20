package modzy

import (
	"testing"
)

func TestNewPaging(t *testing.T) {
	cases := []struct {
		PerPage         int
		Page            int
		ExpectedPerPage int
		ExpectedPage    int
	}{
		{0, 0, 10, 1},
		{0, 2, 10, 2},
		{3, 0, 3, 1},
		{-1, -1, 10, 1},
	}
	for _, c := range cases {
		paging := NewPaging(c.PerPage, c.Page)
		if paging.PerPage != c.ExpectedPerPage {
			t.Errorf("got %d want %d", c.PerPage, c.ExpectedPerPage)
		}
		if paging.Page != c.ExpectedPage {
			t.Errorf("got %d want %d", c.Page, c.Page)
		}
	}
}

func TestPagingChain(t *testing.T) {
	paging := NewPaging(20, 30).
		WithSort(SortDirectionDescending, "a", "b").
		WithFilter(Or("for", "o1", "o2")).
		WithFilter(And("fand", "a1", "a2"))

	if paging.Page != 30 {
		t.Errorf("Page not set")
	}

	// next should propagate
	next := paging.Next()

	if next.Page != 31 {
		t.Errorf("Page not nexted properly, got %d:", next.Page)
	}

	// bot this and the next should be the same in fitlers and sorting
	for _, gotPaging := range []PagingInput{paging, paging.Next()} {
		if gotPaging.PerPage != 20 {
			t.Errorf("PerPage not set")
		}

		if gotPaging.SortDirection != SortDirectionDescending {
			t.Errorf("SortDirection not set")
		}
		if len(gotPaging.SortBy) != 2 || gotPaging.SortBy[0] != "a" || gotPaging.SortBy[1] != "b" {
			t.Errorf("SortBy fields not set, got %v", gotPaging.SortBy)
		}

		if len(gotPaging.Filters) != 2 {
			t.Errorf("Filters total not correct, got %+v", gotPaging.Filters)
		}
		if gotPaging.Filters[0].Field != "for" ||
			(gotPaging.Filters[0].Type != FilterTypeOr || len(gotPaging.Filters[0].Values) != 2) ||
			(gotPaging.Filters[0].Values[0] != "o1" || gotPaging.Filters[0].Values[1] != "o2") {
			t.Errorf("Filters 0 OR not correct, got %+v", gotPaging.Filters[0])
		}

		if gotPaging.Filters[1].Field != "fand" ||
			(gotPaging.Filters[1].Type != FilterTypeAnd || len(gotPaging.Filters[1].Values) != 2) ||
			(gotPaging.Filters[1].Values[0] != "a1" || gotPaging.Filters[1].Values[1] != "a2") {
			t.Errorf("Filters 1 AND not correct, got %+v", gotPaging.Filters[1])
		}
	}
}

func TestPagingNext(t *testing.T) {
	cases := []struct {
		Page             int
		ExpectedNextPage int
	}{
		{-1, 2},
		{0, 2},
		{1, 2},
		{2, 3},
	}

	for _, c := range cases {

		if (PagingInput{Page: c.Page}).Next().Page != c.ExpectedNextPage {
			t.Errorf("%d did not Next to %d", c.Page, c.ExpectedNextPage)
		}
	}
}
