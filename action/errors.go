package action

import "errors"

var (
	ErrAllZonesNotAllowed         = errors.New("all regions is not allowed")
	ErrZonesRequired              = errors.New("regions is required")
	ErrAgeInHoursInvalid          = errors.New("age in hours is invalid")
	ErrCredentialsRequired        = errors.New("credentials are required")
	ErrProjectIDRequired          = errors.New("project id is required")
	ErrResourceLabelKeyRequired   = errors.New("resource label key is required")
	ErrResourceLabelValueRequired = errors.New("resource label value is required")
)
