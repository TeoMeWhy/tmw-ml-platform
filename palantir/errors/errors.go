package errors

import (
	"errors"
	"fmt"
)

var (
	ErrModelNotFound            = errors.New("model not found in model tags")
	ErrFeatureStoreNotFound     = errors.New("feature store not found in model tags")
	ErrInvalidModelURI          = errors.New("invalid model URI")
	ErrPredictionFailed         = errors.New("prediction failed")
	ErrIdNotFound               = errors.New("id not found in feature store")
	ErrFeatureStoreDoesNotExist = errors.New("the feature store does not exist to this model")
)

type ErrModelAPIDoesNotRespond struct {
	ModelURI string
}

func (e ErrModelAPIDoesNotRespond) Error() string {
	return fmt.Sprintf("model URI '%s' does not respond", e.ModelURI)
}
