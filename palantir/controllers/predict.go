package controllers

import (
	"fmt"
	"log"
	"net/http"
	"palantir/configs"
	featurestore "palantir/repository/feature_store"
	"palantir/repository/mlflow"
	"palantir/service/prediction"

	"github.com/gofiber/fiber/v3"
)

type PredictRequestPayload struct {
	ModelName string `json:"model_name"`
	ID        string `json:"id"`
}

type PredictController struct {
	PredictionService *prediction.PredictionService
}

func (c *PredictController) PostPrediction(ctx fiber.Ctx) error {

	payloadRequest := &PredictRequestPayload{}
	if err := ctx.Bind().Body(&payloadRequest); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "invalid request payload")
	}

	pred, err := c.PredictionService.Predict(payloadRequest.ModelName, []string{payloadRequest.ID})
	if err != nil {
		return ctx.Status(http.StatusInternalServerError).JSON(pred)
	}

	return ctx.Status(http.StatusOK).JSON(pred)
}

func NewPredictController(cfg *configs.Config) (*PredictController, error) {

	mlflowRepo := mlflow.NewMLFlowRepository(cfg)

	featureStoreRepo, err := featurestore.NewFeatureStoreRepository(cfg)
	if err != nil {
		txtError := fmt.Sprintf("Erro ao conectar ao Feature Store: %v", err)
		log.Println(txtError)
		return nil, err
	}

	predService := prediction.NewPredictionService(mlflowRepo, featureStoreRepo)
	ctrlr := &PredictController{PredictionService: predService}
	return ctrlr, nil
}
