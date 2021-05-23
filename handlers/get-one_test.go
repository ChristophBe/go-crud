package handlers

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrudHandlersImpl_GetOne(t *testing.T) {

	expectedResponseStatus := http.StatusOK

	tt := []struct {
		name string
		getOneServiceMock
		responseWriterError error
		expectedError       error
	}{
		{
			name: "getOne returns error",
			getOneServiceMock: getOneServiceMock{
				err: errors.New("test"),
			},
			expectedError: errors.New("test"),
		},
		{
			name:                "response writer returns error",
			getOneServiceMock:   getOneServiceMock{},
			responseWriterError: errors.New("test-error"),
			expectedError:       errors.New("test-error"),
		},
		{
			name: "getOne returns model",
			getOneServiceMock: getOneServiceMock{
				model: modelMock{value: "testValue"},
			},
			expectedError: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			responseRecorder := new(responseWriterRecorder)
			errorRecorder := new(errorWriterRecorder)

			responseWriter := newMockResponseWriter(responseRecorder, tc.responseWriterError)

			errorWriter := newMockErrorWriter(errorRecorder)

			handler := NewGetOneHandler(tc.getOneServiceMock, responseWriter, errorWriter)
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
			if tc.getOneServiceMock.err == nil {
				// expect response writer to be called

				if responseRecorder.status != expectedResponseStatus {
					t.Errorf("expected response status to be %v, got %v", expectedResponseStatus, responseRecorder.status)
				}
				resultingModel, ok := responseRecorder.body.(modelMock)
				if !ok {
					t.Fatal("failed to cast model")
				}

				if resultingModel.value != tc.model.value {
					t.Errorf("expected model model to be %v, got %v", tc.model.value, resultingModel.value)
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
