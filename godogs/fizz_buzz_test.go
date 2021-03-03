package godogs

import (
	"github.com/cucumber/godog"

	"godogs/internal"
)

func InitializeScenario(ctx *godog.ScenarioContext) {
	ctx.Step(`^I reset client$`, internal.ResetClient)
	ctx.Step(
		`^I GET (.*)$`,
		func(method, endpoint string) error {
			client := internal.GetClient()
			client.SetEndpoint(endpoint)

			// InitRequest mandatory cause we wish to ensure the call is on the correct value
			return client.ExecuteRequest()
		},
	)
	ctx.Step(`(?:I )?set request query$`, internal.SetQueryParams)
	ctx.Step(`^response should contain$`, internal.ResponseShouldBeEquivalent)
}
