package handlers

import (
	"context"
	"errors"
	"github.com/ChristophBe/go-crud/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockService struct {
	model mockModel
	err   error
}

func (m mockService) GetOne(_ *http.Request) (types.Model, error) {
	return m.model, m.err
}

type mockModel struct {
	value string
}

func (m mockModel) Create(ctx context.Context) (types.Model, error) {
	panic("implement me")
}

func (m mockModel) Update(ctx context.Context) (types.Model, error) {
	panic("implement me")
}

func (m mockModel) Delete(ctx context.Context) error {
	panic("implement me")
}

func TestCrudHandlersImpl_GetOne(t *testing.T) {

	expectedResponseStatus := http.StatusOK

	tt := []struct {
		name string
		mockService
		responseWriterError error
		expectedError       error
	}{
		{
			name: "getOne returns error",
			mockService: mockService{
				err: errors.New("test"),
			},
			expectedError: errors.New("test"),
		},
		{
			name:                "response writer returns error",
			mockService:         mockService{},
			responseWriterError: errors.New("test-error"),
			expectedError:       errors.New("test-error"),
		},
		{
			name: "getOne returns model",
			mockService: mockService{
				model: mockModel{value: "testValue"},
			},
			expectedError: nil,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			responseWriterValues := struct {
				called       bool
				model        interface{}
				resultStatus int
			}{
				called: false,
			}
			errorWriterValues := struct {
				err error
			}{}

			responseWriter := func(model interface{}, resultStatus int, _ http.ResponseWriter, _ *http.Request) error {
				responseWriterValues.called = true
				responseWriterValues.model = model
				responseWriterValues.resultStatus = resultStatus
				return tc.responseWriterError
			}

			errorWriter := func(err error, _ http.ResponseWriter, _ *http.Request) {
				errorWriterValues.err = err
			}

			handler := NewGetOneHandler(tc.mockService, responseWriter, errorWriter)
			w := httptest.ResponseRecorder{}
			handler.ServeHTTP(&w, new(http.Request))

			if tc.expectedError != nil {

				// expect error writer to be called
				if errorWriterValues.err == nil {
					t.Error("error to be not nil")
					return
				}
				if errorWriterValues.err.Error() != tc.expectedError.Error() {
					t.Errorf("expected err to be %v, got %v", tc.expectedError, errorWriterValues.err)
				}
				return
			}
			if tc.mockService.err == nil {
				// expect response writer to be called

				if responseWriterValues.resultStatus != expectedResponseStatus {
					t.Errorf("expected response status to be %v, got %v", expectedResponseStatus, responseWriterValues.resultStatus)
				}
				resultingModel, ok := responseWriterValues.model.(mockModel)
				if !ok {
					t.Fatal("failed to cast model")
				}

				if resultingModel.value != tc.model.value {
					t.Errorf("expected model model to be %v, got %v", tc.model.value, resultingModel.value)
				}

			} else {
				// expect response not to called
				if responseWriterValues.called {
					t.Error("expected response writer not to be called")
				}
			}
		})
	}

}
