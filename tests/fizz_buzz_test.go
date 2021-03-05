package tests

import (
	"github.com/cucumber/godog"

	"github.com/floriantoufet/fizzbuzz/tests/internal"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I reset HTTP client$`, internal.ResetClient)
	ctx.Step(
		`^I get fizzBuzz$`,
		func() error {
			client := internal.GetClient()
			client.SetEndpoint("GET", "http://localhost:8080/v1/fizz_buzz")

			return client.ExecuteRequest()
		},
	)
	ctx.Step(
		`^I get fizzBuzz requests stats$`,
		func() error {
			client := internal.GetClient()
			client.SetEndpoint("GET", "http://localhost:8080/v1/stats")

			return client.ExecuteRequest()
		},
	)
	ctx.Step(`^response status code should be (\d+)$`, internal.ResponseHasStatus)
	ctx.Step(`(?:I )?set request query$`, internal.SetQueryParams)
	ctx.Step(`^json response should resemble$`, internal.ResponseJSONShouldBeEquivalent)
	ctx.Step(`^I reset fizzBuzz request stats$`, func() error {
		client := internal.GetClient()
		client.SetEndpoint("DELETE", "http://localhost:8080/v1/stats")

		return client.ExecuteRequest()
	})
}
