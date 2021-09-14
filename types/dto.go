package types

import "context"

// Validatable is an interface that makes sure the typ can be validated.
type Validatable interface {
	// IsValid validates the values of the dto.
	// If partial is true only values that are not there zero value should be validated. Otherwise, validate all values.
	// It will return an error if the validation fails.
	IsValid(ctx context.Context, partial bool) error
}

// Dto is the type that contains the structure of the data that your api expect to receive.
// It contains a method to validate itself and to convert it to its corresponding model object.
type Dto interface {
	Validatable

	// AssignToModel assigns the value of the dto to a Model.
	AssignToModel(ctx context.Context, model Model) (Model, error)
}
