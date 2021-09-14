package handlers

import (
	"github.com/ChristophBe/go-crud/types"
	"net/http"
)

// NewFunctionHandler creates a http.Handler that handles the creation of a model
func NewFunctionHandler(service types.FunctionHandlerService, responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		ctx := request.Context()
		dto, err := service.ParseValidatableFromRequest(request)
		if err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = dto.IsValid(ctx, false); err != nil {
			errorWriter(err, writer, request)
			return
		}

		result, status, err := service.Function(ctx, dto)
		if err != nil {
			errorWriter(err, writer, request)
			return
		}

		if err = responseWriter(result, status, writer, request); err != nil {
			errorWriter(err, writer, request)
		}
	}
}
