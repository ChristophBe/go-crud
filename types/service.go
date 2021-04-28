package types

import (
	"context"
	"net/http"
)

// Service holds functions to retrieve Model instances  or create Dto objects.
type Service interface {
	// ParseDtoFromRequest creates an dto instance based on a request
	ParseDtoFromRequest(request *http.Request) (Dto, error)

	// CreateEmptyModel returns a empty instance of the model
	CreateEmptyModel(ctx context.Context) Model

	// GetOne returns one Model based on a request.
	GetOne(request *http.Request) (Model, error)

	// GetAll returns a slice of Model based on a request.
	GetAll(request *http.Request) ([]Model, error)
}
