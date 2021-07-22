package model

import (
	"encoding/json"
)

type ModelVersion struct {
	Version string `json:"version"`
}

type ModelVersionSummary struct {
	ID            string   `json:"modelId"`
	LatestVersion string   `json:"latestVersion"`
	Versions      []string `json:"versions"`
}

type ModelDetailsTimeout struct {
	Status int `json:"status"`
	Run    int `json:"run"`
}

type ModelContainerImage struct {
	UploadStatus       string `json:"uploadStatus"`
	LoadStatus         string `json:"loadStatus"`
	UploadPercentage   int    `json:"uploadPercentage"`
	LoadPercentage     int    `json:"loadPercentage"`
	ContainerImageSize int    `json:"containerImageSize"`
	RegistryHost       string `json:"registryHost"`
}

type ModelLoadStatus struct {
	Step       int    `json:"step"`
	StepName   string `json:"stepName"`
	Percentage int    `json:"percentage"`
}

type ModelRunStatusResult struct {
	Status      string    `json:"status"`
	Enging      string    `json:"engine"`
	StartTime   ModzyTime `json:"startTime"`
	EndTime     ModzyTime `json:"endTime"`
	UpdateDate  ModzyTime `json:"updateTime"`
	ElapsedTime int       `json:"elapsedTime"`
}

type ModelRunStatus struct {
	Step       int                  `json:"step"`
	StepName   string               `json:"stepName"`
	Percentage int                  `json:"percentage"`
	Result     ModelRunStatusResult `json:"result"`
}

type ModelTag struct {
	Identifier    string          `json:"identifier"`
	Name          string          `json:"name"`
	DataType      string          `json:"dataType"`
	IsCategorical bool            `json:"isCategorical"`
	Images        json.RawMessage `json:"images"`
}

type ModelImage struct {
	URL          string `json:"url"`
	Caption      string `json:"caption"`
	Alt          string `json:"alt"`
	RelationType string `json:"relationType"`
}

type ModelVisibility struct {
	Scope string   `json:"scope"`
	Teams []string `json:"teams"`
}

type ModelMetdata struct {
	ModelID            string          `json:"modelId"`
	LatestVersion      string          `json:"latestVersion"`
	Versions           []string        `json:"versions"`
	Author             string          `json:"author"`
	Name               string          `json:"name"`
	Description        string          `json:"description"`
	Permalink          string          `json:"permalink"`
	Features           json.RawMessage `json:"features"`
	IsActive           bool            `json:"isActive"`
	IsRecommended      bool            `json:"isRecommended"`
	IsCommercial       bool            `json:"isCommercial"`
	Tags               []ModelTag      `json:"tags"`
	Images             []ModelImage    `json:"images"`
	LastActiveDateTime ModzyTime       `json:"lastActiveDateTime"`
	CreatedByEmail     string          `json:"createdByEmail"`
	Visibility         ModelVisibility `json:"visibility"`
}

type ModelVersionDetails struct {
	ModelID               string              `json:"-"`
	Version               string              `json:"version"`
	CreatedAt             ModzyTime           `json:"createdAt"`
	UpdatedAt             ModzyTime           `json:"updatedAt"`
	InputValidationSchema json.RawMessage     `json:"inputValidationSchema"`
	Timeout               ModelDetailsTimeout `json:"timeout"`
	Requirement           json.RawMessage     `json:"requirement"`
	ContainerImage        ModelContainerImage `json:"containerImage"`
	LoadStatus            ModelLoadStatus     `json:"loadStatus"`
	RunStatus             ModelRunStatus      `json:"runStatus"`
	Inputs                []json.RawMessage   `json:"inputs"`
	Outputs               []json.RawMessage   `json:"outputs"`
	Statistics            []json.RawMessage   `json:"statistics"`
	IsActive              bool                `json:"isActive"`
	LongDescription       string              `json:"longDescription"`
	TechnicalDetails      string              `json:"technicalDetails"`
	IsAvailable           bool                `json:"isAvailable"`
	SourceType            string              `json:"sourceType"`
	VersionHistory        string              `json:"versionHistory"`
	Status                string              `json:"status"`
	PerformanceSummary    string              `json:"performanceSummary"`
	Model                 ModelMetdata        `json:"model"`
}

type ModelDetails struct {
	ModelID             string          `json:"modelID"`
	LatestVersion       string          `json:"latestVersion"`
	LatestActiveVersion string          `json:"latestActiveVersion"`
	Versions            []string        `json:"versions"`
	Author              string          `json:"author"`
	Name                string          `json:"name"`
	Description         string          `json:"description"`
	Permalink           string          `json:"permalink"`
	Features            json.RawMessage `json:"features"`
	IsActive            bool            `json:"isActive"`
	IsRecommended       bool            `json:"isRecommended"`
	IsCommercial        bool            `json:"isCommercial"`
	Tags                []ModelTag      `json:"tags"`
	Images              []ModelImage    `json:"images"`
	SnapshotImages      json.RawMessage `json:"snapshotImages"`
	LastActiveDateTime  ModzyTime       `json:"lastActiveDateTime"`
	Visibility          ModelVisibility `json:"visibility"`
}

type RelatedModel struct {
	ModelID       string          `json:"identifier"`
	LatestVersion string          `json:"latestVersion"`
	Versions      []string        `json:"versions"`
	Author        string          `json:"author"`
	Name          string          `json:"name"`
	Description   string          `json:"description"`
	Permalink     string          `json:"permalink"`
	Features      json.RawMessage `json:"features"`
	IsActive      bool            `json:"isActive"`
	IsRecommended bool            `json:"isRecommended"`
	Tags          []ModelTag      `json:"tags"`
	Images        []ModelImage    `json:"images"`
}

type MinimumEngines struct {
	MinimumProcessingEnginesSum int `json:"minimumProcessingEnginesSum"`
}

type ModelWithTags struct {
	Identifier string     `json:"identifier"`
	Name       string     `json:"name"`
	Tags       []ModelTag `json;"tags"`
}
