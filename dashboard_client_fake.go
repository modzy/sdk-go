package modzy

import (
	"context"
)

// DashboardClientFake is meant to help in mocking the DashboardClient interface easily for unit testing.
type DashboardClientFake struct {
	GetAlertsFunc           func(ctx context.Context, input *GetAlertsInput) (*GetAlertsOutput, error)
	GetAlertDetailsFunc     func(ctx context.Context, input *GetAlertDetailsInput) (*GetAlertDetailsOutput, error)
	GetDataProcessedFunc    func(ctx context.Context, input *GetDataProcessedInput) (*GetDataProcessedOutput, error)
	GetPredictionsMadeFunc  func(ctx context.Context, input *GetPredictionsMadeInput) (*GetPredictionsMadeOutput, error)
	GetActiveUsersFunc      func(ctx context.Context, input *GetActiveUsersInput) (*GetActiveUsersOutput, error)
	GetActiveModelsFunc     func(ctx context.Context, input *GetActiveModelsInput) (*GetActiveModelsOutput, error)
	GetPrometheusMetricFunc func(ctx context.Context, input *GetPrometheusMetricInput) (*GetPrometheusMetricOutput, error)
}

var _ DashboardClient = &DashboardClientFake{}

func (c *DashboardClientFake) GetAlerts(ctx context.Context, input *GetAlertsInput) (*GetAlertsOutput, error) {
	return c.GetAlertsFunc(ctx, input)
}

func (c *DashboardClientFake) GetAlertDetails(ctx context.Context, input *GetAlertDetailsInput) (*GetAlertDetailsOutput, error) {
	return c.GetAlertDetailsFunc(ctx, input)
}

func (c *DashboardClientFake) GetDataProcessed(ctx context.Context, input *GetDataProcessedInput) (*GetDataProcessedOutput, error) {
	return c.GetDataProcessedFunc(ctx, input)
}

func (c *DashboardClientFake) GetPredictionsMade(ctx context.Context, input *GetPredictionsMadeInput) (*GetPredictionsMadeOutput, error) {
	return c.GetPredictionsMadeFunc(ctx, input)
}

func (c *DashboardClientFake) GetActiveUsers(ctx context.Context, input *GetActiveUsersInput) (*GetActiveUsersOutput, error) {
	return c.GetActiveUsersFunc(ctx, input)
}

func (c *DashboardClientFake) GetActiveModels(ctx context.Context, input *GetActiveModelsInput) (*GetActiveModelsOutput, error) {
	return c.GetActiveModelsFunc(ctx, input)
}

func (c *DashboardClientFake) GetPrometheusMetric(ctx context.Context, input *GetPrometheusMetricInput) (*GetPrometheusMetricOutput, error) {
	return c.GetPrometheusMetricFunc(ctx, input)
}
