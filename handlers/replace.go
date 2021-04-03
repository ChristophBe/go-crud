package handlers

import (
	"github.com/ChristophBe/go-crud/types"
	"net/http"
)

func (c crudHandlersImpl) Replace(w http.ResponseWriter, r *http.Request) {
	model, err := c.service.GetOne(r)

	var dto types.Dto
	if dto, err = c.service.ParseDtoFromRequest(r); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = dto.IsValid(false); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if model, err = dto.AssignToModel(c.service.CreateEmptyModel()); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if model, err = model.Update(); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = c.responseWriter(model, http.StatusAccepted, w, r); err != nil {
		c.errorWriter(err, w, r)
	}
}
