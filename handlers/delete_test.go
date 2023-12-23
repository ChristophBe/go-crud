package handlers

import (
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"
)

type deleteMockServiceMock struct {
	deletionError error
}

func (d deleteMockServiceMock) DeleteModel(_ context.Context, _ testModel) error {
	return d.deletionError
}

type deleteServiceMock struct {
	getOneServiceMock
	deleteMockServiceMock
}

func TestCrudHandlersImpl_Delete(t *testing.T) {

	expectedResponseStatus := http.StatusNoContent

	tt := []struct {
		name string
		deleteServiceMock
		responseWriterError error
		expectedError       error
	}{
		{
			name: "service returns error",
			deleteServiceMock: deleteServiceMock{
				getOneServiceMock: getOneServiceMock{
					err: errors.New("test"),
				},
			},
			expectedError: errors.New("test"),
		},
		{
			name: "response writer returns error",
			deleteServiceMock: deleteServiceMock{
				getOneServiceMock: getOneServiceMock{},
			},
			responseWriterError: errors.New("test-error"),
			expectedError:       errors.New("test-error"),
		},
		{
			name: "delete function returns error",
			deleteServiceMock: deleteServiceMock{
				getOneServiceMock: getOneServiceMock{
					model: testModel{"test-value"},
				},
				deleteMockServiceMock: deleteMockServiceMock{
					deletionError: errors.New("test-error"),
				},
			},

			expectedError: errors.New("test-error"),
		},
		{
			name:          "delete everything works fine",
			expectedError: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			responseRecorder := new(responseWriterRecorder)
			errorRecorder := new(errorWriterRecorder)

			responseWriter := newMockResponseWriter(responseRecorder, tc.responseWriterError)

			errorWriter := newMockErrorWriter(errorRecorder)

			handler := NewDeleteHandler[testModel](tc.deleteServiceMock, responseWriter, errorWriter)
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
