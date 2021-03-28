package handlers

import "net/http"

func (c crudHandlersImpl) Create(w http.ResponseWriter, r *http.Request) {

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

	if model, err = model.Create(); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = c.responseWriter(model, http.StatusAccepted, w, r); err != nil {
		c.errorWriter(err, w, r)
	}
}
