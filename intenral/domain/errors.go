package domain

import "fmt"

type DomainValidationError struct {
	Message string
}

func (e *DomainValidationError) Error() string {
	return fmt.Sprintf("domain validation error message: %s", e.Message)
}

func NewDomainValidationError(message string) error {
	return &DomainValidationError{
		Message: message,
	}
}
