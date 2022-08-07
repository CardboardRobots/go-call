package call

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

type CallOptions struct {
	Method  string
	Base    string
	Headers HeaderMap
	Query   any
	Body    any
}

func (c *CallOptions) GetUrl() (string, error) {
	return BuildUrl(c.Base, c.Query)
}

func (c *CallOptions) GetBody() (io.Reader, error) {
	data, err := json.Marshal(c.Body)
	if err != nil {
		return nil, err
	}
	return bytes.NewBuffer(data), nil
}

func (c *CallOptions) GetRequest() (*http.Request, error) {
	url, err := c.GetUrl()
	if err != nil {
		return nil, err
	}

	body, err := c.GetBody()
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest(c.Method, url, body)
	if err != nil {
		return nil, err
	}

	c.updateHeader(req)

	return req, nil
}

func (c *CallOptions) updateHeader(req *http.Request) {
	for key, value := range c.Headers {
		req.Header.Add(key, value)
	}
}

func Call[R any](
	options CallOptions,
	callback func(resp *http.Response, bytes []byte) (R, error),
) (R, error) {
	var zero R

	req, err := options.GetRequest()
	if err != nil {
		return zero, err
	}

	httpClient := http.DefaultClient
	resp, err := httpClient.Do(req)
	if err != nil {
		return zero, err
	}

	bytes, err := Read(resp)
	if err != nil {
		return zero, err
	}

	return callback(resp, bytes)
}

func NewCallOptions(options ...CallOptionFunction) CallOptions {
	c := CallOptions{
		Method:  "Get",
		Headers: HeaderMap{},
	}
	for _, option := range options {
		option(&c)
	}
	return c
}

type CallOptionFunction func(c *CallOptions)

func WithMethod(method string) CallOptionFunction {
	return func(c *CallOptions) {
		c.Method = method
	}
}

func WithBase(base string) CallOptionFunction {
	return func(c *CallOptions) {
		c.Base = base
	}
}

func WithQuery(q any) CallOptionFunction {
	return func(c *CallOptions) {
		c.Query = q
	}
}

func WithBody(b any) CallOptionFunction {
	return func(c *CallOptions) {
		c.Body = b
	}
}

func WithHeader(name string, value string) CallOptionFunction {
	return func(c *CallOptions) {
		c.Headers[name] = value
	}
}

func FromBytes[T any](data []byte) (*T, error) {
	var result T
	err := json.Unmarshal(data, result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func Read(resp *http.Response) ([]byte, error) {
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return bytes, nil
}
