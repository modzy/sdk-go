package Modzy

import (
	"log"
	"modzy.com/modzy-sdk/Modzy/types"

	"github.com/dghubble/sling"
)

type ModelService struct {
	sling *sling.Sling
}

func newModelService(sling *sling.Sling) *ModelService {
	return &ModelService{
		sling: sling.New().Path("/api/models/"),
	}
}

// /api/models/all/versions/all
func (ms *ModelService) GetAllModels(params *types.AllModelRequestOpt) ([]types.ModelDetail, types.RequestError) {

	var modelVersions []types.ModelDetail
	// Definitely need to make use of this for all methods
	var reqErr types.RequestError

	_, e := ms.sling.New().Path("all/versions/all/").QueryStruct(params).Receive(&modelVersions, &reqErr)
	if e != nil {
		log.Fatalf("Error retrieving all models: %v", e)
	}

	return modelVersions, reqErr
}

// /api/models
func (ms *ModelService) GetModels(params *types.ModelRequestOpt) []types.Model {
	var models []types.Model

		_, err := ms.sling.QueryStruct(params).Receive(&models, nil)
		if err != nil {
			log.Fatalf("Error in retrieving all models: %v", err)
		}

	return models
}

// /api/models/{model_id}
func (ms *ModelService) GetModelById(id string) types.Model {
	var m types.Model

	_, e := ms.sling.New().Path(id).Receive(&m, nil)
	if e != nil {
		log.Fatalf("Error grabbing model with id %v: %v", id, e)
	}

	return m

}

// /api/models/{model_id}/related-models
func (ms *ModelService) GetRelatedModels(mid string) ([]types.ModelDetail, types.RequestError) {
	var md []types.ModelDetail
	var re types.RequestError
	_, e := ms.sling.New().Path(mid+"/related-models/").Receive(&md, &re)
	if e != nil {
		log.Fatalf("Error in retrieving related models: %v", e)
	}
	return md, re
}

func (ms *ModelService) GetModelVersions(id string) []types.ModelVersion {
	var m []types.ModelVersion

	_, e := ms.sling.New().Path(id+"/versions/").Receive(&m, nil)
	if e != nil {
		log.Fatalf("Error retrieving versions of model %v: %v", id, e)
	}
	return m
}
