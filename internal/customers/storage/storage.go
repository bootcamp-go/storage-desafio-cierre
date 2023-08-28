package storage

import "errors"

// Customer is a struct that represents a customer
type Customer struct {
	Id        int
	FirstName string
	LastName  string
	Condition int
}

type CustomerConditionInfo struct {
	Condition int
	Total     int
}

// StorageCustomer is an interface that represents a customer storage
type StorageCustomer interface {
	// ReadAll returns all customers
	ReadAll() (cs []*Customer, err error)

	// Create inserts a new customer
	Create(c *Customer) (err error)

	// ConditionInfo returns the total of customers based on their condition
	ConditionInfo() (cs []*CustomerConditionInfo, err error)
}

var (
	// ErrStorageCustomerInternal is returned when an internal error occurs
	ErrStorageCustomerInternal = errors.New("internal storage error")
	// ErrStorageCustomerNotFound is returned when a customer is not found
	ErrStorageCustomerNotFound = errors.New("customer not found")
)