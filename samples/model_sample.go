package main

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	modzy "github.com/modzy/sdk-go"
)

func main() {
	// The system admin can provide the right base API URL, the API key can be downloaded from your profile page on Modzy.
	// You can configure those params as is described in the README file (as environment variables, or by using the .env file),
	// or you can just update the BASE_URL and API_KEY variables and use this sample code (not recommended for production environments).
	ctx := context.TODO()
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
	// Get all models:
	// You can get the full list of models from Modzy by using the get_all method to retrieve the identifier
	// and the latest version of each model
	listModelsInput := (&modzy.ListModelsInput{}).WithPaging(1000, 0)
	out, err := client.Models().ListModels(ctx, listModelsInput)
	if err != nil {
		log.Fatalf("Unexpected error %s", err)
		return
	}
	log.Printf("all models: %d\n", len(out.Models))
	// Also, use the get_models method to search for specific models:
	// Search by author:
	listModelsInput = (&modzy.ListModelsInput{}).WithPaging(1000, 0).WithFilterAnd(modzy.ListModelsFilterFieldAuthor, "Open Source")
	out, err = client.Models().ListModels(ctx, listModelsInput)
	if err != nil {
		log.Fatalf("Unexpected error %s", err)
		return
	}
	log.Printf("Open Source models: %d\n", len(out.Models))
	// Search for active models:
	listModelsInput = (&modzy.ListModelsInput{}).WithPaging(1000, 0).WithFilterAnd(modzy.ListModelsFilterFieldIsActive, "true")
	out, err = client.Models().ListModels(ctx, listModelsInput)
	if err != nil {
		log.Fatalf("Unexpected error %s", err)
		return
	}
	log.Printf("Active models: %d\n", len(out.Models))
	// Search by name (and paginate the results):
	listModelsInput = (&modzy.ListModelsInput{}).WithPaging(5, 0).WithFilterAnd(modzy.ListModelsFilterFieldName, "Image")
	out, err = client.Models().ListModels(ctx, listModelsInput)
	if err != nil {
		log.Fatalf("Unexpected error %s", err)
		return
	}
	log.Printf("Models with name start with 'Image': %d\n", len(out.Models))
	// Combined search:
	listModelsInput = (&modzy.ListModelsInput{}).WithPaging(1000, 0).
		WithFilterAnd(modzy.ListModelsFilterFieldName, "Image").
		WithFilterAnd(modzy.ListModelsFilterFieldAuthor, "Open Source").
		WithFilterAnd(modzy.ListModelsFilterFieldIsActive, "true")
	out, err = client.Models().ListModels(ctx, listModelsInput)
	if err != nil {
		log.Fatalf("Unexpected error %s", err)
		return
	}
	log.Printf("Active open source models which name starts with 'Image': %d\n", len(out.Models))
	// Get model details:
	// The models route only returns the modelId, latestVersion, and versions:
	for _, modelSummary := range out.Models {
		log.Println("Model: ", modelSummary)
		// Use the model identifier to get model details
		model, err := client.Models().GetModelDetails(ctx, &modzy.GetModelDetailsInput{ModelID: modelSummary.ID})
		if err != nil {
			log.Fatalf("Unexpected error %s", err)
			return
		}
		log.Println("Model Details: ", model)
		// Use the version identifier to get version details such as input and output details
		modelVersion, err := client.Models().GetModelVersionDetails(ctx, &modzy.GetModelVersionDetailsInput{ModelID: model.Details.ModelID, Version: model.Details.LatestVersion})
		if err != nil {
			log.Fatalf("Unexpected error %s", err)
			return
		}
		log.Println("Model Version detail keys: ", modelVersion)
		// then you'll get all the details about the specific model version
		log.Printf("ModelVersion Details %s\n", modelVersion.Details)
		// Probably the more interesting are the ones related with the inputs and outputs of the model
		log.Println("  inputs:")
		for _, input := range modelVersion.Details.Inputs {
			log.Printf("    key %s, type %s, description: %s\n", input.Name, input.AcceptedMediaTypes, input.Description)
		}
		log.Println("  outputs:")
		for _, output := range modelVersion.Details.Outputs {
			log.Printf("    key %s, type %s, description: %s\n", output.Name, output.MediaType, output.Description)
		}
	}
	// Get model by name:
	// You can also find models by name
	model, err := client.Models().GetModelDetailsByName(ctx, &modzy.GetModelDetailsByNameInput{Name: "Dataset Joining"})
	// this method returns the first matching model and its details
	log.Printf("Dataset Joining: id:%s, author: %s, is_active: %s, description: %s",
		model.Details.ModelID, model.Details.Author, model.Details.IsActive, model.Details.Description)

	// Finally, you can find related models related with this search:
	related, err := client.Models().GetRelatedModels(ctx, &modzy.GetRelatedModelsInput{ModelID: model.Details.ModelID})
	log.Println("related models")
	for _, model := range related.RelatedModels {
		log.Printf("    %s :: %s (%s)", model.ModelID, model.Name, model.Author)
	}

}
