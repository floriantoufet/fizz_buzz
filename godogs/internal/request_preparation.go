package internal

import (
	"net/http"

	"github.com/google/go-cmp/cmp"
)

type RequestPreparation struct {
	JSONBody  *string
	Arguments map[string]string
	Endpoint  string
	Method    string
}

func (request RequestPreparation) Empty() bool {
	return cmp.Equal(request, RequestPreparation{})
}

// Generate request
func (request RequestPreparation) GenerateRequest(jar http.CookieJar) (*http.Request, error) { //nolint:gocyclo
	var req *http.Request

	q := req.URL.Query()

	for key, value := range request.Arguments {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (request RequestPreparation) SetEndpoint(endpoint string) RequestPreparation {
	request.Endpoint = endpoint
	return request
}

func (request RequestPreparation) ResetArguments() RequestPreparation {
	request.Arguments = make(map[string]string)
	return request
}

func (request RequestPreparation) AddArgument(key, value string) RequestPreparation {
	request.Arguments[key] = value
	return request
}
