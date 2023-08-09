package lemonsqueezy

import "time"

// FileAttributes represents a digital good that can be downloaded by a customer after the product has been purchased.
type FileAttributes struct {
	VariantID     int       `json:"variant_id"`
	Identifier    string    `json:"identifier"`
	Name          string    `json:"name"`
	Extension     string    `json:"extension"`
	DownloadURL   string    `json:"download_url"`
	Size          int       `json:"size"`
	SizeFormatted string    `json:"size_formatted"`
	Version       string    `json:"version"`
	Sort          int       `json:"sort"`
	Status        string    `json:"status"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

// ApiResponseRelationshipsFile relationships of a file
type ApiResponseRelationshipsFile struct {
	Variant ApiResponseLinks `json:"variant"`
}

// FileApiResponse is the api response for one file
type FileApiResponse = ApiResponse[FileAttributes, ApiResponseRelationshipsFile]

// FilesApiResponse is the api response for a list of files.
type FilesApiResponse = ApiResponseList[FileAttributes, ApiResponseRelationshipsFile]
