package internal

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/cucumber/godog"
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
