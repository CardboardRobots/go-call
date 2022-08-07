package call

import (
	"net/url"

	"github.com/google/go-querystring/query"
)

func BuildUrl[T any](base string, options T) (string, error) {
	out, err := url.Parse(base)
	if err != nil {
		return "", err
	}

	queryOptions, err := query.Values(options)
	if err != nil {
		return "", err
	}

	out.RawQuery = queryOptions.Encode()

	return out.String(), nil
}
