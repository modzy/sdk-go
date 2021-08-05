package modzy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"strings"

	"github.com/modzy/sdk-go/internal/impossible"

	"github.com/peterhellberg/link"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type requestDecorator func(*http.Request) *http.Request

type requestor struct {
	baseURL                string
	authorizationDecorator requestDecorator
	requestDebugging       bool
	responseDebugging      bool
	httpClient             *http.Client
}

func (r *requestor) execute(
	ctx context.Context,
	path string, method string, toPostInput interface{}, into interface{},
	reqDecorator requestDecorator,
) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", r.baseURL, path)

	// if we are handed a reader, then don't treat it as a json input
	var toPost io.Reader
	if toPostInput != nil {
		switch v := toPostInput.(type) {
		case io.Reader:
			toPost = v
		default:
			toPostBytes, err := json.Marshal(toPostInput)
			if err != nil {
				return nil, errors.WithMessagef(err, "failed to marshal provided body to %s:%s", method, path)
			}
			toPost = bytes.NewReader(toPostBytes)
		}
	}

	req, err := http.NewRequest(method, url, toPost)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to create request to %s:%s", method, path)
	}
	req = req.WithContext(ctx)

	if r.authorizationDecorator != nil {
		r.authorizationDecorator(req)
	}

	if reqDecorator != nil {
		reqDecorator(req)
	}

	if r.requestDebugging {
		// jsonize again if debugging
		bodyDebug := ""
		switch toPostInput.(type) {
		case io.Reader:
			bodyDebug = "reader provided, will not read"
		default:
			debugJson, debugErr := json.Marshal(toPostInput)
			bodyDebug = fmt.Sprintf("%v => %s", debugErr, string(debugJson))
		}
		logrus.WithFields(logrus.Fields{
			"url":     req.URL,
			"method":  method,
			"body":    bodyDebug,
			"headers": req.Header,
		}).Debug("API request")
	}

	resp, err := r.httpClient.Do(req)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to executing request to %s:%s", method, path)
	}
	defer resp.Body.Close()

	var toDecode io.Reader = resp.Body

	if r.responseDebugging {
		body, debugErr := ioutil.ReadAll(resp.Body)
		logrus.WithFields(logrus.Fields{
			"method":     req.Method,
			"url":        req.URL,
			"statusCode": resp.StatusCode,
			// "headers":    resp.Header,
			"body": fmt.Sprintf("%v => %s", debugErr, string(body)),
		}).Debug("API response")
		toDecode = bytes.NewReader(body)
	}

	if resp.StatusCode >= 400 {
		// non OK response
		apiError := &ModzyHTTPError{}
		if err := json.NewDecoder(toDecode).Decode(apiError); err != nil {
			return resp, errors.WithMessagef(err, "failed parsing non 200 response of %d from %s:%s", resp.StatusCode, method, path)
		}
		return resp, apiError
	}

	if resp.StatusCode == 204 {
		// No Content
		// Do not bother with response body stuff; none expected
	} else {
		if into != nil {
			if err := json.NewDecoder(toDecode).Decode(into); err != nil {
				return resp, errors.WithMessagef(err, "failed parsing the response from %s:%s", method, path)
			}
		}
	}

	return resp, nil
}

func (r *requestor) Get(ctx context.Context, path string, into interface{}) (*http.Response, error) {
	return r.execute(ctx, path, "GET", nil, into, jsonDecorator)
}

func (r *requestor) List(ctx context.Context, path string, paging PagingInput, into interface{}) (*http.Response, link.Group, error) {
	// append paging information to our url query
	partialUrl, err := url.Parse(path)
	if err != nil {
		return nil, link.Group{}, err
	}
	q := partialUrl.Query()
	if paging.PerPage != 0 {
		q.Add("per-page", fmt.Sprintf("%d", paging.PerPage))
	}
	if paging.Page != 0 {
		q.Add("page", fmt.Sprintf("%d", paging.Page))
	}
	if paging.SortDirection != "" {
		q.Add("direction", string(paging.SortDirection))
	}
	if len(paging.SortBy) != 0 {
		q.Add("sort-by", strings.Join(paging.SortBy, ","))
	}

	for _, filter := range paging.Filters {
		switch filter.Type {
		case FilterTypeAnd:
			q.Add(filter.Field, strings.Join(filter.Values, ";"))
		default:
			q.Add(filter.Field, strings.Join(filter.Values, ","))
		}
	}

	partialUrl.RawQuery = q.Encode()

	resp, err := r.Get(ctx, partialUrl.String(), into)
	if err != nil {
		return resp, link.Group{}, err
	}

	// parse out the link response header
	links := link.ParseResponse(resp)

	return resp, links, err
}

func (r *requestor) Post(ctx context.Context, path string, toPost interface{}, into interface{}) (*http.Response, error) {
	return r.execute(ctx, path, "POST", toPost, into, jsonDecorator)
}

func (r *requestor) Patch(ctx context.Context, path string, toPatch interface{}, into interface{}) (*http.Response, error) {
	return r.execute(ctx, path, "PATCH", toPatch, into, jsonDecorator)
}

func (r *requestor) Delete(ctx context.Context, path string, into interface{}) (*http.Response, error) {
	return r.execute(ctx, path, "DELETE", nil, into, jsonDecorator)
}

func (r *requestor) PostMultipart(ctx context.Context, path string, filesDatas map[string]io.Reader, into interface{}) (*http.Response, error) {
	var b bytes.Buffer
	w := multipart.NewWriter(&b)

	for key, keyReader := range filesDatas {
		// the endpoint expects the filename to the be the key, not just a simple part name
		fw, err := w.CreateFormFile(key, key)
		impossible.HandleError(err)

		if _, err = io.Copy(fw, keyReader); err != nil {
			return nil, err
		}
	}
	w.Close()

	return r.execute(ctx, path, "POST", &b, into, func(req *http.Request) *http.Request {
		req.Header.Set("Content-Type", w.FormDataContentType())
		return req
	})
}

func jsonDecorator(req *http.Request) *http.Request {
	req.Header.Set("Content-Type", "application/json")
	return req
}
