package ml

type Predictions map[string]float64

type PredictionsInstances []Predictions

type PayloadRequest struct {
	Values []map[string]interface{} `json:"values"`
}
