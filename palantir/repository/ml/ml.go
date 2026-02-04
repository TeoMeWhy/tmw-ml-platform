package ml

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type MLRepository struct {
	URL        string
	HttpClient *http.Client
}

func (r *MLRepository) GetPredictions(data []map[string]interface{}) (PredictionsClassificationsResponse, error) {

	predictons := PredictionsClassificationsResponse{}

	payload := PayloadRequest{
		Values: data,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		return predictons, err
	}

	bodyReq := bytes.NewReader(jsonData)
	req, err := r.HttpClient.Post(r.URL, "application/json", bodyReq)
	if err != nil {
		return predictons, err
	}

	defer req.Body.Close()

	if req.StatusCode != http.StatusOK {
		return predictons, nil
	}

	if err := json.NewDecoder(req.Body).Decode(&predictons); err != nil {
		return predictons, err
	}

	return predictons, nil
}

func NewMLRepository(url string, client *http.Client) *MLRepository {
	return &MLRepository{
		URL:        url,
		HttpClient: client,
	}
}
