package handlers

import "net/http"

func (c crudHandlersImpl) Delete(w http.ResponseWriter, r *http.Request) {
	model, err := c.service.GetOne(r)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = model.Delete(); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = c.responseWriter(nil, http.StatusNoContent, w, r); err != nil {
		c.errorWriter(err, w, r)
	}
}
