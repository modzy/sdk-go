package modzy

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/peterhellberg/link"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

type Client interface {
	WithAPIKey(apiKey string) Client
	WithTeamKey(teamID string, token string) Client
	WithOptions(opts ...ClientOption) Client
	Jobs() JobsClient
	// for python sdk parity:
	// Models() ModelsClient
	// Tags() TagsClient
}

type standardClient struct {
	baseURL                string
	httpClient             *http.Client
	requestDebugging       bool
	responseDebugging      bool
	authorizationDecorator requestDecorator
	jobsClient             *standardJobsClient
}

func NewClient(baseURL string, opts ...ClientOption) Client {
	var client = &standardClient{
		baseURL:    baseURL,
		httpClient: defaultHTTPClient,
	}
	client.WithOptions(opts...)

	// setup our namespaced groupings that all share this as their base client
	client.jobsClient = &standardJobsClient{
		baseClient: client,
	}

	return client
}

// WithAPIKey -
func (c *standardClient) WithAPIKey(apiKey string) Client {
	c.authorizationDecorator = func(req *http.Request) *http.Request {
		req.Header.Add("Authorization", fmt.Sprintf("ApiKey %s", apiKey))
		return req
	}
	return c
}

// WithTeamKey -
func (c *standardClient) WithTeamKey(teamID string, token string) Client {
	c.authorizationDecorator = func(req *http.Request) *http.Request {
		req.Header.Add("Modzy-Team-Id", teamID)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		return req
	}
	return c
}

// WithOptions -
func (c *standardClient) WithOptions(opts ...ClientOption) Client {
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Jobs returns a client for acess all job related API functions
func (c *standardClient) Jobs() JobsClient {
	return c.jobsClient
}

func (c *standardClient) execute(
	ctx context.Context,
	path string, method string, toPostStruct interface{}, into interface{},
) (*http.Response, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)

	var toPost io.Reader
	if toPostStruct != nil {
		toPostBytes, err := json.Marshal(toPostStruct)
		if err != nil {
			return nil, errors.WithMessagef(err, "failed to marshal provided body to %s:%s", method, path)
		}
		toPost = bytes.NewReader(toPostBytes)
	}

	req, err := http.NewRequest(method, url, toPost)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to create request to %s:%s", method, path)
	}
	req = req.WithContext(ctx)
	req.Header.Add("Content-Type", "application/json")

	if c.authorizationDecorator != nil {
		c.authorizationDecorator(req)
	}

	if c.requestDebugging {
		// jsonize again if debugging
		debugJson, _ := json.Marshal(toPostStruct)

		logrus.WithFields(logrus.Fields{
			"url":    req.URL,
			"method": method,
			"body":   string(debugJson),
		}).Debug("API request")
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, errors.WithMessagef(err, "failed to executing request to %s:%s", method, path)
	}
	defer resp.Body.Close()

	var toDecode io.Reader = resp.Body

	if c.responseDebugging {
		body, _ := ioutil.ReadAll(resp.Body)
		logrus.WithFields(logrus.Fields{
			"method":     req.Method,
			"url":        req.URL,
			"statusCode": resp.StatusCode,
			// "headers":    resp.Header,
			"body": string(body),
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
	} else if resp.StatusCode == 204 {
		// No Content
	} else {
		if err := json.NewDecoder(toDecode).Decode(into); err != nil {
			return resp, errors.WithMessagef(err, "failed parsing the response from %s:%s", method, path)
		}
	}

	return resp, nil
}

func (c *standardClient) get(ctx context.Context, path string, into interface{}) (*http.Response, error) {
	return c.execute(ctx, path, "GET", nil, into)
}

func (c *standardClient) list(ctx context.Context, path string, paging PagingInput, into interface{}) (*http.Response, link.Group, error) {
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

	resp, err := c.get(ctx, partialUrl.String(), into)
	if err != nil {
		return resp, link.Group{}, err
	}

	// parse out the link response header
	links := link.Parse(resp.Header.Get("Link"))

	return resp, links, err
}

func (c *standardClient) post(ctx context.Context, path string, toPost interface{}, into interface{}) (*http.Response, error) {
	return c.execute(ctx, path, "POST", toPost, into)
}

func (c *standardClient) delete(ctx context.Context, path string, into interface{}) (*http.Response, error) {
	return c.execute(ctx, path, "DELETE", nil, into)
}
