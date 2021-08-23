# Modzy Golang SDK

![Modzy Logo](https://www.modzy.com/wp-content/uploads/2020/06/MODZY-RGB-POS.png)

<div align="center">

**Modzy's Java SDK queries models, submits inference jobs, and returns results directly to your editor.**


![GitHub contributors](https://img.shields.io/github/contributors/modzy/sdk-go)
![GitHub last commit](https://img.shields.io/github/last-commit/modzy/sdk-go)
![GitHub Release Date](https://img.shields.io/github/issues-raw/modzy/sdk-go)


[The job lifecycle](https://docs.modzy.com/reference/the-job-lifecycle) | [API Keys](https://docs.modzy.com/reference/api-keys-1) | [Samples](https://github.com/modzy/sdk-go/tree/main/samples) | [Documentation](https://docs.modzy.com/docs)

</div>

## Installation

Add the dependency

```bash
go get -u github.com/modzy/sdk-go
```


### Get your API key



API keys are security credentials required to perform API requests to Modzy. Our API keys are composed of an ID that is split by a dot into two parts: a public and private part.

The *public* part is the API keys' visible part only used to identify the key and by itself, it’s unable to perform API requests.

The *private* part is the public part's complement and it’s required to perform API requests. Since it’s not stored on Modzy’s servers, it cannot be recovered. Make sure to save it securely. If lost, you can [replace the API key](https://docs.modzy.com/reference/update-a-keys-body).


Find your API key in your user profile. To get your full API key click on "Get key":

<img src="key.png" alt="get key" width="10%"/>



## Initialize

Once you have a `model` and `version` identified, get authenticated with your API key.

```go
client := modzy.NewClient("http://url.to.modzy/api").WithAPIKey("API Key")
```

## Basic usage

### Browse models

Modzy’s Marketplace includes pre-trained and re-trainable AI models from industry-leading machine learning companies, accelerating the process from data to value.

The Model service drives the Marketplace and can be integrated with other applications, scripts, and systems. It provides routes to list, search, and filter model and model-version details.

[List models](https://docs.modzy.com/reference/list-models):

```go
out, err := client.Models().ListModels(ctx, input)
if err != nil {
	return err
}
for _, modelSummary := range out.Models {
	fmt.Println("Model: ", modelSummary)
}
```

Tags help categorize and filter models. They make model browsing easier.

[List tags](https://docs.modzy.com/reference/list-tags):

```go
out, err := client.Models().GetTags(ctx)
if err != nil {
    return err
}
for _, tag := range out.Tags {
    fmt.Println("Tag: ", tag)
}
```

[List models by tag](https://docs.modzy.com/reference/list-models-by-tag):

```go
out, err := client.Models().GetTagModels(ctx)
if err != nil {
    return err
}
for _, model := range out.Models {
    fmt.Println("Model: ", model)
}
```

### Get a model's details

Models accept specific *input file [MIME](https://developer.mozilla.org/en-US/docs/Web/HTTP/Basics_of_HTTP/MIME_types) types*. Some models may require multiple input file types to run data accordingly. In this sample, we use a model that requires `text/plain`.

Models require inputs to have a specific *input name* declared in the job request. This name can be found in the model’s details. In this sample, we use a model that requires `input.txt`.

Additionally, users can set their own input names. When multiple input items are processed in a job, these names are helpful to identify and get each input’s results. In this sample, we use a model that requires `input-1` and `input-2`.

[Get a model's details](https://docs.modzy.com/reference/list-model-details):

```go
out, err := client.Models().GetModelDetails(ctx, &modzy.GetModelDetailsInput{ModelID: "ed542963de"})
if err != nil {
    return err
}
fmt.Println("Model: ", out.Details)
```

Model specific sample requests are available in the version details and in the Model Details page.

[Get version details](https://docs.modzy.com/reference/get-version-details):

```go
out, err := client.Models().GetModelVersionDetails(ctx, &modzy.GetModelVersionDetailsInput{ModelID: "ed542963de", Version: "0.0.27"})
if err != nil {
    return err
}
// then you'll get all the details about the specific model version
fmt.Printf("ModelVersion Details %s\n", out.Details)
// Probably the more interesting are the ones related with the inputs and outputs of the model
fmt.Println("  inputs:")
for _, input := range out.Details.Inputs {
    fmt.Printf(
        "    key %s, type %s, description: %s\n", input.Name, input.AcceptedMediaTypes, input.Description
    )
}
fmt.Println("  outputs:")
for _, output := range out.Details.Outputs {
    fmt.Printf(
        "    key %s, type %s, description: %s\n", output.Name, output.MediaType, output.Description
    )
}
```

### Submit a job and get results

A *job* is the process that sends data to a model, sets the model to run the data, and returns results.

Modzy supports several *input types* such as `text`, `embedded` for Base64 strings, `aws-s3` and `aws-s3-folder` for inputs hosted in buckets, and `jdbc` for inputs stored in databases. In this sample, we use `text`.

[Here](https://github.com/modzy/sdk-go/blob/main/samples.adoc) are samples to submit jobs with `embedded`, `aws-s3`, `aws-s3-folder`, and `jdbc` input types.

[Submit a job with the model, version, and input items](https://docs.modzy.com/reference/create-a-job-1):

```go
submitResponse, err := client.Jobs().SubmitJobText(ctx, &modzy.SubmitJobTextInput{
    ModelIdentifier="ed542963de",
    ModelVersion="0.0.27",
    Inputs=map[string]string{
        "my-input": {
            "input.txt": "Modzy is great!"
        }
    }
})
```

[Hold until the inference is complete and results become available](https://docs.modzy.com/reference/get-job-details):

```go
jobDetails, err := submitResponse.JobActions.WaitForCompletion(ctx, 20*time.Second)
```

[Get the results](https://docs.modzy.com/reference/get-results):

Results are available per input item and can be identified with the name provided for each input item upon job request. You can also add an input name to the route and limit the results to any given input item.

Jobs requested for multiple input items may have partial results available prior to job completion.

```go
results, err := jobDetails.GetResults(ctx)
```

### Fetch errors

Errors may arise for different reasons. Fetch errors to know what is their cause and how to fix them.

Error      | Description
---------- | ---------
`ModzyHTTPError` | Wrapper for different errors, check code, message, url attributes.

Submitting jobs:

```go
submitResponse, err := client.Jobs().SubmitJobText(ctx, &modzy.SubmitJobTextInput{
	ModelIdentifier="ed542963de", 
	ModelVersion="0.0.27", 
	Inputs=map[string]string{
		"my-input": {
			"input.txt": "Modzy is great!"
		}
    }
})
if err != nil {
    log.Fatalf("The job submission fails with code %s and message %s", err.Status, err.Message)
    return
}
```

## Features

Modzy supports [batch processing](https://docs.modzy.com/reference/batch-processing), [explainability](https://docs.modzy.com/reference/explainability), and [model drift detection](https://docs.modzy.com/reference/model-drift-1).

## APIs

Here is a list of Modzy APIs. To see all the APIs, check our [Documentation](https://docs.modzy.com/reference/introduction).


| Feature | Code |Api route
| ---     | ---  | ---
|List models|client.Models().ListModels()|[api/models](https://docs.modzy.com/reference/list-models)|
|Get model details|client.Models().GetModelDetails()|[api/models/:model-id](https://docs.modzy.com/reference/list-model-details)|
|List models by name|client.Models().GetModelDetailsByName()|[api/models](https://docs.modzy.com/reference/list-models)|
|List models by tags|client.Models().GetTagsModels()|[api/models/tags/:tag-id](https://docs.modzy.com/reference/list-models-by-tag) |
|Get related models|client.Models().GetRelatedModels()|[api/models/:model-id/related-models](https://docs.modzy.com/reference/get-related-models)|
|Get a model's versions|client.Models().ListModelVersions()|[api/models/:model-id/versions](https://docs.modzy.com/reference/list-versions)|
|Get version details|client.Models().GetModelVersionsDetails()|[api/models/:model-id/versions/:version-id](https://docs.modzy.com/reference/get-version-details)|
|List tags|client.Models().ListTags()|[api/models/tags](https://docs.modzy.com/reference/list-tags)|
|Submit a Job (Text)|client.Jobs().SubmitJobText()|[api/jobs](https://docs.modzy.com/reference/create-a-job-1)|
|Submit a Job (Embedded)|client.Jobs().SubmitJobEmbedded()|[api/jobs](https://docs.modzy.com/reference/create-a-job-1)|
|Submit a Job (AWS S3)|client.Jobs().SubmitJobS3()|[api/jobs](https://docs.modzy.com/reference/create-a-job-1)|
|Submit a Job (JDBC)|client.Jobs().SubmitJobJDBC()|[api/jobs](https://docs.modzy.com/reference/create-a-job-1)|
|Cancel a job|lient.Jobs().CancelJob()|[api/jobs/:job-id](https://docs.modzy.com/reference/cancel-a-job)  |
|Hold until inference is complete|client.Jobs().WaitForJobCompletion()|[api/jobs/:job-id](https://docs.modzy.com/reference/get-job-details)  |
|Get job details|client.Jobs().GetJobDetails()|[api/jobs/:job-id](https://docs.modzy.com/reference/get-job-details)  |
|Get results|client.Jobs().getJobResults()|[api/results/:job-id](https://docs.modzy.com/reference/get-results)  |
|List the job history|client.Jobs().GetJobsHistory()|[api/jobs/history](https://docs.modzy.com/reference/list-the-job-history)  |

## Samples

Check out our [samples](https://github.com/modzy/sdk-go/tree/main/samples) for details on specific use cases.

To run samples:

Set the base url and api key in each sample file:

```go
// TODO: set the base url of modzy api and you api key
client := modzy.NewClient("http://url.to.modzy/api").WithAPIKey("API Key")
```

Or follow the instructions [here](https://github.com/modzy/sdk-go/tree/main/contributing.adoc#set-environment-variables-in-bash) to learn more.

And then, you can:

```bash
$ go run samples/model_samples.go
```
## Contributing

We are happy to receive contributions from all of our users. Check out our [contributing file](https://github.com/modzy/sdk-go/tree/main/contributing.adoc) to learn more.

## Code of conduct


[![Contributor Covenant](https://img.shields.io/badge/Contributor%20Covenant-v2.0%20adopted-ff69b4.svg)](https://github.com/modzy/sdk-go/tree/main/CODE_OF_CONDUCT.md)
