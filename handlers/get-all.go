package handlers

import (
	"github.com/ChristophBe/go-crud/types"
	"net/http"
)

// NewGetAllHandler returns a http.Handler for handling requests a list of model.
func NewGetAllHandler(service types.GetAllService, responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		models, err := service.GetAll(request)
		if err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = responseWriter(models, http.StatusOK, writer, request); err != nil {
			errorWriter(err, writer, request)
		}
	}
}
