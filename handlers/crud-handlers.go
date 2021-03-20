package handlers

import (
	"github.com/ChristophBe/go-crud/types"
	"net/http"
)

type CrudHandlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Replace(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type crudHandlersImpl struct {
	service        types.Service
	responseWriter types.ResponseWriter
	errorWriter    types.ErrorResponseWriter
}

func NewCrudHandlers(service types.Service, responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) CrudHandlers {
	return crudHandlersImpl{
		service:        service,
		responseWriter: responseWriter,
		errorWriter:    errorWriter,
	}
}
