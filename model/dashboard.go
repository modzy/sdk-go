package model

import "encoding/json"

type AlertsListType struct {
	Type  string `json:"type"`
	Count int    `json:"count"`
}

type AlertsList []AlertsListType

type DataProcessedSummary struct {
	RecentBytes   int64   `json:"recent"`
	PreviousBytes int64   `json:"previous"`
	Percentage    float64 `json:"percentage"`
}

type DataProcessingRecent struct {
	Date  ModzyDate `json:"date"`
	Bytes int64     `json:"value"`
}

type PredictionsMadeSummary struct {
	RecentPredictions   int64   `json:"recent"`
	PreviousPredictions int64   `json:"previous"`
	Percentage          float64 `json:"percentage"`
}

type PredictionsMadeRecent struct {
	Date  ModzyDate `json:"date"`
	Bytes int64     `json:"value"`
}

type ActiveUserSummary struct {
	Ranking     int    `json:"ranking"`
	Identifier  string `json:"modelIdentifier"`
	FirstName   string `json:"firstName"`
	LastName    string `json:"lastName"`
	Email       string `json:"email"`
	ElapsedTime int    `json:"elapsedTime"`
	Jobs        int    `json:"jobs"`
}

type ActiveModelSummary struct {
	Ranking     int    `json:"ranking"`
	Identifier  string `json:"modelIdentifier"`
	Name        string `json:"name"`
	Version     string `json:"modelVersion"`
	ElapsedTime int    `json:"elapsedTime"`
	Jobs        int    `json:"jobs"`
}

type PrometheusResponse struct {
	Data PrometheusData `json:"data"`
}

type PrometheusData struct {
	ResultType string             `json:"resultType"`
	Results    []PrometheusResult `json:"result"`
}

type PrometheusResult struct {
	Metric interface{}       `json:"metric"`
	Values []json.RawMessage `json:"values"`
}
