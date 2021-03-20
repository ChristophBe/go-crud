package types

import "fmt"

// CrudHandlerError holds information to create meaning full http error responses
type CrudHandlerError struct {
	Cause      error  `json:"-"`
	Message    string `json:"message"`
	HttpStatus int    `json:"-"`
}

func (c CrudHandlerError) Error() string {
	return fmt.Sprintf("request failed status: %d ; message: %s ; caused by: %v", c.HttpStatus, c.Message, c.Cause)
}
