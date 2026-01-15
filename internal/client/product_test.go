package client

import (
	"net/http"
	"testing"

	"catchpoint-provider/internal/testutil"
)

func TestGetProduct(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{"data":{"products":[{"id":123,"name":"TestProduct"}]},"completed":true}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, OK200), nil
	}

	product, status, err := GetProduct(Token, 123)

	testutil.AssertNil(t, "Error", err)
	testutil.AssertEqual(t, "Product.ID", product.ID, 123)
	testutil.AssertEqual(t, "Product.Name", product.Name, "TestProduct")
	testutil.AssertEqual(t, "ResponseStatus", status, OK200)
}

func TestCreateProduct(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{"data":{"id":456},"completed":true}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, Created201), nil
	}

	_, status, productID, err := CreateProduct(Token, `{"name":"NewProduct"}`)
	testutil.AssertNil(t, "Error", err)

	testutil.AssertEqual(t, "ProductID", productID, 456)
	testutil.AssertEqual(t, "Status", status, "201 created")
}

func TestUpdateProduct(t *testing.T) {
	orig := doFunc
	defer func() { doFunc = orig }()

	fakeBody := `{"completed":true}`
	doFunc = func(req *http.Request) (*http.Response, error) {
		return makeFakeResponse(fakeBody, OK200), nil
	}

	_, status, completed, err := UpdateProduct(Token, 789, `{"name":"UpdatedProduct"}`)
	testutil.AssertNil(t, "Error", err)

	testutil.AssertEqual(t, "Completed", completed, true)
	testutil.AssertEqual(t, "Status", status, OK200)
}
