// This package provides code examples that utilize the Modzy sdk.
package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"time"

	modzy "github.com/modzy/go-sdk"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

var (
	ctx = context.TODO()
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	apiKey := os.Getenv("MODZY_API_KEY")
	baseURL := os.Getenv("MODZY_BASE_URL")
	client := modzy.NewClient(baseURL).WithAPIKey(apiKey)

	if os.Getenv("MODZY_DEBUG") == "1" {
		client = client.WithOptions(modzy.WithHTTPDebugging(false, false))
	}

	// listJobsHistory(client)
	// errorChecking()
	// submitExampleText(client, false)
	// submitExampleText(client, false)
	// describeJob(client, "86b76e20-c506-485d-af4e-2072c41ca35b")
	// describeModel(client, "ed542963de")
	// getRelatedModels(client, "ed542963de")
	// getMinimumEngines(client)
	// getJobFeatures(client)
	// listModels(client)
	// getTags(client)
	// getTagModels(client, []string{"time_series", "equipment_and_machinery"})
	// describeModelByName(client, "Sentiment Analysis")
	// listModelVersions(client, "ed542963de")
	// updateModelProcessingEngines(client, "ed542963de", "0.0.27")
	getModelSampleInputAndOutput(client, "ed542963de", "0.0.27")
}

func listJobsHistory(client modzy.Client) {
	logrus.Info("Will list job histories")

	// This will read the list of job histories, and continue paging until complete
	listJobsHistoryInput := (&modzy.ListJobsHistoryInput{}).
		WithPaging(2, 1).
		WithFilterOr(modzy.ListJobsHistoryFilterFieldStatus, modzy.JobStatusTimedOut). // , modzy.JobStatusPending
		WithSort(modzy.SortDirectionDescending, modzy.ListJobsHistorySortCreatedAt)

	for listJobsHistoryInput != nil {
		listJobsHistoryOut, err := client.Jobs().ListJobsHistory(ctx, listJobsHistoryInput)
		if err != nil {
			logrus.WithError(err).Fatalf("Failed to read job histories")
			return
		}

		logrus.Infof("Found %d job histories", len(listJobsHistoryOut.Jobs))

		// read the next page
		listJobsHistoryInput = listJobsHistoryOut.NextPage
	}
}

func errorChecking() {
	logrus.Info("Will make a call with an unauthenticated client")

	// no api key is provided
	client := modzy.NewClient("")
	_, err := client.Jobs().ListJobsHistory(ctx, &modzy.ListJobsHistoryInput{})
	if err != nil {
		if modzyErr, ok := err.(*modzy.ModzyHTTPError); ok {
			logrus.WithError(err).Warnf("This error is a modzy http error with additional information such as statusCode = %d", modzyErr.StatusCode)
		}

		if errors.Cause(err) == modzy.ErrUnauthorized {
			logrus.WithError(err).Warnf("No authentication mechanism was provided")
		} else {
			logrus.WithError(err).Warnf("An unexpected error occured")
		}

	} else {
		logrus.Error("errorChecking was expected to fail with an unauthenticated error")
	}
}

func submitExampleText(client modzy.Client, cancel bool) {
	logrus.Info("Will submit example text job")
	submittedJob, err := client.Jobs().SubmitJobText(ctx, &modzy.SubmitJobTextInput{
		ModelIdentifier: "ed542963de",
		ModelVersion:    "0.0.27",
		Timeout:         time.Second * 30,
		Inputs: map[string]modzy.TextInputItem{
			"happy-text": {
				"input.txt": "I love AI! :)",
			},
			"angry-text": {
				"input.txt": "I hate AI! abysmal. adverse. alarming. angry. annoy. anxious :(",
			},
			"mixed": {
				"input.txt": "love hate irrational brute",
			},
			"emojis": {
				"input.txt": ":-) :slightly_smiling_face: :disappointed: %%%%%8*",
			},
		},
	})
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to submit text job")
		return
	}

	logrus.WithField("jobIdentifier", submittedJob.Response.JobIdentifier).Info("Text job submitted")

	if cancel {
		logrus.Info("Will cancel job")
		cancelOut, err := submittedJob.Cancel(ctx)
		if err != nil {
			logrus.WithError(err).Fatalf("Failed to cancel job")
		}
		logrus.Infof("Job canceled: %s, %d", cancelOut.Details.Status, cancelOut.Details.HoursDeleteInput)
		return
	} else {
		logrus.Info("Will wait until job completes")
		jobDetails, err := submittedJob.WaitForCompletion(ctx, time.Second*5)
		if err != nil {
			logrus.WithError(err).Fatalf("Failed to wait for job completion")
			return
		}
		logrus.Infof("Job completed: %s -> %s", jobDetails.Details.JobIdentifier, jobDetails.Details.Status)
		jobResults, err := submittedJob.GetResults(ctx)
		if err != nil {
			logrus.WithError(err).Fatalf("Failed to get job results")
			return
		}
		logrus.Infof("Job results: %s -> %d results", jobResults.Results.JobIdentifier, jobResults.Results.Total)
	}
}

func describeJob(client modzy.Client, jobIdentifier string) {
	actions := client.Jobs().NewJobActions(jobIdentifier)
	jobResults, err := actions.GetResults(ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to get resutls for job")
	} else {
		logrus.Info("Dumping job results")
		enc := json.NewEncoder(logrus.StandardLogger().Out)
		enc.SetIndent("", "    ")
		_ = enc.Encode(jobResults)
	}

	modelDetails, err := actions.GetModelDetails(ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to get model details for job %s", jobIdentifier)
	} else {
		logrus.Info("Dumping model details")
		enc := json.NewEncoder(logrus.StandardLogger().Out)
		enc.SetIndent("", "    ")
		_ = enc.Encode(modelDetails)
	}
}

func getJobFeatures(client modzy.Client) {
	out, err := client.Jobs().GetJobFeatures(ctx, &modzy.GetJobFeaturesInput{})
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to list features")
		return
	}
	logrus.Infof("Features: %+v", out.Features)
}

func getMinimumEngines(client modzy.Client) {
	out, err := client.Models().GetMinimumEngines(ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to get minimum engines")
	}
	logrus.Infof("Minimum engines: %d", out.Details.MinimumProcessingEnginesSum)
}

func describeModel(client modzy.Client, modelID string) {
	out, err := client.Models().GetModelDetails(ctx, &modzy.GetModelDetailsInput{ModelID: modelID})
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to get model details for %s", modelID)
	} else {
		logrus.Info("Dumping model details")
		enc := json.NewEncoder(logrus.StandardLogger().Out)
		enc.SetIndent("", "    ")
		_ = enc.Encode(out)
	}
}

func getRelatedModels(client modzy.Client, modelID string) {
	out, err := client.Models().GetRelatedModels(ctx, &modzy.GetRelatedModelsInput{ModelID: modelID})
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to get related models")
	} else {
		logrus.Infof("Found %d related models", len(out.RelatedModels))
	}
}

func listModels(client modzy.Client) {
	out, err := client.Models().ListModels(ctx, (&modzy.ListModelsInput{}).
		WithFilterAnd(modzy.ListModelsFilterFieldAuthor, "modzy").
		WithFilterAnd(modzy.ListModelsFilterFieldIsActive, "false"),
	)
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to list imodels")
		return
	}
	logrus.Infof("Models: %+v", out.Models)
}

func getTags(client modzy.Client) {
	out, err := client.Models().GetTags(ctx)
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to get tags")
	} else {
		logrus.Infof("Found %d tags", len(out.Tags))
	}
}

func getTagModels(client modzy.Client, tagIDs []string) {
	out, err := client.Models().GetTagModels(ctx, &modzy.GetTagModelsInput{TagIDs: tagIDs})
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to get tag models")
	} else {
		logrus.Infof("Found %d tags and %d matching models", len(out.Tags), len(out.Models))
	}
}

func describeModelByName(client modzy.Client, name string) {
	out, err := client.Models().GetModelDetailsByName(ctx, &modzy.GetModelDetailsByNameInput{
		Name: name,
	})
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to get model details by name %s", name)
	} else {
		logrus.Info("Dumping model details")
		enc := json.NewEncoder(logrus.StandardLogger().Out)
		enc.SetIndent("", "    ")
		_ = enc.Encode(out)
	}
}

func listModelVersions(client modzy.Client, modelID string) {
	out, err := client.Models().ListModelVersions(ctx, (&modzy.ListModelVersionsInput{
		ModelID: modelID,
	}))
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to list model verions")
		return
	}
	logrus.Infof("Found %d versions for model %s", len(out.Versions), modelID)
}

func getModelSampleInputAndOutput(client modzy.Client, modelID string, version string) {
	in, err := client.Models().GetModelVersionSampleInput(ctx, &modzy.GetModelVersionSampleInputInput{
		ModelID: modelID,
		Version: version,
	})
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to get sample input")
	} else {
		logrus.Info("Dumping sample input:")
		fmt.Fprintln(logrus.StandardLogger().Out, in.Sample)
	}

	out, err := client.Models().GetModelVersionSampleOutput(ctx, &modzy.GetModelVersionSampleOutputInput{
		ModelID: modelID,
		Version: version,
	})
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to get sample output")
	} else {
		logrus.Info("Dumping sample outptu:")
		fmt.Fprintln(logrus.StandardLogger().Out, out.Sample)
	}
}

func updateModelProcessingEngines(client modzy.Client, modelID string, version string) {
	out, err := client.Models().GetModelVersionDetails(ctx, &modzy.GetModelVersionDetailsInput{
		ModelID: modelID,
		Version: version,
	})
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to get model version")
		return
	}

	// let our max change for testing, but don't climb forever
	newMax := out.Details.Processing.MaximumParallelCapacity + 1
	if newMax > 2 {
		newMax = 1
	}

	newOut, err := client.Models().UpdateModelProcessingEngines(ctx, &modzy.UpdateModelProcessingEnginesInput{
		ModelID:                 modelID,
		Version:                 version,
		MinimumParallelCapacity: out.Details.Processing.MinimumParallelCapacity,
		MaximumParallelCapacity: newMax,
	})
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to patch processing engines")
	} else {
		logrus.Infof("Patched processing engines to be: %+v", newOut.Details.Processing)
	}
}
