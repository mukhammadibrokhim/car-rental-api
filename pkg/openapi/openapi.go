package openapi

import (
	"github.com/getkin/kin-openapi/openapi3"
)

// GenerateOpenAPISpec generates the OpenAPI 3.0 specification for the Car Rental API.
func GenerateOpenAPISpec() *openapi3.T {
	doc := &openapi3.T{
		OpenAPI: "3.0.0",
		Info: &openapi3.Info{
			Title:       "Car Rental API",
			Description: "API documentation for the Car Rental service",
			Version:     "1.0.0",
		},
		Servers: []*openapi3.Server{
			{
				URL: "http://localhost:8080",
			},
		},
	}
	return doc
}
