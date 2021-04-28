package handlers

import (
	"context"
	"net/http"
)

func (c crudHandlersImpl) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	r = r.WithContext(ctx)

	model, err := c.service.GetOne(r)
	if err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = model.Delete(ctx); err != nil {
		c.errorWriter(err, w, r)
		return
	}

	if err = c.responseWriter(nil, http.StatusNoContent, w, r); err != nil {
		c.errorWriter(err, w, r)
	}
}
