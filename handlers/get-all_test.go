package handlers

import (
	"errors"
	"github.com/ChristophBe/grud/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrudHandlersImpl_GetAll(t *testing.T) {

	expectedResponseStatus := http.StatusOK

	tt := []struct {
		name string
		getAllServiceMock
		responseWriterError error
		expectedError       error
	}{
		{
			name: "getAll returns error",
			getAllServiceMock: getAllServiceMock{
				err: errors.New("test"),
			},
			expectedError: errors.New("test"),
		},
		{
			name:                "response writer returns error",
			getAllServiceMock:   getAllServiceMock{},
			responseWriterError: errors.New("test-error"),
			expectedError:       errors.New("test-error"),
		},
		{
			name: "getAll returns a empty list",
			getAllServiceMock: getAllServiceMock{
				models: make([]types.Model, 0),
			},
			expectedError: nil,
		},
		{
			name: "getAll returns a list of models",
			getAllServiceMock: getAllServiceMock{
				models: []types.Model{
					modelMock{value: "a"},
					modelMock{value: "b"},
					modelMock{value: "c"},
				},
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

			handler := NewGetAllHandler(tc.getAllServiceMock, responseWriter, errorWriter)
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
			if tc.getAllServiceMock.err == nil {
				// expect response writer to be called

				if responseRecorder.status != expectedResponseStatus {
					t.Errorf("expected response status to be %v, got %v", expectedResponseStatus, responseRecorder.status)
				}

				resultModelList, ok := responseRecorder.body.([]types.Model)
				if !ok {
					t.Error("response is not a slice of models")
				}

				if len(resultModelList) != len(tc.models) {
					t.Errorf("expectet number of result models %d, got %d", len(tc.models), len(resultModelList))
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
