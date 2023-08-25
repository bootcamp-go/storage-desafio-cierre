package storage

import "errors"

// Customer is a struct that represents a customer
type Customer struct {
	Id        int
	FirstName string
	LastName  string
	Condition bool
}

// StorageCustomer is an interface that represents a customer storage
type StorageCustomer interface {
	// ReadAll returns all customers
	ReadAll() (cs []*Customer, err error)

	// Create inserts a new customer
	Create(c *Customer) (err error)
}

var (
	// ErrStorageCustomerInternal is returned when an internal error occurs
	ErrStorageCustomerInternal = errors.New("internal storage error")
	// ErrStorageCustomerNotFound is returned when a customer is not found
	ErrStorageCustomerNotFound = errors.New("customer not found")
)