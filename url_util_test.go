package call

import (
	"testing"
)

func TestBuildUrl(t *testing.T) {
	type TestQuery struct {
		Name string `url:"name"`
	}

	url, err := BuildUrl("http://localhost:8080", TestQuery{Name: "test"})
	if err != nil {
		t.Errorf("%v", err)
	}

	want := "http://localhost:8080?name=test"
	if url != want {
		t.Errorf("Received: %v, Expected: %v", url, want)
	}
}
