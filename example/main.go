// This package provides code examples that utilize the Modzy sdk.
package main

import (
	"context"
	"encoding/json"
	"os"
	"time"

	modzy "github.com/modzy/go-sdk"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetLevel(logrus.DebugLevel)

	apiKey := os.Getenv("MODZY_API_KEY")
	baseURL := os.Getenv("MODZY_BASE_URL")
	client := modzy.NewClient(baseURL).WithAPIKey(apiKey)

	if os.Getenv("MODZY_DEBUG") == "1" {
		client = client.WithOptions(modzy.WithHTTPDebugging(false, false))
	}

	// listJobs(client, false)
	// listJobsHistory(client)
	// errorChecking()
	// submitExampleText(client, false)
	// submitExampleText(client, false)
	// describeJob(client, "86b76e20-c506-485d-af4e-2072c41ca35b")
	GetJobFeatures(client)
}

// func listJobs(client modzy.Client, outputDetails bool) {
// 	ctx := context.TODO()

// 	logrus.Info("Will list jobs")

// 	// This will read the list of jobs, and continue paging until complete
// 	listJobsInput := (&modzy.ListJobsInput{}).
// 		WithPaging(2, 1)

// 	for listJobsInput != nil {
// 		listJobsOut, err := client.Jobs().ListJobs(ctx, listJobsInput)
// 		if err != nil {
// 			logrus.WithError(err).Fatalf("Failed to read jobs")
// 			return
// 		}

// 		logrus.Infof("Found %d jobs", len(listJobsOut.Jobs))

// 		if outputDetails {
// 			for _, job := range listJobsOut.Jobs {
// 				logrus.Infof("- Job: [%s] %s", job.Status, job.JobIdentifier)

// 				jobDetails, err := client.Jobs().GetJobDetails(ctx, &modzy.GetJobDetailsInput{
// 					JobIdentifier: job.JobIdentifier,
// 				})
// 				if err != nil {
// 					logrus.WithError(err).Fatalf("Failed to read job details: %s", job.JobIdentifier)
// 					return
// 				}
// 				logrus.Infof("  - Model Name: %s", jobDetails.Details.Model.Name)
// 				logrus.Infof("  - Completed: %d", jobDetails.Details.Completed)
// 			}
// 		}

// 		// read the next page
// 		listJobsInput = listJobsOut.NextPage
// 	}
// }

func listJobsHistory(client modzy.Client) {
	ctx := context.TODO()

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
	ctx := context.TODO()

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
	ctx := context.TODO()

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
	ctx := context.TODO()

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

func GetJobFeatures(client modzy.Client) {
	ctx := context.TODO()
	out, err := client.Jobs().GetJobFeatures(ctx, &modzy.GetJobFeaturesInput{})
	if err != nil {
		logrus.WithError(err).Fatalf("Failed to list features")
		return
	}
	logrus.Infof("Features: %+v", out.Features)
}
