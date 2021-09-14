package modzy

const (
	DefaultPageSize = 10
)

type SortDirection string

const (
	SortDirectionAscending  SortDirection = "ASC"
	SortDirectionDescending SortDirection = "DESC"
)

// PagingInput -
type PagingInput struct {
	PerPage       int
	Page          int
	SortDirection SortDirection
	SortBy        []string
	Filters       []Filter
}

// NewPaging creates a PagingInput which allows you to set the paging, sorting and filtering information for a List request.
func NewPaging(perPage int, page int) PagingInput {
	pi := PagingInput{
		PerPage: perPage,
		Page:    page,
	}
	return pi.withDefaults()
}

func (p PagingInput) Next() PagingInput {
	page := p.Page
	if page <= 0 {
		page = 1
	}
	p.Page = page + 1
	return p
}

// WithFilter allows adding a filter to the paging information.  Consider using the `And` and `Or` function helpers for situations where you need more complex filtering.
func (p PagingInput) WithFilter(field string, value string) PagingInput {
	return p.WithFilterAnd(field, value)
}

// WithFilterAnd will create a paging filter that will filter for each value provided
func (p PagingInput) WithFilterAnd(field string, values ...string) PagingInput {
	p.Filters = append(p.Filters, Filter{
		Type:   FilterTypeAnd,
		Field:  field,
		Values: values,
	})
	return p
}

// WithFilterOr will create a paging filter that will filter for any value provided
func (p PagingInput) WithFilterOr(field string, values ...string) PagingInput {
	p.Filters = append(p.Filters, Filter{
		Type:   FilterTypeOr,
		Field:  field,
		Values: values,
	})
	return p
}

func (p PagingInput) WithSort(direction SortDirection, by ...string) PagingInput {
	p.SortDirection = direction
	p.SortBy = by
	return p
}

// withDefaults is helpful to ensure that you are dealing with the explicit paging values instead of those defaulted by the backend.
func (p PagingInput) withDefaults() PagingInput {
	if p.PerPage <= 0 {
		p.PerPage = DefaultPageSize
	}
	if p.Page <= 0 {
		p.Page = 1
	}
	return p
}

type FilterType string

const (
	FilterTypeAnd FilterType = "AND"
	FilterTypeOr  FilterType = "OR"
)

type Filter struct {
	Type   FilterType
	Field  string
	Values []string
}
