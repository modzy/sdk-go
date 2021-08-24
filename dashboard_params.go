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

// GetAlertsDetailsInput -
type GetAlertDetailsInput struct {
	Type AlertType `json:"type"`
}

// GetAlertDetailsOutput -
type GetAlertDetailsOutput struct {
	Type     AlertType `json:"type"`
	Entities []string  `json:"entities"`
}

// The default and minimum accepted time between BtartDate and EndDate is 7 days.
// If only one date is provided the API matches it with a 7 day range.
type GetDataProcessedInput struct {
	DashboardFilters
}

type GetDataProcessedOutput struct {
	Summary model.DataProcessedSummary   `json:"dataProcessed"`
	Recent  []model.DataProcessingRecent `json:"recent"`
}

// The default and minimum accepted time between BtartDate and EndDate is 7 days.
// If only one date is provided the API matches it with a 7 day range.
type GetPredictionsMadeInput struct {
	DashboardFilters
}

type GetPredictionsMadeOutput struct {
	Summary model.PredictionsMadeSummary  `json:"predicationsMade"`
	Recent  []model.PredictionsMadeRecent `json:"recent"`
}

type GetActiveUsersInput struct {
	DashboardFilters
}

type GetActiveUsersOutput struct {
	Users []model.ActiveUserSummary `json:"users"`
}

type GetActiveModelsInput struct {
	DashboardFilters
}

type GetActiveModelsOutput struct {
	Models []model.ActiveModelSummary `json:"models"`
}

type DashboardFilters struct {
	BeginEndFilters
	UserIdentifier  string
	AccessKeyPrefix string
	ModelIdentifier string
	TeamIdentifier  string
}

type PrometheusMetricType string

const (
	// The number of cores requested by a container
	PrometheusMetricTypeCPURequest PrometheusMetricType = "cpu-requested"
	// The cluster’s total number of available CPU cores.
	PrometheusMetricTypeCPUAvailable PrometheusMetricType = "cpu-available"
	// The total amount of “system” time + the total amount of “user” time
	PrometheusMetricTypeCPUUsed PrometheusMetricType = "cpu-used"
	// The number of memory bytes requested by a container
	PrometheusMetricTypeMemoryRequested PrometheusMetricType = "memory-requested"
	//  A node’s total allocatable memory bytes
	PrometheusMetricTypeMemoryAvailable PrometheusMetricType = "memory-available"
	// The current memory usage in bytes, it includes all memory regardless of when it was accessed
	PrometheusMetricTypeMemoryUsed PrometheusMetricType = "memory-used"
	// cpu-used / cpu-available
	PrometheusMetricTypeCPUOverallUsage PrometheusMetricType = "cpu-overall-usage"
	// memory-used / memory-available
	PrometheusMetricTypeMemoryOverallUsage PrometheusMetricType = "memory-overall-usage"
	// cpu-requested / cpu-available
	PrometheusMetricTypeCPUCurrentUsage PrometheusMetricType = "cpu-current-usage"
)

// The default and minimum accepted time between startDate and endDate is 7 days.
type GetPrometheusMetricInput struct {
	BeginEndFilters
	Metric    PrometheusMetricType
	BeginDate model.ModzyDate
	EndDate   model.ModzyDate
}

type BeginEndFilters struct {
	BeginDate model.ModzyDate
	EndDate   model.ModzyDate
}

type PrometheusValue struct {
	Time  time.Time `json:"time"`
	Value string    `json:"value"`
}

type GetPrometheusMetricOutput struct {
	Values []PrometheusValue `json:"values"`
}
