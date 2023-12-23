package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrudHandlersImpl_Function(t *testing.T) {

	tt := []struct {
		name                string
		service             functionServiceMock
		responseWriterError error
		expectedError       error
	}{
		{
			name: "dto parsing failed",
			service: functionServiceMock{
				dtoErr: errors.New("test"),
			},
			expectedError: errors.New("test"),
		},
		{
			name: "dto invalid",
			service: functionServiceMock{
				dto: dtoMock[any]{
					validationError: errors.New("test"),
				},
			},
			expectedError: errors.New("test"),
		},
		{
			name: "function returns error invalid",
			service: functionServiceMock{
				dto:         dtoMock[any]{},
				functionErr: errors.New("test"),
			},
			expectedError: errors.New("test"),
		},
		{
			name: "function succeeded invalid",
			service: functionServiceMock{
				dto:            dtoMock[any]{},
				responseStatus: http.StatusOK,
			},
			expectedError: nil,
		},
		{
			name: "response writer returns error",
			service: functionServiceMock{
				dto:            dtoMock[any]{},
				responseStatus: http.StatusOK,
			},
			responseWriterError: errors.New("test-error"),
			expectedError:       errors.New("test-error"),
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			responseRecorder := new(responseWriterRecorder)
			errorRecorder := new(errorWriterRecorder)

			responseWriter := newMockResponseWriter(responseRecorder, tc.responseWriterError)

			errorWriter := newMockErrorWriter(errorRecorder)

			handler := NewFunctionHandler[dtoMock[any], any](tc.service, responseWriter, errorWriter)
			w := httptest.ResponseRecorder{}
			handler.ServeHTTP(&w, new(http.Request))

			if tc.expectedError != nil {

				// expect error writer to be called
				if errorRecorder.err == nil {
					t.Error("error to be not nil")
					return
				}
				if errorRecorder.err.Error() != tc.expectedError.Error() {
					t.Errorf("expected err to be %v, got %v", tc.expectedError, errorRecorder.err)
				}
				return
			}
			if tc.expectedError == nil {
				// expect response writer to be called

				if responseRecorder.status != tc.service.responseStatus {
					t.Errorf("expected response status to be %v, got %v", tc.service.responseStatus, responseRecorder.status)
				}
				if responseRecorder.body != nil {
					t.Errorf("expected response body to be nil, got %v", responseRecorder.body)
				}

			} else {
				// expect response not to called
				if responseRecorder.called {
					t.Error("expected response writer not to be called")
				}
			}
		})
	}
}
