package client

import (
	"bytes"
	"net/http"
	"strconv"

	"catchpoint-provider/internal"
	"catchpoint-provider/internal/models"
)

// GetTest retrieves a test by its ID.
func GetTest(apiToken string, testID int64) (test *models.TestJSON, responseStatus string, err error) {
	getURL := internal.CatchpointTestURI + "/" + strconv.FormatInt(testID, 10) + "?showInheritedProperties=false"

	req, _ := http.NewRequest(http.MethodGet, getURL, nil)
	response, _, responseStatus, err := doAndUnmarshal[models.GetTestResponse](req, apiToken)

	if err != nil {
		return nil, responseStatus, err
	}

	err = response.ErrorIfIncomplete()
	if err != nil {
		return nil, responseStatus, err
	}

	test = &response.ResponseData.Tests[0]

	return
}

// CreateTest creates a new test with the provided JSON payload.
func CreateTest(apiToken string, jsonPayload string) (responseBody, responseStatus string, testID int64, err error) {
	var postBody = []byte(jsonPayload)

	req, _ := http.NewRequest(http.MethodPost, internal.CatchpointTestURI, bytes.NewBuffer(postBody))
	response, responseBody, responseStatus, err := doAndUnmarshal[models.Response](req, apiToken)

	if err != nil {
		return
	}

	err = response.ErrorIfIncomplete()
	if err != nil {
		return
	}

	testID, err = response.ResponseData.ID.Int64()
	if err != nil {
		return
	}

	return
}

// UpdateTest updates an existing test with the provided JSON payload.
func UpdateTest(apiToken string, testID int64, jsonPayload string) (responseBody, responseStatus string, completed bool, err error) {
	updateURL := internal.CatchpointTestURI + "/" + strconv.FormatInt(testID, 10)
	var jsonPatchDocument = []byte(jsonPayload)

	req, _ := http.NewRequest(http.MethodPatch, updateURL, bytes.NewBuffer(jsonPatchDocument))
	response, responseBody, responseStatus, err := doAndUnmarshal[models.Response](req, apiToken)

	if err != nil {
		return responseBody, responseStatus, false, err
	}

	err = response.ErrorIfIncomplete()
	if err != nil {
		return responseBody, responseStatus, false, err
	}

	completed = response.Completed

	return
}

// DeleteTest deletes a test by its ID and returns the response body, status, and completion status.
func DeleteTest(apiToken string, testID int64) (responseBody, responseStatus string, completed bool, err error) {
	deleteURL := internal.CatchpointTestURI + "/" + strconv.FormatInt(testID, 10)

	req, _ := http.NewRequest(http.MethodDelete, deleteURL, nil)
	response, responseBody, responseStatus, err := doAndUnmarshal[models.DeleteTestResponse](req, apiToken)

	if err != nil {
		return responseBody, responseStatus, false, err
	}

	err = response.ErrorIfIncomplete()
	if err != nil {
		return responseBody, responseStatus, false, err
	}

	completed = response.Completed

	return
}
