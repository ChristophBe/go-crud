package types

// Model is the representation of one data Type that is managed by your api.
type Model interface {
	// Assign is used to to assign the values of an given dto the current value of the model.
	// It returns an error if it fails otherwise it returns nil.
	Assign(dto interface{}) error

	// Create is used to persist a new instance of that model.
	// It returns an error if it fails otherwise it returns nil.
	Create() error

	// Update is used to persist new the values of an existing model.
	// It returns an error if it fails otherwise it returns nil.
	Update() error

	// Delete deletes it self from the persisted records.
	// It returns an error if it fails otherwise it returns nil.
	Delete() error
}
