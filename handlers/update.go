package handlers

import (
	"github.com/ChristophBe/grud/types"
	"net/http"
)

// NewUpdateHandler creates a http.Handler that handles partial updates for existing models.
func NewUpdateHandler[M types.ModelTypeInterface, D types.Dto[M]](service types.UpdateService[M, D], responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) http.HandlerFunc {
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

		var model M
		if model, err = service.GetOne(request); err != nil {
			errorWriter(err, writer, request)
			return
		}

		model, err = dto.AssignToModel(ctx, model)
		if err != nil {
			errorWriter(err, writer, request)
			return
		}

		if model, err = service.UpdateModel(ctx, model); err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = responseWriter(model, http.StatusAccepted, writer, request); err != nil {
			errorWriter(err, writer, request)
		}
	}
}
