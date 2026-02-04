package mlflow

type RegisteredModel struct {
	Name                 string               `json:"name"`
	CreationTimestamp    int                  `json:"creation_timestamp"`
	LastUpdatedTimestamp int                  `json:"last_updated_timestamp"`
	LatestVersions       ModelVersions        `json:"latest_versions"`
	Tags                 []RegisteredModelTag `json:"tags"`
}

type RegisteredModelTag struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ModelVersion struct {
	Name                 string `json:"name"`
	Version              string `json:"version"`
	CreationTimestamp    int    `json:"creation_timestamp"`
	LastUpdatedTimestamp int    `json:"last_updated_timestamp"`
	CurrentStage         string `json:"current_stage"`
	Description          string `json:"description"`
	Source               string `json:"source"`
	RunId                string `json:"run_id"`
	Status               string `json:"status"`
	RunLink              string `json:"run_link"`
}

type ModelVersions []ModelVersion

type GetLatestsModelVersionsResponse struct {
	ModelVersions ModelVersions `json:"model_versions"`
}

type GetRegisteredModelResponse struct {
	RegisteredModel RegisteredModel `json:"registered_model"`
}
