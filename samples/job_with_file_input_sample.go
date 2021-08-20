package main

import (
    "bytes"
    "context"
    "github.com/joho/godotenv"
    modzy "github.com/modzy/sdk-go"
    "io/ioutil"
    "log"
    "os"
    "time"
)

var (
    ctx = context.TODO()
)

func main() {
    // The system admin can provide the right base API URL, the API key can be downloaded from your profile page on Modzy.
    // You can configure those params as is described in the README file (as environment variables, or by using the .env file),
    // or you can just update the BASE_URL and API_KEY variables and use this sample code (not recommended for production environments).
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }
    // The MODZY_BASE_URL should point to the API services route which may be different from the Modzy page URL.
    // (ie: https://modzy.example.com).
    baseURL := os.Getenv("MODZY_BASE_URL")
    // The MODZY_API_KEY is your own personal API key. It is composed by a public part, a dot character, and a private part
    // (ie: AzQBJ3h4B1z60xNmhAJF.uQyQh8putLIRDi1nOldh).
    apiKey := os.Getenv("MODZY_API_KEY")
    // Client initialization:
    //   Initialize the ApiClient instance with the BASE_URL and the API_KEY to store those arguments
    //   for the following API calls.
    client := modzy.NewClient(baseURL).WithAPIKey(apiKey)
    // Create a Job with a text input, wait, and retrieve results:
    // Get the model object:
    // If you already know the model identifier (i.e.: you got it from the URL of the model details page or from the input sample),
    // you can skip this step. If you don't, you can find the model identifier by using its name as follows:
    model, err := client.Models().GetModelDetailsByName(ctx, &modzy.GetModelDetailsByNameInput{Name: "Multi-Language OCR"})
    // Or if you already know the model id and want to know more about the model, you can use this instead:
    // model, err := client.Models().GetModelDetails(ctx, &modzy.GetModelDetailsInput{ModelID: "c60c8dbd79"})
    // You can find more information about how to query the models on the model_sample.go file.
    // The model identifier is under the ModelID key. You can take a look at the other properties under ModelDetails struct
    // Or just log the model identifier, and potencially the latest version
    log.Printf("The model identifier is %s and the latest version is %s\n", model.Details.ModelID, model.Details.LatestVersion)
    // Get the model version object:
    // If you already know the model version and the input key(s) of the model version you can skip this step. Also, you can
    // use the following code block to know about the inputs keys and skip the call on future job submissions.
    modelVersion, err := client.Models().GetModelVersionDetails(ctx, &modzy.GetModelVersionDetailsInput{ModelID: model.Details.ModelID, Version: model.Details.LatestVersion})
    if err != nil {
        log.Fatalf("Unexpected error %s", err)
        return
    }
    // The info stored in modelVersion provides insights about the amount of time that the model can spend processing, the inputs, and
    // output keys of the model.
    log.Printf("The model version is %s\n", modelVersion.Details.Version)
    log.Printf("  timeouts: status %dms, run %dms\n", modelVersion.Details.Timeout.Status, modelVersion.Details.Timeout.Run)
    log.Println("  inputs:")
    for _, input := range modelVersion.Details.Inputs {
        log.Printf("    key %s, type %s, description: %s\n", input.Name, input.AcceptedMediaTypes, input.Description)
    }
    log.Println("  outputs:")
    for _, output := range modelVersion.Details.Outputs {
        log.Printf("    key %s, type %s, description: %s\n", output.Name, output.MediaType, output.Description)
    }
    // Send the job:
    // With the info about the model (identifier), the model version (version string, input/output keys), you are ready to
    // submit the job. Just prepare the source map:
    // A file input can be a byte array or any file path. This input type fits for any size files.
    imagePath  := "./samples/image.png"
    configPath := "./samples/config.json"
    mapSource := make(map[string]modzy.FileInputItem)
    mapInput := make(modzy.FileInputItem)
    mapInput["input"] = modzy.FileInputFile(imagePath)
    mapInput["config.json"] = modzy.FileInputFile(configPath)
    mapSource["source-key"] = mapInput
    // An inference job groups input data that you send to a model. You can send any amount of inputs to
    // process and you can identify and refer to a specific input by the key that you assign, for example we can add:
    mapInput = make(modzy.FileInputItem)
    mapInput["input"] = modzy.FileInputFile(imagePath)
    mapInput["config.json"] = modzy.FileInputFile(configPath)
    mapSource["second-key"] = mapInput
    // You don't need to load all the inputs from files, you can just convert the files to bytes as follows:
    imageBytes, err := ioutil.ReadFile(imagePath)
    configBytes := []byte("{\"languages\":[\"spa\"]}")
    mapInput = make(modzy.FileInputItem)
    mapInput["input"] = modzy.FileInputReader(bytes.NewReader(imageBytes))
    mapInput["config.json"] = modzy.FileInputReader(bytes.NewReader(configBytes))
    mapSource["another-key"] = mapInput
    //If you send a wrong input key, the model fails to process the input.
    mapInput = make(modzy.FileInputItem)
    mapInput["a.wrong.key"] = modzy.FileInputFile(imagePath)
    mapInput["config.json"] = modzy.FileInputFile(configPath)
    mapSource["wrong-key"] = mapInput
    //If you send a correct input key but some wrong values, the model fails too.
    mapInput = make(modzy.FileInputItem)
    mapInput["input"] = modzy.FileInputFile(configPath)
    mapInput["config.json"] = modzy.FileInputFile(imagePath)
    mapSource["wrong-value"] = mapInput
    job, err := client.Jobs().SubmitJobFile(ctx, &modzy.SubmitJobFileInput{
        ModelIdentifier: model.Details.ModelID,
        ModelVersion: modelVersion.Details.Version,
        Inputs: mapSource,
    })
    if err != nil {
        log.Fatalf("Unexpected error %s", err)
        return
    }
    // Modzy creates the job and queue for processing. The job object contains all the info that you need to keep track
    // of the process, the most important being the job identifier and the job status.
    log.Printf("job: %s \n", job.Response.JobIdentifier)
    // The job moves to SUBMITTED, meaning that Modzy acknowledged the job and sent it to the queue to be processed.
    // We provide a helper method to listen until the job finishes processing. Its a good practice to set a max timeout
    // if you're doing a test (ie: 2*status+run). Otherwise, if the timeout is set to None, it will listen until the job finishes and moves to
    // COMPLETED, CANCELED, or TIMEOUT.
    job2, err := job.JobActions.WaitForCompletion(ctx, 20*time.Second)
    if err != nil {
        log.Fatalf("Unexpected error %s", err)
        return
    }
    // Get the results:
    // Check the status of the job. Jobs may be canceled or may reach a timeout.
    if job2.Details.Status == "COMPLETED" {
        // A completed job means that all the inputs were processed by the model. Check the results for each
        // input key provided in the source map to see the model output.
        results, err := job.JobActions.GetResults(ctx)
        // You can also get the results with the identifier (if you don't have the job object)
        //results, err := client.Jobs().GetJobResults(ctx, &modzy.GetJobResultsInput{JobIdentifier: job.Response.JobIdentifier})
        if err != nil {
            log.Fatalf("Unexpected error %s", err)
            return
        }
        // The result object has some useful info:
        log.Printf("Result: finished: %t, total: %d, completed: %d, failed: %d",
            results.Results.Finished,
            results.Results.Total,
            results.Results.Completed,
            results.Results.Failed)
        // Notice that we are iterating through the same input source keys
        for key, _ := range mapSource{
            // The result object has the individual results of each job input. In this case the output key is called
            // results.json, so we can get the results as follows:
            if result, exists := results.Results.Results[key]; exists {
                // The output for this model comes in a JSON format, so we can directly log the model results:
                log.Printf("    %s:\n", key)
                modelRes := result.Data["results.json"].(map[string]interface{})
                for key2, val2 := range modelRes {
                    log.Printf("      %s:%f\n", key2, val2)
                }
            } else {
                // If the model raises an error, we can get the specific error message:
                log.Fatalf("    %s: %s\n", key, result.Data)
            }
        }

    } else {
        log.Fatalf("The job ends with status %s", job2.Details.Status)
    }
}