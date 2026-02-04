package main

import (
	"log"
	"palantir/configs"
	featurestore "palantir/repository/feature_store"
	"palantir/repository/mlflow"
	"palantir/service/prediction"
)

func main() {

	config, err := configs.LoadConfig()
	if err != nil {
		log.Fatal("Deu merda!", err)
	}

	mlflowRepo := mlflow.NewMLFlowRepository(config)
	featureStoreRepo, err := featurestore.NewFeatureStoreRepository(config)
	if err != nil {
		log.Fatalf("Erro ao conectar ao Feature Store: %v", err)
	}

	predictionService := prediction.NewPredictionService(mlflowRepo, featureStoreRepo)

	model := "Churn TMW"
	ids := []string{
		"000ff655-fa9f-4baa-a108-47f581ec52a1",
		"3f55b86f-dc21-4ac8-8e89-7c1535359eaf",
	}

	result, err := predictionService.Predict(model, ids)
	if err != nil {
		log.Fatalf("Erro ao obter predições: %v", err)
	}

	log.Printf("Predições recebidas: %+v", result)

}
