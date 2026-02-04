package prediction

import (
	"fmt"
	"net/http"
	featurestore "palantir/repository/feature_store"
	"palantir/repository/ml"
	"palantir/repository/mlflow"
)

type PredictionService struct {
	MLFlowRepo       *mlflow.MLFlowRepository
	FeatureStoreRepo *featurestore.FeatureStoreRepository
}

func (s *PredictionService) Predict(modelName string, ids []string) (ml.PredictionsClassificationsResponse, error) {

	model, err := s.MLFlowRepo.GetRegisteredModel(modelName)
	if err != nil {
		return ml.PredictionsClassificationsResponse{}, err
	}

	URI := GetTagByKey(model.Tags, "uri")
	if URI == "" {
		return ml.PredictionsClassificationsResponse{}, fmt.Errorf("error: key 'uri' not found on the registered model on MLFlow")
	}

	featureStoreName := GetTagByKey(model.Tags, "feature_store")
	if featureStoreName == "" {
		return ml.PredictionsClassificationsResponse{}, fmt.Errorf("error: key 'feature_store' not found on the registered model on MLFlow")
	}

	data, err := s.FeatureStoreRepo.GetFeatures(featureStoreName, ids)
	if err != nil {
		return ml.PredictionsClassificationsResponse{}, err
	}

	return s.PredictData(URI, data)

}

func (s *PredictionService) PredictData(modelURI string, data []map[string]interface{}) (ml.PredictionsClassificationsResponse, error) {

	httpClient := &http.Client{}
	mlRepo := ml.NewMLRepository(modelURI, httpClient)

	predictions, err := mlRepo.GetPredictions(data)
	if err != nil {
		return ml.PredictionsClassificationsResponse{}, err
	}

	return predictions, nil
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
