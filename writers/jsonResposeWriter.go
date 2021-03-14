package writers

import (
	"encoding/json"
	"net/http"
)

func JsonResponseWriter(responseBody interface{}, statusCode int, writer http.ResponseWriter, _ *http.Request) (err error) {
	jsonResponse, err := json.Marshal(responseBody)

	if err != nil {
		return
	}

	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.Header().Set("X-Content-Type-Options", "nosniff")
	writer.WriteHeader(statusCode)
	_, err = writer.Write(jsonResponse)
	return
}
