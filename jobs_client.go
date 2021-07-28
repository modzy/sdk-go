package modzy

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"time"

	"github.com/docker/go-units"
	"github.com/modzy/go-sdk/model"
	"github.com/pkg/errors"
)

type JobsClient interface {
	GetJobDetails(ctx context.Context, input *GetJobDetailsInput) (*GetJobDetailsOutput, error)
	ListJobsHistory(ctx context.Context, input *ListJobsHistoryInput) (*ListJobsHistoryOutput, error)
	SubmitJob(ctx context.Context, input *SubmitJobInput) (*SubmitJobOutput, error)
	SubmitJobText(ctx context.Context, input *SubmitJobTextInput) (*SubmitJobTextOutput, error)
	SubmitJobEmbedded(ctx context.Context, input *SubmitJobEmbeddedInput) (*SubmitJobEmbeddedOutput, error)
	SubmitJobFile(ctx context.Context, input *SubmitJobFileInput) (*SubmitJobFileOutput, error)
	// SubmitJobS3(ctx context.Context, input *SubmitJobS3Input) (*SubmitJobS3Output, error)
	// SubmitJobJDBC(ctx context.Context, input *SubmitJobJDBCInput) (*SubmitJobJDBCOutput, error)
	WaitForJobCompletion(ctx context.Context, input *WaitForJobCompletionInput, pollInterval time.Duration) (*GetJobDetailsOutput, error)
	CancelJob(ctx context.Context, input *CancelJobInput) (*CancelJobOutput, error)
	GetJobResults(ctx context.Context, input *GetJobResultsInput) (*GetJobResultsOutput, error)
	GetJobFeatures(ctx context.Context) (*GetJobFeaturesOutput, error)
}

type standardJobsClient struct {
	baseClient *standardClient
}

var _ JobsClient = &standardJobsClient{}

func (c *standardJobsClient) GetJobDetails(ctx context.Context, input *GetJobDetailsInput) (*GetJobDetailsOutput, error) {
	var out model.JobDetails
	url := fmt.Sprintf("/api/jobs/%s", input.JobIdentifier)
	_, err := c.baseClient.requestor.Get(ctx, url, &out)
	if err != nil {
		return nil, err
	}

	return &GetJobDetailsOutput{
		Details: out,
	}, nil
}

func (c *standardJobsClient) ListJobsHistory(ctx context.Context, input *ListJobsHistoryInput) (*ListJobsHistoryOutput, error) {
	input.Paging = input.Paging.withDefaults()

	var items []model.JobSummary
	url := "/api/jobs/history"
	_, links, err := c.baseClient.requestor.List(ctx, url, input.Paging, &items)
	if err != nil {
		return nil, err
	}

	// decide if we have a next page (the next link is not always accurate?)
	var nextPage *ListJobsHistoryInput
	if _, hasNextLink := links["next"]; len(items) == input.Paging.PerPage && hasNextLink {
		nextPage = &ListJobsHistoryInput{
			Paging: input.Paging.Next(),
		}
	}

	return &ListJobsHistoryOutput{
		Jobs:     items,
		NextPage: nextPage,
	}, nil
}

func (c *standardJobsClient) SubmitJob(ctx context.Context, input *SubmitJobInput) (*SubmitJobOutput, error) {
	textInputs := map[string]TextInputItem{}
	embeddedInputs := map[string]EmbeddedInputItem{}
	chunkedInputs := map[string]FileInputItem{}

	for k, v := range input.Inputs {
		for innerK, innerV := range v {
			jobInputDescription, err := innerV()
			if err != nil {
				return nil, errors.WithMessagef(err, "failed to get data reader for item %s/%s", k, innerK)
			}

			switch jobInputDescription.Type {
			case jobInputableTypeString:
				if _, has := textInputs[k]; !has {
					textInputs[k] = TextInputItem{}
				}
				if fullStringBytes, err := ioutil.ReadAll(jobInputDescription.Data); err != nil {
					return nil, errors.WithMessagef(err, "failed to get data for item %s/%s", k, innerK)
				} else {
					textInputs[k][innerK] = string(fullStringBytes)
				}

			case jobInputableTypeEmbedded:
				if _, has := embeddedInputs[k]; !has {
					embeddedInputs[k] = EmbeddedInputItem{}
				}
				embeddedInputs[k][innerK] = URIEncodedReader(jobInputDescription.Data)

			case jobInputableTypeByte:
				if _, has := chunkedInputs[k]; !has {
					chunkedInputs[k] = FileInputItem{}
				}
				chunkedInputs[k][innerK] = ChunkReader(jobInputDescription.Data)
			}
		}
	}

	if len(textInputs) > 0 {
		return c.SubmitJobText(ctx, &SubmitJobTextInput{
			ModelIdentifier: input.ModelIdentifier,
			ModelVersion:    input.ModelVersion,
			Explain:         input.Explain,
			Timeout:         input.Timeout,
			Inputs:          textInputs,
		})
	}

	if len(embeddedInputs) > 0 {
		return c.SubmitJobEmbedded(ctx, &SubmitJobEmbeddedInput{
			ModelIdentifier: input.ModelIdentifier,
			ModelVersion:    input.ModelVersion,
			Explain:         input.Explain,
			Timeout:         input.Timeout,
			Inputs:          embeddedInputs,
		})
	}

	if len(chunkedInputs) > 0 {
		return c.SubmitJobFile(ctx, &SubmitJobFileInput{
			ModelIdentifier: input.ModelIdentifier,
			ModelVersion:    input.ModelVersion,
			Explain:         input.Explain,
			Timeout:         input.Timeout,
			Inputs:          chunkedInputs,
		})
	}

	return nil, fmt.Errorf("No inputs were provided")
}

func (c *standardJobsClient) SubmitJobText(ctx context.Context, input *SubmitJobTextInput) (*SubmitJobTextOutput, error) {

	toPostSources := map[string]model.TextInputItem{}
	for k, v := range input.Inputs {
		input := map[string]string{}
		for innerK, innerV := range v {
			input[innerK] = innerV
		}
		toPostSources[k] = input
	}

	toPost := model.SubmitTextJob{
		Model: model.SubmitJobModelInfo{
			Identifier: input.ModelIdentifier,
			Version:    input.ModelVersion,
		},
		Explain: input.Explain,
		Timeout: int(input.Timeout / time.Millisecond),
		Input: model.TextInput{
			Type:    "text",
			Sources: toPostSources,
		},
	}

	var response model.SubmitJobResponse

	url := "/api/jobs"
	_, err := c.baseClient.requestor.Post(ctx, url, toPost, &response)
	if err != nil {
		return nil, err
	}

	return &SubmitJobTextOutput{
		Response:   response,
		JobActions: NewJobActions(c.baseClient, response.JobIdentifier),
	}, nil
}

func (c *standardJobsClient) SubmitJobEmbedded(ctx context.Context, input *SubmitJobEmbeddedInput) (*SubmitJobEmbeddedOutput, error) {
	toPostSources := map[string]model.EmbeddedInputItem{}
	for k, v := range input.Inputs {
		input := map[string]string{}
		for innerK, innerV := range v {
			dataReader, err := innerV()
			if err != nil {
				return nil, errors.WithMessagef(err, "Failed to get data reader for item %s/%s", k, innerK)
			}
			encodedString, err := io.ReadAll(dataReader)
			if err != nil {

				return nil, errors.WithMessagef(err, "Failed to stream data for item %s/%s", k, innerK)
			}
			input[innerK] = string(encodedString)
		}
		toPostSources[k] = input
	}

	toPost := model.SubmitEmbeddedJob{
		Model: model.SubmitJobModelInfo{
			Identifier: input.ModelIdentifier,
			Version:    input.ModelVersion,
		},
		Explain: input.Explain,
		Timeout: int(input.Timeout / time.Millisecond),
		Input: model.EmbeddedInput{
			Type:    "embedded",
			Sources: toPostSources,
		},
	}

	var response model.SubmitJobResponse

	url := "/api/jobs"
	_, err := c.baseClient.requestor.Post(ctx, url, toPost, &response)
	if err != nil {
		return nil, err
	}

	return &SubmitJobEmbeddedOutput{
		Response:   response,
		JobActions: NewJobActions(c.baseClient, response.JobIdentifier),
	}, nil
}

func (c *standardJobsClient) SubmitJobFile(ctx context.Context, input *SubmitJobFileInput) (*SubmitJobFileOutput, error) {
	chunkSize, err := c.getMaxChunkSize(ctx, input.ChunkSize)
	if err != nil {
		return nil, errors.WithMessage(err, "Failed to get max chunk size")
	}

	noInputJob := model.SubmitChunkedJob{
		Model: model.SubmitJobModelInfo{
			Identifier: input.ModelIdentifier,
			Version:    input.ModelVersion,
		},
		Explain: input.Explain,
		Timeout: int(input.Timeout / time.Millisecond),
	}
	var response model.SubmitJobResponse
	if _, err := c.baseClient.requestor.Post(ctx, "/api/jobs", noInputJob, &response); err != nil {
		return nil, errors.WithMessage(err, "failed to post open job before posting input chunks")
	}
	jobActions := NewJobActions(c.baseClient, response.JobIdentifier)

	chunkErr := c.postInputsAsChunks(ctx, response.JobIdentifier, chunkSize, input.Inputs)
	if chunkErr != nil {
		// uploading the inputs failed, close the job
		_, _ = jobActions.Cancel(ctx)
		return nil, errors.WithMessage(err, "job canceled due to failure to upload data")
	} else {
		// close the job since everything is posted
		closeURL := fmt.Sprintf("/api/jobs/%s/close", response.JobIdentifier)
		if _, err := c.baseClient.requestor.Post(ctx, closeURL, noInputJob, &response); err != nil {
			return nil, errors.WithMessage(err, "failed to close open job after sucessfully uploading inputs")
		}
	}

	return &SubmitJobFileOutput{
		Response:   response,
		JobActions: jobActions,
	}, nil
}

func (c *standardJobsClient) getMaxChunkSize(ctx context.Context, defaultChunkSize int) (int64, error) {
	features, err := c.GetJobFeatures(ctx)
	if err != nil {
		return 0, err
	}
	maxChunkSize, err := units.FromHumanSize(features.Features.InputChunkMaximumSize)
	if err != nil {
		return 0, errors.WithMessage(err, "failed to parse InputChunkMaximumSize as an integer")
	}
	if maxChunkSize == 0 {
		maxChunkSize = 1024 * 1024
	}
	chunkSize := int64(defaultChunkSize)
	if chunkSize == 0 || chunkSize > maxChunkSize {
		chunkSize = maxChunkSize
	}
	return chunkSize, nil
}

func (c *standardJobsClient) postInputsAsChunks(ctx context.Context, jobID string, chunkSize int64, inputs map[string]FileInputItem) error {
	// go through each input and submit the data in chunks as necessary
	for k, v := range inputs {
		for innerK, innerV := range v {
			dataReader, err := innerV()
			if err != nil {
				return errors.WithMessagef(err, "failed to get data reader for item %s/%s", k, innerK)
			}

			// post as many chunks as necessary
			buf, err := ioutil.ReadAll(dataReader)
			if err != nil {
				return errors.WithMessage(err, "failed reading a chunk of data")
			}
			start := 0
			end := 0
			for {
				end = start + int(chunkSize)
				if end > len(buf) {
					end = len(buf)
				}
				if start == end {
					break
				}
				chunk := buf[start:end]
				chunkURL := fmt.Sprintf("/api/jobs/%s/%s/%s", jobID, k, innerK)
				chunkReader := bytes.NewReader(chunk)
				if _, err := c.baseClient.requestor.PostMultipart(ctx, chunkURL, map[string]io.Reader{"input": chunkReader}, nil); err != nil {
					return errors.WithMessage(err, "failed post a chunk of data")
				}
				start = end
			}
		}
	}
	return nil
}

// WaitForJobCompletion will wait until the provided job is done processing.
// The minimum pollInterval is 5 seconds.
// If the provided context is canceled, this wait will error.
func (c *standardJobsClient) WaitForJobCompletion(ctx context.Context, input *WaitForJobCompletionInput, pollInterval time.Duration) (*GetJobDetailsOutput, error) {
	timer := time.NewTimer(pollInterval)

	for {
		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("Wait for job completion was canceled due to provided context being canceled")
		case <-timer.C:
			job, err := c.GetJobDetails(ctx, &GetJobDetailsInput{input.JobIdentifier})
			if err != nil {
				return nil, err
			}
			// check
			if job.Details.Status == JobStatusCanceled ||
				job.Details.Status == JobStatusCompleted ||
				job.Details.Status == JobStatusTimedOut {
				// job is done
				return job, nil
			}
			if job.Details.Status == JobStatusOpen {
				return nil, fmt.Errorf("Job is currently OPEN and will never complete")
			}
			// not done -- wait and try again
			timer.Reset(pollInterval)
		}
	}
}

func (c *standardJobsClient) CancelJob(ctx context.Context, input *CancelJobInput) (*CancelJobOutput, error) {
	var response model.JobDetails

	url := fmt.Sprintf("/api/jobs/%s", input.JobIdentifier)
	_, err := c.baseClient.requestor.Delete(ctx, url, &response)
	if err != nil {
		return nil, err
	}

	return &CancelJobOutput{
		Details: response,
	}, nil
}

func (c *standardJobsClient) GetJobResults(ctx context.Context, input *GetJobResultsInput) (*GetJobResultsOutput, error) {
	var response model.JobResults

	url := fmt.Sprintf("/api/results/%s", input.JobIdentifier)
	_, err := c.baseClient.requestor.Get(ctx, url, &response)
	if err != nil {
		return nil, err
	}

	return &GetJobResultsOutput{
		Results: response,
	}, nil
}

func (c *standardJobsClient) GetJobFeatures(ctx context.Context) (*GetJobFeaturesOutput, error) {
	var response model.JobFeatures

	url := fmt.Sprintf("/api/jobs/features")
	_, err := c.baseClient.requestor.Get(ctx, url, &response)
	if err != nil {
		return nil, err
	}

	return &GetJobFeaturesOutput{
		Features: response,
	}, nil
}
