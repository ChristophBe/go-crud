package handlers

import (
	"errors"
	"github.com/ChristophBe/grud/types"
	"net/http"
	"net/http/httptest"
	"testing"
)

type replaceServiceMock struct {
	getOneServiceMock
	createEmptyModelServiceMock
	parseDtoFromRequestServiceMock
}

func TestCrudHandlersImpl_Replace(t *testing.T) {

	expectedResponseStatus := http.StatusAccepted

	tt := []struct {
		name                string
		service             types.ReplaceService
		responseWriterError error
		expectedError       error
		resultModel         modelMock
	}{
		{
			name: "parse dto form request turns error",
			service: replaceServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					err: errors.New("test"),
				},
			},
			expectedError: errors.New("test"),
		},

		{
			name: "dto is invalid",
			service: replaceServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					dto: dtoMock{
						validationError: errors.New("test"),
					},
				},
			},
			expectedError: errors.New("test"),
		},
		{
			name: "getting exiting model failed",
			service: replaceServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					dto: dtoMock{
						validationError: nil,
					},
				},
				getOneServiceMock: getOneServiceMock{
					err: errors.New("test"),
				},
			},
			expectedError: errors.New("test"),
		},
		{
			name: "dto assign to model failed",
			service: replaceServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					dto: dtoMock{
						assignModelResult: modelErrorHolder{
							err: errors.New("test"),
						},
					},
				},
			},
			expectedError: errors.New("test"),
		},
		{
			name: "model save model failed",
			service: replaceServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					dto: dtoMock{
						assignModelResult: modelErrorHolder{
							model: modelMock{
								updateResult: modelErrorHolder{
									err: errors.New("test"),
								},
							},
						},
					},
				},
				getOneServiceMock: getOneServiceMock{
					model: modelMock{},
				},
			},
			expectedError: errors.New("test"),
		},
		{
			name: "response writer returns error",
			service: replaceServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					dto: dtoMock{
						assignModelResult: modelErrorHolder{
							model: modelMock{
								createResult: modelErrorHolder{
									model: modelMock{
										value: "test-value",
									},
								},
							},
						},
					},
				},
			},
			responseWriterError: errors.New("test-error"),
			expectedError:       errors.New("test-error"),
		},
		{
			name: "success",
			service: replaceServiceMock{
				parseDtoFromRequestServiceMock: parseDtoFromRequestServiceMock{
					dto: dtoMock{
						assignModelResult: modelErrorHolder{
							model: modelMock{
								updateResult: modelErrorHolder{
									model: modelMock{
										value: "test-value",
									},
								},
							},
						},
					},
				},
			},
			expectedError: nil,
			resultModel: modelMock{
				value: "test-value",
			},
		},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {

			responseRecorder := new(responseWriterRecorder)
			errorRecorder := new(errorWriterRecorder)

			responseWriter := newMockResponseWriter(responseRecorder, tc.responseWriterError)

			errorWriter := newMockErrorWriter(errorRecorder)

			handler := NewReplaceHandler(tc.service, responseWriter, errorWriter)
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
				result, ok := responseRecorder.body.(modelMock)
				if !ok {
					t.Fatal("failed to cast model")
				}

				if tc.resultModel.value != result.value {
					t.Errorf("expected result value to be %v, got %v", tc.resultModel.value, result.value)
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
