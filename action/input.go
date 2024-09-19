package action

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"go.uber.org/multierr"
)

type Input struct {
	CredentialsJSON    string `env:"INPUT_CREDENTIALS-JSON"`
	Zones              string `env:"INPUT_ZONES"`
	ProjectID          string `env:"INPUT_PROJECT-ID"`
	ResourceLabelKey   string `env:"INPUT_RESOURCE-LABEL-KEY"`
	ResourceLabelValue string `env:"INPUT_RESOURCE-LABEL-VALUE"`
	AgeInHours         int    `env:"INPUT_AGE-IN-HOURS"`
	DryRun             bool   `env:"INPUT_DRY-RUN"`
}

// NewInput creates a new input from the environment variables.
func NewInput() (*Input, error) {
	input := &Input{}
	if err := env.Parse(input); err != nil {
		return nil, fmt.Errorf("parsing environment variables: %w", err)
	}

	return input, nil
}

func (i *Input) Validate() error {
	var err error

	if len(i.CredentialsJSON) == 0 {
		err = multierr.Append(err, ErrCredentialsRequired)
	}

	if i.Zones == "" {
		err = multierr.Append(err, ErrZonesRequired)
	}

	if i.ProjectID == "" {
		err = multierr.Append(err, ErrProjectIDRequired)
	}

	if i.ResourceLabelKey == "" {
		err = multierr.Append(err, ErrResourceLabelKeyRequired)
	}

	if i.ResourceLabelValue == "" {
		err = multierr.Append(err, ErrResourceLabelValueRequired)
	}

	if i.AgeInHours < 0 {
		err = multierr.Append(err, ErrAgeInHoursInvalid)
	}

	return err
}
