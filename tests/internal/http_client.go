package internal

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/http/httptrace"
)

var ErrNoRequest = errors.New("trying to emit empty request")

type HttpClient struct {
	client        *http.Client
	trace         *httptrace.ClientTrace
	initialClient *http.Client

	// Current storage
	request      *http.Request
	httpResponse *http.Response
	Response     *Response
	tracing      bool
}

func (cli *HttpClient) EmitRequest(req RequestPreparation) (err error) {
	if req.Empty() {
		return ErrNoRequest
	}

	cli.request, err = req.GenerateRequest()
	if err != nil {
		return err
	}

	if cli.tracing {
		cli.request = cli.request.WithContext(
			httptrace.WithClientTrace(cli.request.Context(), cli.trace),
		)
	}

	// nolint: bodyclose
	cli.httpResponse, err = cli.client.Do(cli.request)
	if err != nil {
		return
	}

	body, errBody := ioutil.ReadAll(cli.httpResponse.Body)

	if errBody != nil {
		return errBody
	}

	cli.Response = NewResponse(cli.httpResponse.StatusCode, body, cli.httpResponse.Cookies(), cli.httpResponse.Header)

	return
}
