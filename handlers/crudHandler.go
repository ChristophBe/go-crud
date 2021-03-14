package handlers

import (
	"github.com/ChristophBe/go-crud/types"
	"net/http"
)

type crudHandlerImpl struct {
	service        types.Service
	responseWriter types.ResponseWriter
	errorWriter    types.ErrorResponseWriter
}

func NewCrudHandler(service types.Service, responseWriter types.ResponseWriter, errorWriter types.ErrorResponseWriter) types.CrudHandlers {
	return crudHandlerImpl{
		service:        service,
		responseWriter: responseWriter,
		errorWriter:    errorWriter,
	}
}
func (c crudHandlerImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	model, err := c.service.GetAll(r)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}
	if err = c.responseWriter(model, http.StatusOK, w, r); err != nil {
		c.errorWriter(err, w, r)
	}

}
func (c crudHandlerImpl) GetOne(w http.ResponseWriter, r *http.Request) {
	model, err := c.service.GetOne(r)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}
	if err = c.responseWriter(model, http.StatusOK, w, r); err != nil {
		c.errorWriter(err, w, r)
	}

}
func (c crudHandlerImpl) Delete(w http.ResponseWriter, r *http.Request) {
	model, err := c.service.GetOne(r)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}
	if err := model.Delete(); err != nil {
		c.errorWriter(err, w, r)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
func (c crudHandlerImpl) getModelFromRequest(r *http.Request) (model types.Model, err error) {
	var dto types.Dto
	dto, err = c.service.ParseDtoFromRequest(r)
	if err != nil {
		return
	}

	if err = dto.IsValid(false); err != nil {
		return
	}
	model, err = dto.ConvertToModel()

	return
}
func (c crudHandlerImpl) Create(w http.ResponseWriter, r *http.Request) {

	dto, err := c.service.ParseDtoFromRequest(r)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = dto.IsValid(false); err != nil {
		c.errorWriter(err, w, r)
		return
	}
	model, err := dto.ConvertToModel()

	if err != nil {
		c.errorWriter(err, w, r)
		return
	}
	if err = model.Create(); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = c.responseWriter(model, http.StatusAccepted, w, r); err != nil {
		c.errorWriter(err, w, r)
	}
}
func (c crudHandlerImpl) Replace(w http.ResponseWriter, r *http.Request) {
	model, err := c.service.GetOne(r)

	var dto types.Dto
	dto, err = c.service.ParseDtoFromRequest(r)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = dto.IsValid(false); err != nil {
		c.errorWriter(err, w, r)
		return
	}
	if err = model.Assign(dto); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = model.Update(); err != nil {
		c.errorWriter(err, w, r)
		return
	}
	if err = c.responseWriter(model, http.StatusAccepted, w, r); err != nil {
		c.errorWriter(err, w, r)
	}
}
func (c crudHandlerImpl) Update(w http.ResponseWriter, r *http.Request) {

	dto, err := c.service.ParseDtoFromRequest(r)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = dto.IsValid(true); err != nil {
		c.errorWriter(err, w, r)
		return
	}
	var model types.Model

	if model, err = c.service.GetOne(r); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	err = model.Assign(dto)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}
	if err = model.Update(); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = c.responseWriter(model, http.StatusAccepted, w, r); err != nil {
		c.errorWriter(err, w, r)
	}
}
