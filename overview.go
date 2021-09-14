// Package modzy provides an SDK to easily access the modzy http API.
//
// To use this SDK you need to create a modzy Client and then call any useful functions.  For example, if you need to get details about a model:
// 	client := modzy.NewClient("https://your-base-url.example.com").WithAPIKey("your-api-key")
// 	details, err := client.Models().GetModelDetails(ctx, &modzy.GetJobDetailsInput{
// 		ModelIdentifier: "e3f73163d3",
// 	})
// 	if err != nil {
// 		log.Fatalf("Failed to get model details: %v", err)
//	 }
// 	fmt.Printf("Found latest version: %s\n", details.Details.LatestVersion)
//
// The methods are organized into collections under the main client. For example:
//	client.Jobs().GetJobDetails(...)
//	client.Models().GetModelDetails(...)
//
// All SDK functions require a context to be provided, and this context is passed to the resulting http.Request.  You may choose to use this knowledge to implement custom tracing or other supportability requirements you may have.
// If cancel a context, the resulting http request, or any other internal processes will halt.
//
// For all cases where a known list of values exists, the SDK functions will use a specific type to help you find those easily.  For example, when filtering and/or sorting the job history, there are const values for the various inputs:
//
// 	listJobsHistoryInput := (&modzy.ListJobsHistoryInput{}).
//		WithFilterOr(modzy.ListJobsHistoryFilterFieldStatus, modzy.JobStatusTimedOut).
//		WithSort(modzy.SortDirectionDescending, modzy.ListJobsHistorySortFieldCreatedAt)
//
// In the case that the API is updated before the SDK can be updated, you can use the const types yourself:
//
//	WithFilterOr(modzy.ListJobsHistoryFilterField("new-field"), "a value")
//
// Fake types are provided for each of the main client interfaces to allow for easy mocking during testing:
//
// 	var clientMock = &modzy.ClientFake{
//		JobsFunc: func() modzy.JobsClient {
// 			return &modzy.JobsClientFake {
// 				GetJobDetailsFunc: func(ctx context.Context, input *modzy.GetJobDetailsInput) (*modzy.GetJobDetailsOutput, error) {
// 					return nil, fmt.Errorf("No details!")
// 				},
// 			}
// 		},
//	}
package modzy
