package client

import (
	"context"
	"encoding/json"
	"net/http"
)

// FilesService is the API client for the `/v1/files` endpoint
type FilesService service

// Get returns the file with the given ID.
//
// https://docs.lemonsqueezy.com/api/files#retrieve-a-file
func (service *FilesService) Get(ctx context.Context, fileID string) (*FileApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/files/"+fileID)
	if err != nil {
		return nil, response, err
	}

	file := new(FileApiResponse)
	if err = json.Unmarshal(*response.Body, file); err != nil {
		return nil, response, err
	}

	return file, response, nil
}

// List returns a paginated list of files.
//
// https://docs.lemonsqueezy.com/api/files#list-all-files
func (service *FilesService) List(ctx context.Context) (*FilesApiResponse, *Response, error) {
	response, err := service.client.do(ctx, http.MethodGet, "/v1/files")
	if err != nil {
		return nil, response, err
	}

	files := new(FilesApiResponse)
	if err = json.Unmarshal(*response.Body, files); err != nil {
		return nil, response, err
	}

	return files, response, nil
}
