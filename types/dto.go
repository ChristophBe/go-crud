package types

// Dto is the type that contains the structure of the data that your api expect to receive.
// It contains a method to validate it self and to convert it to its corresponding model object.
type Dto interface {
	// IsValid validates the values of the dto.
	// If partial is true only values that are not there zero value should be validated. Otherwise validate all values.
	// It will return an error if the validation fails.
	IsValid(partial bool) error

	// AssignToModel assigns the value of the dto to a Model.
	AssignToModel(model Model) (Model, error)
}
