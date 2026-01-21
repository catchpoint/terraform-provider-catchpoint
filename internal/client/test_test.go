package client

import (
	"errors"
	"net/http"
	"testing"

	"catchpoint-provider/internal/models"
	"catchpoint-provider/internal/testutil"
)

func TestGetTest(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{"data":{"tests":[{"id":123,"name":"TestTest"}]},"completed":true}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, OK200), nil
	}

	test, status, err := GetTest(Token, 123)

	testutil.AssertNil(t, "Error", err)
	testutil.AssertEqual(t, "Test.ID", test.ID, 123)
	testutil.AssertEqual(t, "Test.Name", test.Name, "TestTest")
	testutil.AssertEqual(t, "ResponseStatus", status, OK200)
}

func TestGetTestError(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{"data":{"tests":[{"id":123,"name":"TestTest"}]},"completed":true}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, OK200), errors.New("test error")
	}

	test, status, err := GetTest(Token, 123)

	testutil.AssertNotNil(t, "Error", err)
	testutil.AssertNil(t, "Test", test)
	testutil.AssertEqual(t, "ResponseStatus", status, OK200)
}

func TestGetTestNotFound(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{
  "data": {},
  "messages": [],
  "errors": [
    {
      "message": "No valid Test found for input '123'."
    }
  ],
  "completed": false,
  "traceId": "deb9cf7389cbf3da77fa973f7a045096"
}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, OK200), nil
	}

	test, status, err := GetTest(Token, 123)

	testutil.AssertNotNil(t, "Error", err)

	_, ok := err.(*models.ObjectNotFoundError)
	testutil.AssertEqual(t, "ErrorType", ok, true)

	testutil.AssertNil(t, "Test", test)
	testutil.AssertEqual(t, "ResponseStatus", status, OK200)
}

// If the JSON is invalid, the unmarshaling should fail and return its own error.
// This will return a nil test but still give the status of the HTTP response.
func TestGetInvalidJSON(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{invalidjson`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, OK200), nil
	}

	test, status, err := GetTest(Token, 123)

	testutil.AssertNotNil(t, "Error", err)
	testutil.AssertNil(t, "Test", test)
	testutil.AssertEqual(t, "ResponseStatus", status, OK200)
}

// If there's an error with no response, the body should be nil, and the status should be empty.
func TestGetTestNoResponse(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	doFunc = func(req *http.Request) (*http.Response, error) {
		return nil, errors.New("test error")
	}

	test, status, err := GetTest(Token, 123)

	testutil.AssertNotNil(t, "Error", err)
	testutil.AssertNil(t, "Test", test)
	testutil.AssertEqual(t, "ResponseStatus", status, "")
}

func TestCreateTest(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{"data":{"id":"456"},"completed":true}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, Created201), nil
	}

	body, status, testID, err := CreateTest(Token, `{"name":"NewTest"}`)
	testutil.AssertNil(t, "Error", err)

	testutil.AssertEqual(t, "TestID", testID, 456)
	testutil.AssertEqual(t, "Status", status, Created201)
	testutil.AssertEqual(t, "Body", body, fakeBody)
}

// For API errors, the test ID will be empty but we should still return a status code and body
// as the body may contain error details.
func TestCreateTestIDNotZero(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{
  "data": {},
  "messages": [],
  "errors": [
    {
      "message": "Test Id must be zero."
    }
  ],
  "completed": false,
  "traceId": "ea3c68cb4d22ffe3b0eb550ea742ba55"
}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, "400"), nil
	}

	body, status, testID, err := CreateTest(Token, `{"name":"NewTest"}`)
	testutil.AssertNotNil(t, "Error", err)
	testutil.AssertEqual(t, "TestID", testID, 0)
	testutil.AssertEqual(t, "Status", status, "400")
	testutil.AssertEqual(t, "Body", body, fakeBody)
}

func TestUpdateTest(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{"completed":true}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, OK200), nil
	}

	body, status, completed, err := UpdateTest(Token, 789, `{"name":"UpdatedTest"}`)
	testutil.AssertNil(t, "Error", err)

	testutil.AssertEqual(t, "Completed", completed, true)
	testutil.AssertEqual(t, "Status", status, OK200)
	testutil.AssertEqual(t, "Body", body, fakeBody)
}

func TestDeleteTest(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{"completed":true}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, OK200), nil
	}

	body, status, completed, err := DeleteTest(Token, 789)
	testutil.AssertNil(t, "Error", err)

	testutil.AssertEqual(t, "Completed", completed, true)
	testutil.AssertEqual(t, "Status", status, OK200)
	testutil.AssertEqual(t, "Body", body, fakeBody)
}
