package handlers

import (
	"context"
	"github.com/ChristophBe/go-crud/types"
	"net/http"
)

type getOneServiceMock struct {
	model modelMock
	err   error
}

func (m getOneServiceMock) GetOne(_ *http.Request) (types.Model, error) {
	return m.model, m.err
}

type modelErrorHolder struct {
	model types.Model
	err   error
}
type modelMock struct {
	value        string
	createResult modelErrorHolder
	updateResult modelErrorHolder
	deleteResult error
}

func (m modelMock) Create(_ context.Context) (types.Model, error) {
	return m.updateResult.model, m.updateResult.err
}

func (m modelMock) Update(_ context.Context) (types.Model, error) {
	return m.createResult.model, m.createResult.err
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
