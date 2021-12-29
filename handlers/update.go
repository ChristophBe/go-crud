package handlers

import (
	"github.com/ChristophBe/grud/types"
	"net/http"
)

// NewUpdateHandler creates a http.Handler that handles partial updates for existing models.
func NewUpdateHandler(service types.UpdateService, responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		dto, err := service.ParseDtoFromRequest(request)
		if err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = dto.IsValid(ctx, true); err != nil {
			errorWriter(err, writer, request)
			return
		}

		var model types.Model
		if model, err = service.GetOne(request); err != nil {
			errorWriter(err, writer, request)
			return
		}

		model, err = dto.AssignToModel(ctx, model)
		if err != nil {
			errorWriter(err, writer, request)
			return
		}

		if model, err = model.Update(ctx); err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = responseWriter(model, http.StatusAccepted, writer, request); err != nil {
			errorWriter(err, writer, request)
		}
	}
}
