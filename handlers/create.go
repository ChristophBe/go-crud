package handlers

import (
	"github.com/ChristophBe/grud/types"
	"net/http"
)

// NewCreateHandler creates a http.Handler that handles the creation of a model
func NewCreateHandler[Model types.ModelTypeInterface, Dto types.Dto[Model]](service types.CreateService[Model, Dto], responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) http.HandlerFunc {
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

		var model Model
		model, err = dto.AssignToModel(ctx, model)
		if err != nil {
			errorWriter(err, writer, request)
			return
		}

		if model, err = service.CreateModel(ctx, model); err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = responseWriter(model, http.StatusAccepted, writer, request); err != nil {
			errorWriter(err, writer, request)
		}
	}
}
