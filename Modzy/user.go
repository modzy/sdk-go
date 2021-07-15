package Modzy

import (
	"github.com/dghubble/sling"
	"log"
	"modzy.com/modzy-sdk/Modzy/types"
	"strings"
)

type UserService struct {
	sling  *sling.Sling
	apiKey string
}

//TODO: Libraries do not log statements out. Need to refactor to return errors for handling by user

func newUserService(sling *sling.Sling, apiKey string) *UserService {
	return &UserService{
		sling:  sling.New().Path("/api/accounting/"),
		apiKey: apiKey,
	}
}

func (us *UserService) GetApiKeyDetails() types.KeyDetailResponse {
	keyPrefix := strings.Split(us.apiKey, ".")
	var response types.KeyDetailResponse
	_, err := us.sling.Path("access-keys/"+keyPrefix[0]).Receive(&response, nil)
	if err != nil {
		log.Println("Error getting key details:\n", err)
	}
	return response
}

func (us *UserService) GetAvailableModzyRoles() []types.Roles {
	var response []types.Roles
	_, err := us.sling.Path("roles").Receive(&response, nil)
	if err != nil {
		log.Println("Error getting key roles:\n", err)
	}
	return response
}
