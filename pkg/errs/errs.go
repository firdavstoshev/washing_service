package errs

import "errors"

var (
	ErrWashingServiceNotFound = errors.New("washing service not found")
	ErrCustomerNotFound       = errors.New("customer not found")
)
