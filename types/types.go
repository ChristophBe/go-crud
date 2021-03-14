package types

import "net/http"

type Model interface {
	Assign(dto interface{}) error
	Create() error
	Update() error
	Delete() error
}

type Dto interface {
	IsValid(partial bool) error
	ConvertToModel() (Model, error)
}

type Service interface {
	ParseDtoFromRequest(request *http.Request) (Dto, error)
	GetOne(request *http.Request) (Model, error)
	GetAll(request *http.Request) ([]Model, error)
}

type CrudHandlers interface {
	Create(w http.ResponseWriter, r *http.Request)
	GetAll(w http.ResponseWriter, r *http.Request)
	GetOne(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Replace(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
}

type ErrorResponseWriter func(error, http.ResponseWriter, *http.Request)
type ResponseWriter func(interface{}, int, http.ResponseWriter, *http.Request) error
