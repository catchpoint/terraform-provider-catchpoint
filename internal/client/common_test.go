package client

import (
	"bytes"
	"io"
	"net/http"
)

const (
	OK200      = "200 ok"
	Created201 = "201 created"
	Token      = "fake-token"
)

func makeFakeResponse(body string, status string) *http.Response {
	return &http.Response{
		Status:     status,
		StatusCode: 200,
		Body:       io.NopCloser(bytes.NewBufferString(body)),
	}
}
