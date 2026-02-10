package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"palantir/errors"
	"palantir/repository/ml"
	"palantir/service/prediction"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostPredictionSucess(t *testing.T) {

	type TestUnit struct {
		Name             string
		PayloadRequest   *PredictRequestPayload
		ExpectedResponse prediction.PredictionServiceResponse
		ExpectedError    error
	}

	tests := []TestUnit{
		{
			Name: "Teste 01 - Churn TMW: 3f55b86f-dc21-4ac8-8e89-7c1535359eaf",
			PayloadRequest: &PredictRequestPayload{
				ModelName: "Churn TMW",
				ID:        "3f55b86f-dc21-4ac8-8e89-7c1535359eaf",
			},
			ExpectedResponse: prediction.PredictionServiceResponse{
				Predictions: map[string]ml.PredictionsClassification{
					"3f55b86f-dc21-4ac8-8e89-7c1535359eaf": {
						"0": 0.6571428571428571,
						"1": 0.34285714285714286,
					},
				},
			},
			ExpectedError: nil,
		},
		{
			Name: "Teste 01 - Churn TMW: 564ab6da-2ad3-4c5f-9e59-60de8d89fbcc",
			PayloadRequest: &PredictRequestPayload{
				ModelName: "Churn TMW",
				ID:        "564ab6da-2ad3-4c5f-9e59-60de8d89fbcc",
			},
			ExpectedResponse: prediction.PredictionServiceResponse{
				Predictions: map[string]ml.PredictionsClassification{
					"564ab6da-2ad3-4c5f-9e59-60de8d89fbcc": {
						"0": 1,
						"1": 0,
					},
				},
			},
			ExpectedError: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			client := &http.Client{}

			jsonData, err := json.Marshal(tt.PayloadRequest)
			assert.NoError(t, err)

			bodyRequest := bytes.NewReader(jsonData)

			req, err := client.Post("http://localhost:3000/predict", "application/json", bodyRequest)
			assert.NoError(t, err)

			defer req.Body.Close()

			var response prediction.PredictionServiceResponse
			err = json.NewDecoder(req.Body).Decode(&response)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusOK, req.StatusCode)
			assert.Equal(t, tt.ExpectedResponse.Predictions, response.Predictions)

		})
	}

}

func TestPostPredictionIdNotFound(t *testing.T) {

	type TestUnit struct {
		Name             string
		PayloadRequest   *PredictRequestPayload
		ExpectedResponse prediction.PredictionServiceResponse
		ExpectedError    error
	}

	err := errors.ErrIdNotFound.Error()
	msgErr := &err

	tests := []TestUnit{
		{
			Name: "Teste 01 - Churn TMW: Id teste_not_exists",
			PayloadRequest: &PredictRequestPayload{
				ModelName: "Churn TMW",
				ID:        "teste_not_exists",
			},
			ExpectedResponse: prediction.PredictionServiceResponse{Err: msgErr},
			ExpectedError:    errors.ErrIdNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			client := &http.Client{}

			jsonData, err := json.Marshal(tt.PayloadRequest)
			assert.NoError(t, err)

			bodyRequest := bytes.NewReader(jsonData)

			resp, err := client.Post("http://localhost:3000/predict", "application/json", bodyRequest)
			assert.NoError(t, err)

			defer resp.Body.Close()

			response := prediction.PredictionServiceResponse{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			assert.Equal(t, tt.ExpectedResponse.Predictions, response.Predictions)

		})
	}

}

func TestPostPredictionModelNameNotFound(t *testing.T) {

	type TestUnit struct {
		Name             string
		PayloadRequest   *PredictRequestPayload
		ExpectedResponse prediction.PredictionServiceResponse
		ExpectedError    error
	}

	err := errors.ErrModelNotFound.Error()
	msgErr := &err

	tests := []TestUnit{
		{
			Name: "Teste 01 - Churn TMW: ModelName teste_not_exists",
			PayloadRequest: &PredictRequestPayload{
				ModelName: "Model Does Not Exist",
				ID:        "564ab6da-2ad3-4c5f-9e59-60de8d89fbcc",
			},
			ExpectedResponse: prediction.PredictionServiceResponse{Err: msgErr},
			ExpectedError:    errors.ErrIdNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			client := &http.Client{}

			jsonData, err := json.Marshal(tt.PayloadRequest)
			assert.NoError(t, err)

			bodyRequest := bytes.NewReader(jsonData)

			resp, err := client.Post("http://localhost:3000/predict", "application/json", bodyRequest)
			assert.NoError(t, err)

			defer resp.Body.Close()

			response := prediction.PredictionServiceResponse{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			assert.Equal(t, tt.ExpectedResponse.Predictions, response.Predictions)

		})
	}

}

func TestPostPredictionTagURINotFound(t *testing.T) {

	type TestUnit struct {
		Name             string
		PayloadRequest   *PredictRequestPayload
		ExpectedResponse prediction.PredictionServiceResponse
		ExpectedError    error
	}

	err := errors.ErrModelNotFound.Error()
	msgErr := &err

	tests := []TestUnit{
		{
			Name: "Teste 01 - Teste 01: TagURI teste_not_exists",
			PayloadRequest: &PredictRequestPayload{
				ModelName: "Teste 01",
				ID:        "564ab6da-2ad3-4c5f-9e59-60de8d89fbcc",
			},
			ExpectedResponse: prediction.PredictionServiceResponse{Err: msgErr},
			ExpectedError:    errors.ErrIdNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			client := &http.Client{}

			jsonData, err := json.Marshal(tt.PayloadRequest)
			assert.NoError(t, err)

			bodyRequest := bytes.NewReader(jsonData)

			resp, err := client.Post("http://localhost:3000/predict", "application/json", bodyRequest)
			assert.NoError(t, err)

			defer resp.Body.Close()

			response := &prediction.PredictionServiceResponse{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			assert.Equal(t, tt.ExpectedResponse.Predictions, response.Predictions)
			assert.Equal(t, *tt.ExpectedResponse.Err, *response.Err)

		})
	}

}

func TestPostPredictionTagFeatureStoreNotFound(t *testing.T) {

	type TestUnit struct {
		Name             string
		PayloadRequest   *PredictRequestPayload
		ExpectedResponse prediction.PredictionServiceResponse
		ExpectedError    error
	}

	err := errors.ErrFeatureStoreNotFound.Error()
	msgErr := &err

	tests := []TestUnit{
		{
			Name: "Teste 01: Tag Feature Store",
			PayloadRequest: &PredictRequestPayload{
				ModelName: "Teste 02",
				ID:        "564ab6da-2ad3-4c5f-9e59-60de8d89fbcc",
			},
			ExpectedResponse: prediction.PredictionServiceResponse{Err: msgErr},
			ExpectedError:    errors.ErrFeatureStoreNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			client := &http.Client{}

			jsonData, err := json.Marshal(tt.PayloadRequest)
			assert.NoError(t, err)

			bodyRequest := bytes.NewReader(jsonData)

			resp, err := client.Post("http://localhost:3000/predict", "application/json", bodyRequest)
			assert.NoError(t, err)

			defer resp.Body.Close()

			response := &prediction.PredictionServiceResponse{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			assert.Equal(t, tt.ExpectedResponse.Predictions, response.Predictions)
			assert.Equal(t, *tt.ExpectedResponse.Err, *response.Err)

		})
	}

}

func TestPostPredictionTagFeatureStoreDoesNotExist(t *testing.T) {

	type TestUnit struct {
		Name             string
		PayloadRequest   *PredictRequestPayload
		ExpectedResponse prediction.PredictionServiceResponse
		ExpectedError    error
	}

	err := errors.ErrFeatureStoreDoesNotExist.Error()
	msgErr := &err

	tests := []TestUnit{
		{
			Name: "Teste 01: Feature Store Does Not Exists",
			PayloadRequest: &PredictRequestPayload{
				ModelName: "Teste 03",
				ID:        "564ab6da-2ad3-4c5f-9e59-60de8d89fbcc",
			},
			ExpectedResponse: prediction.PredictionServiceResponse{Err: msgErr},
			ExpectedError:    errors.ErrFeatureStoreDoesNotExist,
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			client := &http.Client{}

			jsonData, err := json.Marshal(tt.PayloadRequest)
			assert.NoError(t, err)

			bodyRequest := bytes.NewReader(jsonData)

			resp, err := client.Post("http://localhost:3000/predict", "application/json", bodyRequest)
			assert.NoError(t, err)

			defer resp.Body.Close()

			response := &prediction.PredictionServiceResponse{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			assert.Equal(t, tt.ExpectedResponse.Predictions, response.Predictions)
			assert.Equal(t, *tt.ExpectedResponse.Err, *response.Err)

		})
	}

}

func TestPostPredictionModelURIDoesNotRespond(t *testing.T) {

	type TestUnit struct {
		Name             string
		PayloadRequest   *PredictRequestPayload
		ExpectedResponse prediction.PredictionServiceResponse
		ExpectedError    error
	}

	err := errors.ErrModelAPIDoesNotRespond{ModelURI: "http://ml_churn_app:5002/predict"}.Error()
	msgErr := &err

	tests := []TestUnit{
		{
			Name: "Teste 01: Model URI Does Not Respond",
			PayloadRequest: &PredictRequestPayload{
				ModelName: "Teste 04",
				ID:        "564ab6da-2ad3-4c5f-9e59-60de8d89fbcc",
			},
			ExpectedResponse: prediction.PredictionServiceResponse{Err: msgErr},
			ExpectedError:    errors.ErrModelAPIDoesNotRespond{ModelURI: "http://ml_churn_app:5002/predict"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {

			client := &http.Client{}

			jsonData, err := json.Marshal(tt.PayloadRequest)
			assert.NoError(t, err)

			bodyRequest := bytes.NewReader(jsonData)

			resp, err := client.Post("http://localhost:3000/predict", "application/json", bodyRequest)
			assert.NoError(t, err)

			defer resp.Body.Close()

			response := &prediction.PredictionServiceResponse{}
			err = json.NewDecoder(resp.Body).Decode(&response)
			assert.NoError(t, err)
			assert.Equal(t, http.StatusInternalServerError, resp.StatusCode)
			assert.Equal(t, tt.ExpectedResponse.Predictions, response.Predictions)
			assert.Equal(t, *tt.ExpectedResponse.Err, *response.Err)

		})
	}

}
