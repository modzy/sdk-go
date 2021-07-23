package modzy

import (
	"context"
	"net/http"

	"github.com/peterhellberg/link"
)

type fakeRequestor struct {
	GetFunc    func(ctx context.Context, path string, into interface{}) (*http.Response, error)
	ListFunc   func(ctx context.Context, path string, paging PagingInput, into interface{}) (*http.Response, link.Group, error)
	PostFunc   func(ctx context.Context, path string, toPost interface{}, into interface{}) (*http.Response, error)
	DeleteFunc func(ctx context.Context, path string, into interface{}) (*http.Response, error)
}

var _ requestor = &fakeRequestor{}

func (r *fakeRequestor) Get(ctx context.Context, path string, into interface{}) (*http.Response, error) {
	return r.GetFunc(ctx, path, into)
}
func (r *fakeRequestor) List(ctx context.Context, path string, paging PagingInput, into interface{}) (*http.Response, link.Group, error) {
	return r.ListFunc(ctx, path, paging, into)
}

func (r *fakeRequestor) Post(ctx context.Context, path string, toPost interface{}, into interface{}) (*http.Response, error) {
	return r.PostFunc(ctx, path, toPost, into)
}

func (r *fakeRequestor) Delete(ctx context.Context, path string, into interface{}) (*http.Response, error) {
	return r.DeleteFunc(ctx, path, into)
}
