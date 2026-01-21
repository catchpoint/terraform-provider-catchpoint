package client

import (
	"bytes"
	"net/http"
	"strconv"

	"catchpoint-provider/internal"
	"catchpoint-provider/internal/models"
)

func GetProduct(apiToken string, productId int64) (product *models.ProductJSON, responseStatus string, err error) {
	getURL := internal.CatchpointProductURI + "/" + strconv.FormatInt(productId, 10)

	req, _ := http.NewRequest(http.MethodGet, getURL, nil)
	response, _, responseStatus, err := doAndUnmarshal[models.ProductResponse](req, apiToken)
	if err != nil {
		return nil, responseStatus, err
	}

	err = response.ErrorIfIncomplete()
	if err != nil {
		return nil, responseStatus, err
	}

	product = &response.ResponseData.Products[0]

	return
}

func CreateProduct(apiToken string, jsonPayload string) (responseBody, responseStatus string, productID int64, err error) {
	var postBody = []byte(jsonPayload)

	req, _ := http.NewRequest(http.MethodPost, internal.CatchpointProductURI, bytes.NewBuffer(postBody))
	response, responseBody, responseStatus, err := doAndUnmarshal[models.Response](req, apiToken)

	if err != nil {
		return
	}
	err = response.ErrorIfIncomplete()
	if err != nil {
		return
	}

	productID, err = response.ResponseData.ID.Int64()
	if err != nil {
		return
	}

	return
}

func UpdateProduct(apiToken string, productId int64, jsonPayload string) (responseBody, responseStatus string, completed bool, err error) {
	updateURL := internal.CatchpointProductURI + "/" + strconv.FormatInt(productId, 10)
	var jsonPatchDocument = []byte(jsonPayload)

	req, _ := http.NewRequest(http.MethodPatch, updateURL, bytes.NewBuffer(jsonPatchDocument))
	response, responseBody, responseStatus, err := doAndUnmarshal[models.Response](req, apiToken)

	if err != nil {
		return
	}

	err = response.ErrorIfIncomplete()
	if err != nil {
		return
	}

	completed = response.Completed

	return
}
