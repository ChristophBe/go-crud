package handlers

import (
	"github.com/ChristophBe/go-crud/types"
	"net/http"
)

// NewReplaceHandler creates a http.Handler that handles replacing an exing model.
func NewReplaceHandler(service types.ReplaceService, responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()

		var (
			dto   types.Dto
			model types.Model
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

		if model, err = model.Update(ctx); err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = responseWriter(model, http.StatusAccepted, writer, request); err != nil {
			errorWriter(err, writer, request)
		}
	}
}

// Replace is a http.Handler that handles replacing an exing model.
func (c crudHandlersImpl) Replace(writer http.ResponseWriter, request *http.Request) {
	NewReplaceHandler(c.service, c.responseWriter, c.errorWriter).ServeHTTP(writer, request)
}
