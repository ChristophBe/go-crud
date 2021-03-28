package types

// Model is the representation of one data Type that is managed by your api.
type Model interface {
	// Assign is used to to assign the values of an given dto the current value of the model.
	// It returns a module with the values of the dto assigned to it, and a error if it fails and otherwise nil.
	Assign(dto interface{}) (Model, error)

	// Create is used to persist a new instance of that model.
	// It returns the created model and an error if it fails otherwise nil.
	Create() (Model, error)

	// Update is used to persist new the values of an existing model.
	// It returns the updated model and an error if it fails otherwise nil.
	Update() (Model, error)

	// Delete deletes it self from the persisted records.
	// It returns an error if it fails otherwise it returns nil.
	Delete() error
}
