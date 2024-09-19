package action

import (
	"context"
	"fmt"
	"strings"
)

type GCPJanitorAction interface {
	Cleanup(ctx context.Context, input *Input) error
}

func New(dryrun bool) GCPJanitorAction {
	return &action{
		dryrun: dryrun,
	}
}

type action struct {
	dryrun bool
}

type Cleaner struct {
	Service string
	Run     CleanupFunc
}

func (a *action) Cleanup(ctx context.Context, input *Input) error {
	cleaners := []Cleaner{
		{Service: "computeVMs", Run: a.cleanVMs},
	}

	inputZones := strings.Split(input.Zones, ",")

	for _, cleaner := range cleaners {
		scope := &CleanupScope{
			CredentialsJSON:    []byte(input.CredentialsJSON),
			Zones:              inputZones,
			ProjectID:          input.ProjectID,
			ResourceLabelKey:   input.ResourceLabelKey,
			ResourceLabelValue: input.ResourceLabelValue,
		}
		if err := cleaner.Run(ctx, scope); err != nil {
			return fmt.Errorf("failed to clean %s: %w", cleaner.Service, err)
		}
	}

	return nil
}
