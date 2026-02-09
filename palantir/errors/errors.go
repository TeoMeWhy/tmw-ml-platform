package errors

import "errors"

var (
	ErrModelNotFound        = errors.New("model not found in model tags")
	ErrFeatureStoreNotFound = errors.New("feature store not found in model tags")
	ErrInvalidModelURI      = errors.New("invalid model URI")
	ErrPredictionFailed     = errors.New("prediction failed")
	ErrIdNotFound           = errors.New("id not found in feature store")
)
