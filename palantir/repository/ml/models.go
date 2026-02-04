package ml

type PredictionsClassification map[string]float64

type PredictionsClassificationsResponse struct {
	Predictions map[string]PredictionsClassification `json:"predictions"`
}

type PayloadRequest struct {
	Values []map[string]interface{} `json:"values"`
}
