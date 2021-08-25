package model

type ResourcesProcessingModel struct {
	Identifier           string                                  `json:"identifier"`
	Version              string                                  `json:"version"`
	DeployedAt           ModzyTime                               `json:"deployedAt"`
	Engines              []ResourcesProcessingModelEngine        `json:"engines"`
	Inputs               ResourcesProcessingModelInputs          `json:"inputs"`
	Jobs                 ResourcesProcessingModelJobs            `json:"jobs"`
	Situations           []string                                `json:"situations"`
	ModelDeploymentState ResourcesProcessingModelDeploymentState `json:"modelDeploymentState"`
}

type ResourcesProcessingModelEngine struct {
	Name       string                                    `json:"name"`
	CreatedAt  ModzyTime                                 `json:"createdAt"`
	Ready      bool                                      `json:"ready"`
	Conditions []ResourcesProcessingModelEngineCondition `json:"conditions"`
}

type ResourcesProcessingModelEngineCondition struct {
	Type   string `json:"type"`
	Status string `json:"status"`
}

type ResourcesProcessingModelInputs struct {
	Queued int `json:"queued"`
}

type ResourcesProcessingModelJobs struct {
	Queued int `json:"queued"`
}

type ResourcesProcessingModelDeploymentState struct {
	HasError       bool `json:"hasError"`
	Ready          bool `json:"ready"`
	BeingMonitored bool `json:"beingMonitored"`
}
