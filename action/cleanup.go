package action

import (
	"context"
)

type CleanupScope struct {
	CredentialsJSON    []byte
	Zones              []string
	ProjectID          string
	AgeInHours         int
	DryRun             bool
	ResourceLabelKey   string
	ResourceLabelValue string
}

type CleanupFunc func(ctx context.Context, input *CleanupScope) error
