package types

import (
	"context"
	"net/http"
)

type ModelTypeInterface any

// GetOneService defines functions that are needed for GetOne.
type GetOneService[M ModelTypeInterface] interface {
	// GetOne returns one Model based on a request.
	GetOne(request *http.Request) (M, error)
}

// GetAllService defines functions that are needed for GetAll.
type GetAllService[M ModelTypeInterface] interface {
	// GetAll returns a slice of Model based on a request.
	GetAll(request *http.Request) ([]M, error)
}

type CreateModelService[M ModelTypeInterface] interface {
	CreateModel(ctx context.Context, model M) (M, error)
}

type UpdateModelService[M ModelTypeInterface] interface {
	UpdateModel(ctx context.Context, model M) (M, error)
}
type DeleteModelService[M ModelTypeInterface] interface {
	DeleteModel(ctx context.Context, model M) error
}

// ParseDtoFromRequestService defines the ParseDtoFromRequest function that is used in multiple handlers.
type ParseDtoFromRequestService[M ModelTypeInterface] interface {
	// ParseDtoFromRequest creates a dto instance based on a request
	ParseDtoFromRequest(request *http.Request) (Dto[M], error)
}

// FunctionHandlerService defines a service to handle a request by a  Function
type FunctionHandlerService[Dto Validatable, Result any] interface {
	// ParseValidatableFromRequest parses a Validatable for the request
	ParseValidatableFromRequest(request *http.Request) (Dto, error)
	// Function a function generates a response based on a Validatable
	Function(ctx context.Context, dto Dto) (Result, int, error)
}

type DeleteService[M ModelTypeInterface] interface {
	DeleteModelService[M]
	GetOneService[M]
}

// CreateService defines functions that are need for the create model handler
type CreateService[M ModelTypeInterface] interface {
	CreateModelService[M]
	ParseDtoFromRequestService[M]
}

// UpdateService defines functions that are need for the update model handler
type UpdateService[M ModelTypeInterface] interface {
	UpdateModelService[M]
	GetOneService[M]
	ParseDtoFromRequestService[M]
}

// ReplaceService defines functions that are need for the replace model handler
type ReplaceService[M ModelTypeInterface] interface {
	UpdateModelService[M]
	GetOneService[M]
	ParseDtoFromRequestService[M]
}

// Service holds functions to retrieve Model instances  or create Dto objects.
type Service[M ModelTypeInterface] interface {
	ParseDtoFromRequestService[M]
	CreateModelService[M]
	UpdateModelService[M]
	GetOneService[M]
	GetAllService[M]
	DeleteService[M]
}
