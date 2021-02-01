package http

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"
)

type client struct {
	host  string
	token string
}

type httpClient interface {
	Get(u string) (*http.Response, error)
	Post(u string, files map[string]string) (*http.Response, error)
}

func newHttpClient(host, token string) httpClient {
	return &client{host, token}
}

func (c *client) makeRequest(req *http.Request) (*http.Response, error) {
	if c.token != "" {
		req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.token))
	}
	client := &http.Client{}
	return client.Do(req)
}

func (c *client) Get(u string) (*http.Response, error) {
	uri := fmt.Sprintf("%s%s", c.host, u)
	req, _ := http.NewRequest("GET", uri, nil)
	return c.makeRequest(req)
}

func (c *client) Post(u string, files map[string]string) (*http.Response, error) {
	var (
		uri    = fmt.Sprintf("%s%s", c.host, u)
		body   = new(bytes.Buffer)
		writer = multipart.NewWriter(body)
	)

	for field, path := range files {
		// gets file name from file path
		filename := filepath.Base(path)
		// creates a new form file writer
		fw, err := writer.CreateFormFile(field, filename)
		if err != nil {
			return nil, err
		}

		// prepare the file to be read
		file, err := os.Open(path)
		if err != nil {
			return nil, err
		}

		// copies the file content to the form file writer
		if _, err := io.Copy(fw, file); err != nil {
			return nil, err
		}
	}

	if err := writer.Close(); err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", uri, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())
	return c.makeRequest(req)
}
