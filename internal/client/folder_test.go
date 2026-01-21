package client

import (
	"net/http"
	"testing"

	"catchpoint-provider/internal/testutil"
)

func TestGetFolder(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{"data":{"folders":[{"id":123,"name":"TestFolder"}]},"completed":true}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, OK200), nil
	}

	folder, status, err := GetFolder(Token, 123)

	testutil.AssertNil(t, "Error", err)
	testutil.AssertEqual(t, "Folder.ID", folder.ID, 123)
	testutil.AssertEqual(t, "Folder.Name", folder.Name, "TestFolder")
	testutil.AssertEqual(t, "ResponseStatus", status, OK200)
}

func TestGetFolderDetailed(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{
  "data": {
    "folders": [
      {
        "scheduleSetting": {
          "scheduleSettingType": {
            "id": 0,
            "name": "Inherit"
          }
        },
        "id": 108061,
        "divisionId": 3,
        "productId": 88801,
        "name": "Web Tests",
        "testFolderTypeId": {
          "id": 1,
          "name": "Synthetic"
        },
        "advancedSettings": {
          "advancedSettingType": {
            "id": 0,
            "name": "Inherit"
          }
        },
        "requestSetting": {
          "requestSettingType": {
            "id": 0,
            "name": "Inherit"
          }
        },
        "alertGroup": {
          "alertSettingType": {
            "id": 0,
            "name": "Inherit"
          },
          "alertGroupItems": []
        },
        "insights": {
          "insightSettingType": {
            "id": 0,
            "name": "Inherit"
          }
        }
      }
    ]
  },
  "messages": [],
  "errors": [],
  "completed": true,
  "traceId": "4000c824-0001-f700-b63f-84710c7967bb"
}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, OK200), nil
	}

	folder, status, err := GetFolder(Token, 108061)

	testutil.AssertNil(t, "Error", err)
	testutil.AssertEqual(t, "Folder.ID", folder.ID, 108061)
	testutil.AssertEqual(t, "Folder.Name", folder.Name, "Web Tests")
	testutil.AssertEqual(t, "ResponseStatus", status, OK200)

	testutil.AssertEqual(t, "Folder.ScheduleSetting.ScheduleSettingType.ID", folder.ScheduleSettings.ScheduleSettingType.ID, 0)
	testutil.AssertEqual(t, "Folder.AdvancedSettings.AdvancedSettingType.ID", folder.AdvancedSettings.AdvancedSettingType.ID, 0)
	testutil.AssertEqual(t, "Folder.RequestSettings.RequestSettingType.ID", folder.RequestSettings.RequestSettingType.ID, 0)
	testutil.AssertEqual(t, "Folder.InsightData.InsightSettingType.ID", folder.InsightData.InsightSettingType.ID, 0)
	testutil.AssertEqual(t, "Folder.AlertGroup.AlertSettingType.ID", folder.AlertGroup.AlertSettingType.ID, 0)
}

func TestCreateFolder(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{"data":{"id":"456"},"completed":true}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, Created201), nil
	}

	_, status, folderID, err := CreateFolder(Token, `{"name":"NewFolder"}`)
	testutil.AssertNil(t, "Error", err)

	testutil.AssertEqual(t, "FolderID", folderID, 456)
	testutil.AssertEqual(t, "Status", status, "201 created")
}

func TestUpdateFolder(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{"completed":true}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, OK200), nil
	}

	_, status, completed, err := UpdateFolder(Token, 789, `{"name":"UpdatedFolder"}`)
	testutil.AssertNil(t, "Error", err)

	testutil.AssertEqual(t, "Completed", completed, true)
	testutil.AssertEqual(t, "Status", status, OK200)
}
