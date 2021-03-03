package internal

import (
	"errors"

	"github.com/cucumber/godog"
)

var client *Client

type Client struct {
	cli *HttpClient

	request RequestPreparation
}

func GetClient() *Client {
	if client == nil {
		client = &Client{}
	}
	return client
}

func ResetClient() error {
	client = &Client{}

	return nil
}

// ExecuteRequest builds and executes request through http client.
func (cli *Client) ExecuteRequest() error {
	if err := cli.cli.EmitRequest(cli.request); err != nil {
		return err
	}

	return nil
}

// SetQueryParams replaces query parameters with new ones.
func SetQueryParams(args *godog.Table) error {
	cli := GetClient()
	if cli.request.Empty() {
		cli.InitRequest()
	}

	cli.request = cli.request.ResetArguments()
	cli.AddQueryParams(args)

	return nil
}

func (cli *Client) SetEndpoint(endpoint string) {
	if cli.request.Empty() {
		cli.InitRequest()
	}

	cli.request = cli.request.SetEndpoint(endpoint)
}

// AddQueryParams adds query parameters.
func (cli *Client) AddQueryParams(args *godog.Table) {
	for _, row := range args.Rows {
		cli.AddQueryParam(row.Cells[0].Value, row.Cells[1].Value)
	}
}

// AddQueryParam adds a single query parameter.
func (cli *Client) AddQueryParam(key, val string) {
	if cli.request.Empty() {
		cli.InitRequest()
	}

	cli.request.AddArgument(key, val)
}

// InitRequest starts a new request with default parameter.
func (cli *Client) InitRequest() {
	cli.request = RequestPreparation{
		Method:    "GET",
		Arguments: make(map[string]string),
	}
}

// ResponseShouldBeEquivalent asserts response body is a HTML resembling provided.
func ResponseShouldBeEquivalent(body *godog.DocString) error {
	cli := GetClient()
	if cli == nil {
		return errors.New("empty client")
	}
	return cli.cli.Response.HTMLResemble(body)
}
