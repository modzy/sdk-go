package Modzy

import (
	"github.com/dghubble/sling"
	"log"
	"modzy.com/modzy-sdk/Modzy/types"
)

type JobService struct {
	sling *sling.Sling
}

//TODO: Libraries do not log statements out. Need to refactor to return errors for handling by user

func newJobService(sling *sling.Sling) *JobService {
	return &JobService{
		sling: sling.New().Path("/api/jobs/"),
	}
}

func (js *JobService) GetAllJobs(params *types.AllJobsParams) []types.AllJobSummary {
	var jobs []types.AllJobSummary
	_, e := js.sling.QueryStruct(params).Receive(&jobs, nil)
	if e != nil {
		log.Printf("Error retrieving jobs: %v", e)
	}
	return jobs
}

func (js *JobService) GetJobDetails(jobId string) *types.JobDetails {
	var jobDetails types.JobDetails
	_, e := js.sling.Path(jobId).Receive(&jobDetails, nil)
	if e != nil {
		log.Printf("Error retrieving job details: %v", e)
	}
	return &jobDetails
}
