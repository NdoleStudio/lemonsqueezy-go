package client

// ApiResponse represents an API response
type ApiResponse[T any] struct {
	Jsonapi ApiResponseJSONAPI `json:"jsonapi"`
	Links   ApiResponseLink    `json:"links"`
	Data    ApiResponseData[T] `json:"data"`
}

// ApiResponseData contains the api response data
type ApiResponseData[T any] struct {
	Type          string                   `json:"type"`
	ID            string                   `json:"id"`
	Attributes    T                        `json:"attributes"`
	Relationships ApiResponseRelationships `json:"relationships"`
	Links         ApiResponseLink          `json:"links"`
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
	Related string `json:"related"`
	Self    string `json:"self"`
}

// ApiResponseRelationships relationships of an object
type ApiResponseRelationships struct {
	Store     ApiResponseLinks `json:"store"`
	Order     ApiResponseLinks `json:"order"`
	OrderItem ApiResponseLinks `json:"order-item"`
	Product   ApiResponseLinks `json:"product"`
	Variant   ApiResponseLinks `json:"variant"`
}
