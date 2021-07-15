package main

import (
	"fmt"
	"modzy.com/modzy-sdk/Modzy"
)

// The whole purpose of this main.go file is to test the functionality of the SDK and test the dev experience. Won't ship
func main() {
	modzy := Modzy.NewClient(" ", "")

	modzy.Models.GetAllModels(nil)
}

func getJobDetails(mc *Modzy.Client) {
	jobDetails := mc.Jobs.GetJobDetails("205a5d72-4868-45d1-b4f7-ca0eb050fa56")
	fmt.Println(jobDetails.Model.Name)

}
