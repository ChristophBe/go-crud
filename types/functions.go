package types

import "net/http"

// ErrorResponseWriter is used generate and write an error response.
type ErrorResponseWriter func(error, http.ResponseWriter, *http.Request)

// ResponseWriter is used to generate and write a response.
type ResponseWriter func(interface{}, int, http.ResponseWriter, *http.Request) error
