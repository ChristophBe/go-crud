package types

import "net/http"

// Service holds functions to retrieve Model instances  or create Dto objects.
type Service interface {
	// ParseDtoFromRequest creates an dto instance based on a request
	ParseDtoFromRequest(request *http.Request) (Dto, error)

	// GetOne returns one Model based on a request.
	GetOne(request *http.Request) (Model, error)

	// GetAll returns a slice of Model based on a request.
	GetAll(request *http.Request) ([]Model, error)
}
