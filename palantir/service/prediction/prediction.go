package prediction

import (
	"net/http"
	"palantir/errors"
	featurestore "palantir/repository/feature_store"
	"palantir/repository/ml"
	"palantir/repository/mlflow"
)

type PredictionService struct {
	MLFlowRepo       *mlflow.MLFlowRepository
	FeatureStoreRepo *featurestore.FeatureStoreRepository
}

type PredictionServiceResponse struct {
	Predictions map[string]ml.PredictionsClassification `json:"predictions"`
	Err         *string                                 `json:"error,omitempty"`
}

func (s *PredictionService) Predict(modelName string, ids []string) (PredictionServiceResponse, error) {

	model, err := s.MLFlowRepo.GetRegisteredModel(modelName)
	if err != nil {
		errMsg := err.Error()
		return PredictionServiceResponse{Err: &errMsg}, err
	}

	URI := GetTagByKey(model.Tags, "uri")
	if URI == "" {
		err := errors.ErrModelNotFound
		errMsg := err.Error()
		return PredictionServiceResponse{Err: &errMsg}, err
	}

	featureStoreName := GetTagByKey(model.Tags, "feature_store")
	if featureStoreName == "" {
		err := errors.ErrFeatureStoreNotFound
		errMsg := err.Error()
		return PredictionServiceResponse{Err: &errMsg}, err
	}

	data, err := s.FeatureStoreRepo.GetFeatures(featureStoreName, ids)
	if err != nil {
		errMsg := err.Error()
		return PredictionServiceResponse{Err: &errMsg}, err
	}

	if len(data) == 0 {
		err := errors.ErrIdNotFound
		errMsg := err.Error()
		return PredictionServiceResponse{Err: &errMsg}, err
	}

	pred, err := s.PredictData(URI, data)
	if err != nil {
		errMsg := err.Error()
		return PredictionServiceResponse{Err: &errMsg}, err
	}

	return PredictionServiceResponse{Predictions: pred.Predictions}, nil
}

func (s *PredictionService) PredictData(modelURI string, data []map[string]interface{}) (ml.PredictionsClassificationsResponse, error) {
	httpClient := &http.Client{}
	mlRepo := ml.NewMLRepository(modelURI, httpClient)
	return mlRepo.GetPredictions(data)
}

func NewPredictionService(mlflowRepo *mlflow.MLFlowRepository, featureStoreRepo *featurestore.FeatureStoreRepository) *PredictionService {
	return &PredictionService{
		MLFlowRepo:       mlflowRepo,
		FeatureStoreRepo: featureStoreRepo,
	}
}

func GetTagByKey(tags []mlflow.RegisteredModelTag, key string) string {
	for _, tag := range tags {
		if tag.Key == key {
			return tag.Value
		}
	}
	return ""
}
