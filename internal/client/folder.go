package client

import (
	"bytes"
	"net/http"
	"strconv"

	"catchpoint-provider/internal"
	"catchpoint-provider/internal/models"
)

// GetFolder retrieves a folder by its ID.
func GetFolder(apiToken string, folderId int64) (folder *models.FolderJSON, responseStatus string, err error) {
	getURL := internal.CatchpointFolderURI + "/" + strconv.FormatInt(folderId, 10) + "?showInheritedProperties=false"

	req, _ := http.NewRequest(http.MethodGet, getURL, nil)
	response, _, responseStatus, err := doAndUnmarshal[models.FolderResponse](req, apiToken)

	if err != nil {
		return nil, responseStatus, err
	}

	err = response.ErrorIfIncomplete()
	if err != nil {
		return nil, responseStatus, err
	}

	folder = &response.ResponseData.Folders[0]

	return
}

// CreateFolder creates a new folder with the provided JSON payload.
func CreateFolder(apiToken string, jsonPayload string) (responseBody, responseStatus string, folderID int64, err error) {
	postBody := []byte(jsonPayload)

	req, _ := http.NewRequest(http.MethodPost, internal.CatchpointFolderURI, bytes.NewBuffer(postBody))
	response, responseBody, responseStatus, err := doAndUnmarshal[models.Response](req, apiToken)

	if err != nil {
		return
	}
	err = response.ErrorIfIncomplete()
	if err != nil {
		return
	}

	folderID, err = response.ResponseData.ID.Int64()
	if err != nil {
		return
	}

	return
}

// UpdateFolder updates an existing folder with the provided JSON payload.
func UpdateFolder(apiToken string, folderId int64, jsonPayload string) (responseBody, responseStatus string, completed bool, err error) {
	updateURL := internal.CatchpointFolderURI + "/" + strconv.FormatInt(folderId, 10)
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
