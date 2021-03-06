package modzy

import (
	"time"

	"github.com/modzy/sdk-go/model"
)

type AlertType string

const (
	AlertTypeJobStuck        AlertType = "JOB_STUCK"
	AlertTypeModelExpiration AlertType = "MODEL_EXPIRATION"
)

// GetAlertsInput -
type GetAlertsInput struct{}

type AlertSummary struct {
	Type  AlertType `json:"type"`
	Count int       `json:"count"`
}

// GetAlertsOutput -
type GetAlertsOutput struct {
	Alerts []AlertSummary `json:"alerts"`
}

// GetAlertDetailsInput -
type GetAlertDetailsInput struct {
	Type AlertType `json:"type"`
}

// GetAlertDetailsOutput -
type GetAlertDetailsOutput struct {
	Type     AlertType `json:"type"`
	Entities []string  `json:"entities"`
}

// GetDataProcessedInput -
// The default and minimum accepted time between BtartDate and EndDate is 7 days.
// If only one date is provided the API matches it with a 7 day range.
type GetDataProcessedInput struct {
	BeginDate       model.ModzyDate
	EndDate         model.ModzyDate
	UserIdentifier  string
	AccessKeyPrefix string
	ModelIdentifier string
	TeamIdentifier  string
}

type GetDataProcessedOutput struct {
	Summary model.DataProcessedSummary   `json:"dataProcessed"`
	Recent  []model.DataProcessingRecent `json:"recent"`
}

// GetPredictionsMadeInput -
// The default and minimum accepted time between BtartDate and EndDate is 7 days.
// If only one date is provided the API matches it with a 7 day range.
type GetPredictionsMadeInput struct {
	BeginDate       model.ModzyDate
	EndDate         model.ModzyDate
	UserIdentifier  string
	AccessKeyPrefix string
	ModelIdentifier string
	TeamIdentifier  string
}

type GetPredictionsMadeOutput struct {
	Summary model.PredictionsMadeSummary  `json:"predictionsMade"`
	Recent  []model.PredictionsMadeRecent `json:"recent"`
}

type GetActiveUsersInput struct {
	BeginDate       model.ModzyDate
	EndDate         model.ModzyDate
	UserIdentifier  string
	AccessKeyPrefix string
	ModelIdentifier string
	TeamIdentifier  string
}

type GetActiveUsersOutput struct {
	Users []model.ActiveUserSummary `json:"users"`
}

type GetActiveModelsInput struct {
	BeginDate       model.ModzyDate
	EndDate         model.ModzyDate
	UserIdentifier  string
	AccessKeyPrefix string
	ModelIdentifier string
	TeamIdentifier  string
}

type GetActiveModelsOutput struct {
	Models []model.ActiveModelSummary `json:"models"`
}

type PrometheusMetricType string

const (
	// PrometheusMetricTypeCPURequest - The number of cores requested by a container
	PrometheusMetricTypeCPURequest PrometheusMetricType = "cpu-requested"
	// PrometheusMetricTypeCPUAvailable - The cluster???s total number of available CPU cores.
	PrometheusMetricTypeCPUAvailable PrometheusMetricType = "cpu-available"
	// PrometheusMetricTypeCPUUsed - The total amount of ???system??? time + the total amount of ???user??? time
	PrometheusMetricTypeCPUUsed PrometheusMetricType = "cpu-used"
	// PrometheusMetricTypeMemoryRequested - The number of memory bytes requested by a container
	PrometheusMetricTypeMemoryRequested PrometheusMetricType = "memory-requested"
	// PrometheusMetricTypeMemoryAvailable - A node???s total allocatable memory bytes
	PrometheusMetricTypeMemoryAvailable PrometheusMetricType = "memory-available"
	// PrometheusMetricTypeMemoryUsed - The current memory usage in bytes, it includes all memory regardless of when it was accessed
	PrometheusMetricTypeMemoryUsed PrometheusMetricType = "memory-used"
	// PrometheusMetricTypeCPUOverallUsage - cpu-used / cpu-available
	PrometheusMetricTypeCPUOverallUsage PrometheusMetricType = "cpu-overall-usage"
	// PrometheusMetricTypeMemoryOverallUsage - memory-used / memory-available
	PrometheusMetricTypeMemoryOverallUsage PrometheusMetricType = "memory-overall-usage"
	// PrometheusMetricTypeCPUCurrentUsage - cpu-requested / cpu-available
	PrometheusMetricTypeCPUCurrentUsage PrometheusMetricType = "cpu-current-usage"
)

// GetPrometheusMetricInput - The default and minimum accepted time between startDate and endDate is 7 days.
type GetPrometheusMetricInput struct {
	BeginDate model.ModzyDate
	EndDate   model.ModzyDate
	Metric    PrometheusMetricType
}

type PrometheusValue struct {
	Time  time.Time `json:"time"`
	Value string    `json:"value"`
}

type GetPrometheusMetricOutput struct {
	Values []PrometheusValue `json:"values"`
}
