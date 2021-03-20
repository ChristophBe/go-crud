package handlers

import "net/http"

func (c crudHandlersImpl) GetAll(w http.ResponseWriter, r *http.Request) {
	model, err := c.service.GetAll(r)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = c.responseWriter(model, http.StatusOK, w, r); err != nil {
		c.errorWriter(err, w, r)
	}

}
