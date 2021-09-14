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
	// CreateEmptyModel returns an empty instance of the model
	CreateEmptyModel(ctx context.Context) Model
}

// ParseDtoFromRequestService defines the ParseDtoFromRequest function that is used in multiple handlers.
type ParseDtoFromRequestService interface {
	// ParseDtoFromRequest creates a dto instance based on a request
	ParseDtoFromRequest(request *http.Request) (Dto, error)
}

// FunctionHandlerService defines a service to handle a request by a  Function
type FunctionHandlerService interface {
	// ParseValidatableFromRequest parses a Validatable for the request
	ParseValidatableFromRequest(request *http.Request) (Validatable, error)
	// Function a function generates a response based on a Validatable
	Function(ctx context.Context, dto Validatable) (interface{}, int, error)
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

// ReplaceService defines functions that are need for the replace model handler
type ReplaceService interface {
	GetOneService
	CreateEmptyModelService
	ParseDtoFromRequestService
}

// Service holds functions to retrieve Model instances  or create Dto objects.
type Service interface {
	ParseDtoFromRequestService
	CreateEmptyModelService
	GetOneService
	GetAllService
}
