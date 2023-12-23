package handlers

import (
	"context"
	"errors"
	"github.com/ChristophBe/grud/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

type testModel struct {
	Value string
}
type parseDtoFromRequestServiceMock struct {
	dto types.Dto[testModel]
	err error
}

func (p parseDtoFromRequestServiceMock) ParseDtoFromRequest(_ *http.Request) (types.Dto[testModel], error) {
	return p.dto, p.err
}

type createEmptyModelServiceMock struct {
	emptyModel types.Model
}

func (c createEmptyModelServiceMock) CreateEmptyModel(_ context.Context) types.Model {
	return c.emptyModel
}

type createModelServiceMock struct {
	createdModel testModel
	err          error
}

func (c createModelServiceMock) CreateModel(_ context.Context, _ testModel) (testModel, error) {
	return c.createdModel, c.err
}

type createServiceMock struct {
	createModelServiceMock
	parseDtoFromRequestServiceMock
}

func TestCrudHandlersImpl_Create(t *testing.T) {

	expectedResponseStatus := http.StatusAccepted

	validModel := testModel{
		Value: "value",
	}

	tt := []struct {
		name                string
		service             types.CreateService[testModel]
		responseWriterError error
		expectedError       error
		resultModel         testModel
	}{
		{
			name: "parse dto form request turns error",
			service: createServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					err: errors.New("test"),
				},
			},
			expectedError: errors.New("test"),
		},
		{
			name: "dto is invalid",
			service: createServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					dto: dtoMock[testModel]{
						validationError: errors.New("test"),
					},
				},
			},
			expectedError: errors.New("test"),
		},
		{
			name: "dto assign to model failed",
			service: createServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					dto: dtoMock[testModel]{
						assignModelResult: modelErrorHolder[testModel]{
							err: errors.New("test"),
						},
					},
				},
			},
			expectedError: errors.New("test"),
		},
		{
			name: "model save model failed",
			service: createServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					dto: dtoMock[testModel]{
						assignModelResult: modelErrorHolder[testModel]{
							model: testModel{Value: "test"},
						},
					},
				},
				createModelServiceMock: createModelServiceMock{
					err: errors.New("test"),
				},
			},
			expectedError: errors.New("test"),
		},
		{
			name: "response writer returns error",
			service: createServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					dto: dtoMock[testModel]{
						assignModelResult: modelErrorHolder[testModel]{
							model: validModel,
						},
					},
				},
			},
			responseWriterError: errors.New("test-error"),
			expectedError:       errors.New("test-error"),
		},
		{
			name: "success",
			service: createServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					dto: dtoMock[testModel]{
						assignModelResult: modelErrorHolder[testModel]{
							model: validModel,
						},
					},
				},
				createModelServiceMock: createModelServiceMock{
					createdModel: validModel,
				},
			},
			expectedError: nil,
			resultModel:   validModel,
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			responseRecorder := new(responseWriterRecorder)
			errorRecorder := new(errorWriterRecorder)

			responseWriter := newMockResponseWriter(responseRecorder, tc.responseWriterError)

			errorWriter := newMockErrorWriter(errorRecorder)

			handler := NewCreateHandler[testModel](tc.service, responseWriter, errorWriter)
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

				if responseRecorder.status != expectedResponseStatus {
					t.Errorf("expected response status to be %v, got %v", expectedResponseStatus, responseRecorder.status)
				}
				result, ok := responseRecorder.body.(testModel)
				if !ok {
					t.Fatal("failed to cast model")
				}

				if tc.resultModel.Value != result.Value {
					t.Errorf("expected result value to be %v, got %v", tc.resultModel.Value, result.Value)
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
