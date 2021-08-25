package modzy

import (
	"context"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/modzy/sdk-go/internal/impossible"
	"github.com/modzy/sdk-go/model"
)

type DashboardClient interface {
	GetAlerts(ctx context.Context, input *GetAlertsInput) (*GetAlertsOutput, error)
	GetAlertDetails(ctx context.Context, input *GetAlertDetailsInput) (*GetAlertDetailsOutput, error)
	GetDataProcessed(ctx context.Context, input *GetDataProcessedInput) (*GetDataProcessedOutput, error)
	GetPredictionsMade(ctx context.Context, input *GetPredictionsMadeInput) (*GetPredictionsMadeOutput, error)
	GetActiveUsers(ctx context.Context, input *GetActiveUsersInput) (*GetActiveUsersOutput, error)
	GetActiveModels(ctx context.Context, input *GetActiveModelsInput) (*GetActiveModelsOutput, error)
	GetPrometheusMetric(ctx context.Context, input *GetPrometheusMetricInput) (*GetPrometheusMetricOutput, error)
}

type standardDashboardClient struct {
	baseClient *standardClient
}

var _ DashboardClient = &standardDashboardClient{}

func (c *standardDashboardClient) GetAlerts(ctx context.Context, input *GetAlertsInput) (*GetAlertsOutput, error) {
	var out model.AlertsList
	path := "/api/notifications/alerts"
	_, err := c.baseClient.requestor.Get(ctx, path, &out)
	if err != nil {
		return nil, err
	}

	summaries := []AlertSummary{}
	for _, s := range out {
		summaries = append(summaries, AlertSummary{
			Type:  AlertType(s.Type),
			Count: s.Count,
		})
	}
	return &GetAlertsOutput{
		Alerts: summaries,
	}, nil
}

func (c *standardDashboardClient) GetAlertDetails(ctx context.Context, input *GetAlertDetailsInput) (*GetAlertDetailsOutput, error) {
	var out []string
	path := fmt.Sprintf("/api/notifications/alerts/%s", input.Type)
	_, err := c.baseClient.requestor.Get(ctx, path, &out)
	if err != nil {
		return nil, err
	}
	return &GetAlertDetailsOutput{
		Type:     input.Type,
		Entities: out,
	}, nil
}

func (c *standardDashboardClient) GetDataProcessed(ctx context.Context, input *GetDataProcessedInput) (*GetDataProcessedOutput, error) {
	url := c.parseDashboardFilters("/api/metrics/data-processed", dashboardFilters{
		BeginDate:       input.BeginDate,
		EndDate:         input.EndDate,
		UserIdentifier:  input.UserIdentifier,
		AccessKeyPrefix: input.AccessKeyPrefix,
		ModelIdentifier: input.ModelIdentifier,
		TeamIdentifier:  input.TeamIdentifier,
	})

	var out GetDataProcessedOutput
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *standardDashboardClient) GetPredictionsMade(ctx context.Context, input *GetPredictionsMadeInput) (*GetPredictionsMadeOutput, error) {
	url := c.parseDashboardFilters("/api/metrics/predictions-made", dashboardFilters{
		BeginDate:       input.BeginDate,
		EndDate:         input.EndDate,
		UserIdentifier:  input.UserIdentifier,
		AccessKeyPrefix: input.AccessKeyPrefix,
		ModelIdentifier: input.ModelIdentifier,
		TeamIdentifier:  input.TeamIdentifier,
	})

	var out GetPredictionsMadeOutput
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *standardDashboardClient) GetActiveUsers(ctx context.Context, input *GetActiveUsersInput) (*GetActiveUsersOutput, error) {
	url := c.parseDashboardFilters("/api/metrics/active-users", dashboardFilters{
		BeginDate:       input.BeginDate,
		EndDate:         input.EndDate,
		UserIdentifier:  input.UserIdentifier,
		AccessKeyPrefix: input.AccessKeyPrefix,
		ModelIdentifier: input.ModelIdentifier,
		TeamIdentifier:  input.TeamIdentifier,
	})

	var out []model.ActiveUserSummary
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetActiveUsersOutput{
		Users: out,
	}, nil
}

func (c *standardDashboardClient) GetActiveModels(ctx context.Context, input *GetActiveModelsInput) (*GetActiveModelsOutput, error) {
	url := c.parseDashboardFilters("/api/metrics/active-models", dashboardFilters{
		BeginDate:       input.BeginDate,
		EndDate:         input.EndDate,
		UserIdentifier:  input.UserIdentifier,
		AccessKeyPrefix: input.AccessKeyPrefix,
		ModelIdentifier: input.ModelIdentifier,
		TeamIdentifier:  input.TeamIdentifier,
	})

	var out []model.ActiveModelSummary
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetActiveModelsOutput{
		Models: out,
	}, nil
}

func (c *standardDashboardClient) GetPrometheusMetric(ctx context.Context, input *GetPrometheusMetricInput) (*GetPrometheusMetricOutput, error) {
	url := c.parseDashboardFilters(
		fmt.Sprintf("/api/metrics/prometheus/%s", input.Metric),
		dashboardFilters{
			BeginDate: input.BeginDate,
			EndDate:   input.EndDate,
		},
	)

	var out model.PrometheusResponse
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	var parsedValues []PrometheusValue

	for _, result := range out.Data.Results {
		for _, value := range result.Values {
			msg := fmt.Sprintf("%s", value)
			msg = strings.TrimPrefix(msg, "[")
			msg = strings.TrimSuffix(msg, "]")
			msgParts := strings.Split(msg, ",")

			if len(msgParts) < 2 {
				continue
			}

			parsedIntTime, err := strconv.ParseInt(strings.TrimSpace(msgParts[0]), 10, 64)
			if err != nil {
				continue
			}

			parsedValue := strings.TrimSuffix(strings.TrimPrefix(strings.TrimSpace(msgParts[1]), `"`), `"`)
			parsedValues = append(parsedValues, PrometheusValue{
				Time:  time.Unix(parsedIntTime, 0),
				Value: fmt.Sprintf("%v", parsedValue),
			})
		}
	}

	return &GetPrometheusMetricOutput{
		Values: parsedValues,
	}, nil
}

func (c *standardDashboardClient) parseDashboardFilters(path string, filters dashboardFilters) string {
	partialUrl, err := url.Parse(path)
	impossible.HandleError(err)

	q := partialUrl.Query()
	if !filters.BeginDate.IsZero() {
		q.Add("begin-date", filters.BeginDate.String())
	}
	if !filters.EndDate.IsZero() {
		q.Add("end-date", filters.EndDate.String())
	}
	if filters.UserIdentifier != "" {
		q.Add("user-identifier", filters.UserIdentifier)
	}
	if filters.AccessKeyPrefix != "" {
		q.Add("access-key", filters.AccessKeyPrefix)
	}
	if filters.ModelIdentifier != "" {
		q.Add("model-identifier", filters.ModelIdentifier)
	}
	if filters.TeamIdentifier != "" {
		q.Add("team-identifier", filters.TeamIdentifier)
	}
	partialUrl.RawQuery = q.Encode()
	return partialUrl.String()
}

type dashboardFilters struct {
	BeginDate       model.ModzyDate
	EndDate         model.ModzyDate
	UserIdentifier  string
	AccessKeyPrefix string
	ModelIdentifier string
	TeamIdentifier  string
}
