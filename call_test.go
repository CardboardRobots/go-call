package call

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestWithMethod(t *testing.T) {
	want := "Post"

	c := NewCallOptions(
		WithMethod(want),
	)

	if c.Method != want {
		t.Errorf("Received: %v, Expected: %v", c.Base, want)
	}
}

func TestWithBase(t *testing.T) {
	want := "http://localhost"

	c := NewCallOptions(
		WithBase(want),
	)

	if c.Base != want {
		t.Errorf("Received: %v, Expected: %v", c.Base, want)
	}
}

func TestWithQuery(t *testing.T) {
	type TestQuery struct {
		Name string `url:"name"`
	}

	want := TestQuery{}

	c := NewCallOptions(
		WithQuery(want),
	)

	if c.Query != want {
		t.Errorf("Received: %v, Expected: %v", c.Query, want)
	}
}

func TestWithBody(t *testing.T) {
	type TestBody struct {
		Name string
	}

	want := TestBody{}

	c := NewCallOptions(
		WithBody(want),
	)

	if c.Body != want {
		t.Errorf("Received: %v, Expected: %v", c.Body, want)
	}
}

func TestWithHeader(t *testing.T) {
	name := "Content-Type"
	want := "application/json"

	c := NewCallOptions(
		WithHeader(name, want),
	)

	if c.Headers[name] != want {
		t.Errorf("Received: %v, Expected: %v", c.Headers[name], want)
	}
}

func TestCall(t *testing.T) {
	svc := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	}))

	type TestQuery struct {
		Name string `url:"name"`
	}

	result, err := Call(NewCallOptions(
		WithBase(svc.URL),
		WithHeader("Content-Type", "application/json"),
		WithQuery(TestQuery{
			Name: "test",
		}),
	), func(resp *http.Response, bytes []byte) (string, error) {
		return "", nil
	})

	if err != nil {
		t.Errorf("%v", err)
	}

	if result != "" {
		t.Errorf("%v", result)
	}
}
