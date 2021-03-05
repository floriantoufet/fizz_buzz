package internal

import (
	"bytes"
	"context"
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
func (request RequestPreparation) GenerateRequest() (*http.Request, error) { //nolint:gocyclo
	var b bytes.Buffer

	req, _ := http.NewRequestWithContext(context.Background(), request.Method, request.Endpoint, &b) //nolint:errcheck

	q := req.URL.Query()

	for key, value := range request.Arguments {
		q.Add(key, value)
	}

	req.URL.RawQuery = q.Encode()

	return req, nil
}

func (request RequestPreparation) SetEndpoint(method, endpoint string) RequestPreparation {
	request.Endpoint = endpoint
	request.Method = method

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
