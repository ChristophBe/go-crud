package writers

import (
	"fmt"
	"github.com/ChristophBe/go-crud/types"
	"net/http"
)

func JsonErrorResponseWriter(error error, writer http.ResponseWriter, request *http.Request) {
	crudError, ok := error.(types.CrudHandlerError)
	if ok {
		fmt.Printf("Request Failed with status: %d, massage: %s cause: %v", crudError.HttpStatus, crudError.Message, crudError.Cause)
		err := JsonResponseWriter(crudError, crudError.HttpStatus, writer, request)
		if err != nil {
			panic(err)
		}
	} else {
		JsonErrorResponseWriter(types.CrudHandlerError{
			Cause:      error,
			Message:    "unexpected error",
			HttpStatus: http.StatusInternalServerError,
		}, writer, request)
	}

}
