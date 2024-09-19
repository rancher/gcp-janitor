package main

import (
	"context"

	"github.com/rancher-sandbox/gcp-janitor/action"
)

func main() {
	action.Log("running gcp janitor")

	input, err := action.NewInput()
	if err != nil {
		action.LogErrorAndExit("failed to get input: %s", err.Error())
	}

	if err := input.Validate(); err != nil {
		action.LogErrorAndExit("failed input validation: %s", err.Error())
	}

	a := action.New(input.DryRun)

	ctx := context.Background()
	if err := a.Cleanup(ctx, input); err != nil {
		action.LogErrorAndExit("failed to cleanup gcp resources: %s", err.Error())
	}
}
