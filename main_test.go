package main

import (
	"bytes"
	"errors"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockClient struct {
	URL        string
	statusCode int
	status     string
	body       string
	err        string
}

func (client *mockClient) Get() (resp *http.Response, err error) {
	resp = &http.Response{
		Status:     client.status,
		StatusCode: client.statusCode,
		Body:       ioutil.NopCloser(bytes.NewBufferString(client.body)),
	}

	err = errors.New(client.err)

	return
}

// TestFetch will test the HTTP client get
func TestFetch(t *testing.T) {
	client := &mockClient{
		URL:        "whatever",
		statusCode: http.StatusNotFound,
		status:     "404 NotFound",
		body:       "",
		err:        "404",
	}

	resp, err := fetch(client)

	assert.Equal(t, 404, resp.StatusCode)
	assert.Equal(t, "404", err.Error())
}
