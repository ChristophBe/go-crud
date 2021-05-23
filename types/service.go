package types

import (
	"context"
	"net/http"
)

// GetOneService defines functions that are needed for GetOne.
type GetOneService interface {
	// GetOne returns one Model based on a request.
	GetOne(request *http.Request) (Model, error)
}

// GetAllService defines functions that are needed for GetAll.
type GetAllService interface {
	// GetAll returns a slice of Model based on a request.
	GetAll(request *http.Request) ([]Model, error)
}

// CreateEmptyModelService defines the CreateEmptyModel function that is used in multiple handlers.
type CreateEmptyModelService interface {
	// CreateEmptyModel returns a empty instance of the model
	CreateEmptyModel(ctx context.Context) Model
}

// ParseDtoFromRequestService defines the ParseDtoFromRequest function that is used in multiple handlers.
type ParseDtoFromRequestService interface {
	// ParseDtoFromRequest creates an dto instance based on a request
	ParseDtoFromRequest(request *http.Request) (Dto, error)
}

// CreateService defines functions that are need for the create model handler
type CreateService interface {
	CreateEmptyModelService
	ParseDtoFromRequestService
}

// UpdateService defines functions that are need for the update model handler
type UpdateService interface {
	GetOneService
	ParseDtoFromRequestService
}

// Service holds functions to retrieve Model instances  or create Dto objects.
type Service interface {
	ParseDtoFromRequestService
	CreateEmptyModelService
	GetOneService
	GetAllService
}
