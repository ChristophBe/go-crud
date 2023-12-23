package handlers

import (
	"context"
	"github.com/ChristophBe/grud/types"
	"net/http"
)

type getOneServiceMock struct {
	model testModel
	err   error
}

func (m getOneServiceMock) GetOne(_ *http.Request) (testModel, error) {
	return m.model, m.err
}

type getAllServiceMock struct {
	models []testModel
	err    error
}

func (m getAllServiceMock) GetAll(_ *http.Request) ([]testModel, error) {
	return m.models, m.err
}

type functionServiceMock struct {
	functionErr    error
	dtoErr         error
	result         any
	responseStatus int
	dto            dtoMock[any]
}

func (f functionServiceMock) Function(_ context.Context, _ dtoMock[any]) (any, int, error) {
	return f.result, f.responseStatus, f.functionErr
}
func (f functionServiceMock) ParseValidatableFromRequest(_ *http.Request) (dtoMock[any], error) {
	return f.dto, f.dtoErr
}

type modelErrorHolder[T any] struct {
	model T
	err   error
}

type updateModelServiceMock struct {
	model testModel
	err   error
}

func (u updateModelServiceMock) UpdateModel(_ context.Context, _ testModel) (testModel, error) {
	return u.model, u.err
}

type modelMock struct {
	value        string
	createResult modelErrorHolder[types.Model]
	updateResult modelErrorHolder[types.Model]
	deleteResult error
}

func (m modelMock) Create(_ context.Context) (types.Model, error) {
	return m.createResult.model, m.createResult.err
}

func (m modelMock) Update(_ context.Context) (types.Model, error) {
	return m.updateResult.model, m.updateResult.err
}

func (m modelMock) Delete(_ context.Context) error {
	return m.deleteResult
}

type errorWriterRecorder struct {
	called bool
	err    error
}

func newMockErrorWriter(recorder *errorWriterRecorder) types.ErrorResponseWriter {
	return func(err error, _ http.ResponseWriter, _ *http.Request) {
		recorder.called = true
		recorder.err = err
	}
}

type responseWriterRecorder struct {
	called bool
	body   interface{}
	status int
}

func newMockResponseWriter(recorder *responseWriterRecorder, err error) types.ResponseWriter {
	return func(body interface{}, status int, _ http.ResponseWriter, _ *http.Request) error {
		recorder.called = true
		recorder.body = body
		recorder.status = status
		return err
	}
}

type dtoMock[T any] struct {
	validationError   error
	assignModelResult modelErrorHolder[T]
}

func (d dtoMock[T]) IsValid(_ context.Context, _ bool) error {
	return d.validationError
}

func (d dtoMock[T]) AssignToModel(_ context.Context, _ T) (T, error) {
	return d.assignModelResult.model, d.assignModelResult.err
}
