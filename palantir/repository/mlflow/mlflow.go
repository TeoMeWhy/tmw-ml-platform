package mlflow

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MLFlowRepository struct {
	BaseURL    string
	HttpClient *http.Client
}

func (r *MLFlowRepository) GetLatestsModelVersions(modelName string) (ModelVersions, error) {

	mVersions := ModelVersions{}
	bodyMap := bytes.NewReader([]byte(`{"name":"` + modelName + `"}`))

	resp, err := r.HttpClient.Post(r.BaseURL+"/api/2.0/mlflow/registered-models/get-latest-versions", "application/json", bodyMap)
	if err != nil {
		return mVersions, err
	}

	if resp.StatusCode != http.StatusOK {
		return mVersions, fmt.Errorf("error: %s", resp.Status)
	}

	responsePayload := GetLatestsModelVersionsResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&responsePayload); err != nil {
		return mVersions, err
	}

	return responsePayload.ModelVersions, nil
}

func (r *MLFlowRepository) GetRegisteredModel(modelName string) (RegisteredModel, error) {

	registeredModel := RegisteredModel{}

	req, err := http.NewRequest("GET", r.BaseURL+"/api/2.0/mlflow/registered-models/get", nil)
	if err != nil {
		return registeredModel, err
	}

	q := req.URL.Query()
	q.Add("name", modelName)
	req.URL.RawQuery = q.Encode()

	resp, err := r.HttpClient.Do(req)
	if err != nil {
		return registeredModel, err
	}

	if resp.StatusCode != http.StatusOK {
		return registeredModel, fmt.Errorf("error: %s", resp.Status)
	}

	defer resp.Body.Close()

	responsePayload := GetRegisteredModelResponse{}
	if err := json.NewDecoder(resp.Body).Decode(&responsePayload); err != nil {
		return registeredModel, err
	}

	return responsePayload.RegisteredModel, nil
}

func NewMLFlowRepository(baseURL string) *MLFlowRepository {

	client := &http.Client{}

	return &MLFlowRepository{
		BaseURL:    baseURL,
		HttpClient: client,
	}
}
