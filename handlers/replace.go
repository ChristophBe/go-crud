package handlers

import (
	"github.com/ChristophBe/grud/types"
	"net/http"
)

// NewReplaceHandler creates a http.Handler that handles replacing an exing model.
func NewReplaceHandler[M types.ModelTypeInterface, D types.Dto[M]](service types.ReplaceService[M, D], responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		var (
			dto   types.Dto[M]
			model M
			err   error
		)
		if dto, err = service.ParseDtoFromRequest(request); err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = dto.IsValid(ctx, false); err != nil {
			errorWriter(err, writer, request)
			return
		}

		if model, err = service.GetOne(request); err != nil {
			errorWriter(err, writer, request)
			return
		}

		if model, err = dto.AssignToModel(ctx, model); err != nil {
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
