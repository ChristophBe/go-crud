package handlers

import (
	"github.com/ChristophBe/grud/types"
	"net/http"
)

// CrudHandlers aggregates common crud http handlers.
type CrudHandlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Replace(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

// NewCrudHandlers creates a instance of CrudHandlers.
func NewCrudHandlers[M types.ModelTypeInterface](service types.Service[M], responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) CrudHandlers {
	return crudHandlersImpl[M]{
		service:        service,
		responseWriter: responseWriter,
		errorWriter:    errorWriter,
	}
}

type crudHandlersImpl[M types.ModelTypeInterface] struct {
	service        types.Service[M]
	responseWriter types.ResponseWriter
	errorWriter    types.ErrorResponseWriter
}

// Create is a http.Handler that handles the creation of a model
func (c crudHandlersImpl[M]) Create(writer http.ResponseWriter, request *http.Request) {
	NewCreateHandler[M](c.service, c.responseWriter, c.errorWriter).ServeHTTP(writer, request)
}

// GetAll is a http.Handler for fetch a list of model.
func (c crudHandlersImpl[M]) GetAll(writer http.ResponseWriter, request *http.Request) {
	NewGetAllHandler[M](c.service, c.responseWriter, c.errorWriter).ServeHTTP(writer, request)
}

// GetOne returns a http handler for handling requests one specific model.
func (c crudHandlersImpl[M]) GetOne(w http.ResponseWriter, r *http.Request) {
	NewGetOneHandler[M](c.service, c.responseWriter, c.errorWriter)(w, r)
}

// Update is a http.Handler that handles partial updates for existing models.
func (c crudHandlersImpl[M]) Update(writer http.ResponseWriter, request *http.Request) {
	NewUpdateHandler[M](c.service, c.responseWriter, c.errorWriter).ServeHTTP(writer, request)
}

// Replace is a http.Handler that handles replacing an exing model.
func (c crudHandlersImpl[M]) Replace(writer http.ResponseWriter, request *http.Request) {
	NewReplaceHandler[M](c.service, c.responseWriter, c.errorWriter).ServeHTTP(writer, request)
}

// Delete is a http handler for handling the deletion of specific model.
func (c crudHandlersImpl[M]) Delete(writer http.ResponseWriter, request *http.Request) {
	NewDeleteHandler[M](c.service, c.responseWriter, c.errorWriter).ServeHTTP(writer, request)
}
