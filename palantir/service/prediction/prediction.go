package prediction

import (
	"fmt"
	"net/http"
	featurestore "pml/repository/feature_store"
	"pml/repository/ml"
	"pml/repository/mlflow"
)

type PredictionService struct {
	MLFlowRepo       *mlflow.MLFlowRepository
	FeatureStoreRepo *featurestore.FeatureStoreRepository
}

func (s *PredictionService) Predict(modelName string, ids []string) (ml.PredictionsInstances, error) {

	model, err := s.MLFlowRepo.GetRegisteredModel(modelName)
	if err != nil {
		return ml.PredictionsInstances{}, err
	}

	URI := GetTagByKey(model.Tags, "uri")
	if URI == "" {
		return ml.PredictionsInstances{}, fmt.Errorf("error: key 'uri' not found on the registered model on MLFlow")
	}

	featureStoreName := GetTagByKey(model.Tags, "feature_store")
	if featureStoreName == "" {
		return ml.PredictionsInstances{}, fmt.Errorf("error: key 'feature_store' not found on the registered model on MLFlow")
	}

	data, err := s.FeatureStoreRepo.GetFeatures(featureStoreName, ids)
	if err != nil {
		return ml.PredictionsInstances{}, err
	}

	return s.PredictData(URI, data)

}

func (s *PredictionService) PredictData(modelURI string, data []map[string]interface{}) (ml.PredictionsInstances, error) {

	httpClient := &http.Client{}
	mlRepo := ml.NewMLRepository(modelURI, httpClient)

	predictions, err := mlRepo.GetPredictions(data)
	if err != nil {
		return ml.PredictionsInstances{}, err
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
