// nolint:errcheck
package modzy

import (
	"context"
	"testing"
)

func TestDashboardClientFake(t *testing.T) {
	expectedCtx := context.WithValue(context.TODO(), testContextKey("a"), "b")

	calls := 0
	fake := &DashboardClientFake{
		GetAlertsFunc: func(ctx context.Context, input *GetAlertsInput) (*GetAlertsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetAlertDetailsFunc: func(ctx context.Context, input *GetAlertDetailsInput) (*GetAlertDetailsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetDataProcessedFunc: func(ctx context.Context, input *GetDataProcessedInput) (*GetDataProcessedOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetPredictionsMadeFunc: func(ctx context.Context, input *GetPredictionsMadeInput) (*GetPredictionsMadeOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetActiveUsersFunc: func(ctx context.Context, input *GetActiveUsersInput) (*GetActiveUsersOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetActiveModelsFunc: func(ctx context.Context, input *GetActiveModelsInput) (*GetActiveModelsOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
		GetPrometheusMetricFunc: func(ctx context.Context, input *GetPrometheusMetricInput) (*GetPrometheusMetricOutput, error) {
			calls++
			if ctx != expectedCtx {
				t.Errorf("not expected ctx")
			}
			if input == nil {
				t.Errorf("input was not passed through")
			}
			return nil, nil
		},
	}

	fake.GetAlerts(expectedCtx, &GetAlertsInput{})
	fake.GetAlertDetails(expectedCtx, &GetAlertDetailsInput{})
	fake.GetDataProcessed(expectedCtx, &GetDataProcessedInput{})
	fake.GetPredictionsMade(expectedCtx, &GetPredictionsMadeInput{})
	fake.GetActiveUsers(expectedCtx, &GetActiveUsersInput{})
	fake.GetActiveModels(expectedCtx, &GetActiveModelsInput{})
	fake.GetPrometheusMetric(expectedCtx, &GetPrometheusMetricInput{})

	if calls != 7 {
		t.Errorf("Did not call all of the funcs: %d", calls)
	}
}
