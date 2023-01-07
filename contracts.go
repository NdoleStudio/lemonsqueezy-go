package client

// ApiResponse represents an API response
type ApiResponse[T any, R any] struct {
	Jsonapi ApiResponseJSONAPI    `json:"jsonapi"`
	Links   ApiResponseLink       `json:"links"`
	Data    ApiResponseData[T, R] `json:"data"`
}

// ApiResponseList represents an API response with a list of items
type ApiResponseList[T any, R any] struct {
	Jsonapi ApiResponseJSONAPI      `json:"jsonapi"`
	Links   ApiResponseListLink     `json:"links"`
	Meta    ApiResponseListMeta     `json:"meta"`
	Data    []ApiResponseData[T, R] `json:"data"`
}

// ApiResponseData contains the api response data
type ApiResponseData[T any, R any] struct {
	Type          string          `json:"type"`
	ID            string          `json:"id"`
	Attributes    T               `json:"attributes"`
	Relationships R               `json:"relationships"`
	Links         ApiResponseLink `json:"links"`
}

// Resource returns a lemonsqueezy resource. It's similar to ApiResponseData but without the links
type Resource[T any] struct {
	Type       string `json:"type"`
	ID         string `json:"id"`
	Attributes T      `json:"attributes"`
}

// ApiResponseJSONAPI API version
type ApiResponseJSONAPI struct {
	Version string `json:"version"`
}

// ApiResponseLinks contains links to related resources
type ApiResponseLinks struct {
	Links ApiResponseLink `json:"links"`
}

// ApiResponseLink defines a link
type ApiResponseLink struct {
	Related string `json:"related,omitempty"`
	Self    string `json:"self"`
}

// ApiResponseListLink defines a link for list os resources
type ApiResponseListLink struct {
	First string `json:"first"`
	Last  string `json:"last"`
}

// ApiResponseListMeta defines the meta data for a list api response
type ApiResponseListMeta struct {
	Page ApiResponseListMetaPage `json:"page"`
}

// ApiResponseListMetaPage defines the pagination meta data for a list api response
type ApiResponseListMetaPage struct {
	CurrentPage int `json:"currentPage"`
	From        int `json:"from"`
	LastPage    int `json:"lastPage"`
	PerPage     int `json:"perPage"`
	To          int `json:"to"`
	Total       int `json:"total"`
}
