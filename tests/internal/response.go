package internal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/cucumber/godog"
	"github.com/google/go-cmp/cmp"
)

type Response struct {
	Body    []byte
	Cookies map[string]*http.Cookie
	Headers http.Header
	Status  int
}

func NewResponse(status int, body []byte, cookies []*http.Cookie, headers http.Header) *Response {
	cookieMap := make(map[string]*http.Cookie)

	for _, cookie := range cookies {
		cookieMap[cookie.Name] = cookie
	}

	return &Response{
		Status:  status,
		Body:    body,
		Cookies: cookieMap,
		Headers: headers,
	}
}

func (r Response) HTMLResemble(expectedBody *godog.DocString) error {
	var expected, actual interface{}

	expected = expectedBody.Content
	actual = string(r.Body)

	// the matching may be adapted per different requirements.
	if expected != actual {
		return fmt.Errorf("%w, %v vs. %v", errors.New("no match"), expected, actual)
	}

	return nil
}

func ResponseJSONShouldBeEquivalent(expected *godog.DocString) error {
	cli := GetClient()

	return cli.cli.Response.JSONResemble(expected)
}

func (r Response) JSONResemble(expectedBody *godog.DocString) error {
	var expected, actual interface{}
	var err error

	// re-encode expected response
	if err = json.Unmarshal([]byte(expectedBody.Content), &expected); err != nil {
		return err
	}

	// re-encode actual response too
	if err = json.Unmarshal(r.Body, &actual); err != nil {
		return err
	}

	// the matching may be adapted per different requirements.
	if !cmp.Equal(expected, actual) {
		return fmt.Errorf("%w, %v vs. %v", errors.New("no match"), expected, actual)
	}

	return nil
}

// ResponseHasStatus asserts Response has expected status.
func ResponseHasStatus(expectedStatus int) error {
	cli := GetClient()
	if cli.cli.Response.Status != expectedStatus {
		return fmt.Errorf("%w: expected %d - got %d", errors.New("status code does not match expected"), expectedStatus, cli.cli.Response.Status)
	}

	return nil
}
