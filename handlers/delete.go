package handlers

import (
	"github.com/ChristophBe/go-crud/types"
	"net/http"
)

// NewDeleteHandler returns a http handler for that handles the deletion of one specific model.
func NewDeleteHandler(service types.GetOneService, responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		model, err := service.GetOne(request)
		if err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = model.Delete(request.Context()); err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = responseWriter(nil, http.StatusNoContent, writer, request); err != nil {
			errorWriter(err, writer, request)
		}
	}
}
