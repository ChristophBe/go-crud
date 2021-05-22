package handlers

import (
	"github.com/ChristophBe/go-crud/types"
	"net/http"
)

// NewGetOneHandler returns a http handler for handling requests one specific model.
func NewGetOneHandler(service types.GetOneService, responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		model, err := service.GetOne(r)
		if err != nil {
			errorWriter(err, w, r)
			return
		}
		if err = responseWriter(model, http.StatusOK, w, r); err != nil {
			errorWriter(err, w, r)
		}
	}
}

// GetOne returns a http handler for handling requests one specific model.
func (c crudHandlersImpl) GetOne(w http.ResponseWriter, r *http.Request) {
	NewGetOneHandler(c.service, c.responseWriter, c.errorWriter)(w, r)
}
