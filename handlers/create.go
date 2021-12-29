package handlers

import (
	"github.com/ChristophBe/grud/types"
	"net/http"
)

// NewCreatHandler creates a http.Handler that handles the creation of a model
func NewCreatHandler(service types.CreateService, responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		dto, err := service.ParseDtoFromRequest(request)
		if err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = dto.IsValid(ctx, false); err != nil {
			errorWriter(err, writer, request)
			return
		}

		model, err := dto.AssignToModel(ctx, service.CreateEmptyModel(ctx))
		if err != nil {
			errorWriter(err, writer, request)
			return
		}

		if model, err = model.Create(ctx); err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = responseWriter(model, http.StatusAccepted, writer, request); err != nil {
			errorWriter(err, writer, request)
		}
	}
}
