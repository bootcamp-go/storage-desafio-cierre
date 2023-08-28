package storage

import "errors"

// Customer is a struct that represents a customer
type Customer struct {
	Id        int
	FirstName string
	LastName  string
	Condition int
}

// CustomerConditionInfo is a struct that represents a customer condition info
type CustomerConditionInfo struct {
	Condition int
	Total     int
}

// CustomerAmountSpent is a struct that represents a customer amount spent
type CustomerAmountSpent struct {
	FirstName string
	LastName  string
	Amount	  float64
}

// StorageCustomer is an interface that represents a customer storage
type StorageCustomer interface {
	// ReadAll returns all customers
	ReadAll() (cs []*Customer, err error)

	// Create inserts a new customer
	Create(c *Customer) (err error)

	// ConditionInfo returns the total of customers based on their condition
	ConditionInfo() (cs []*CustomerConditionInfo, err error)

	// TopActiveCustomers returns the top active customers who have spent the most money
	TopActiveCustomers() (cs []*CustomerAmountSpent, err error)
}

var (
	// ErrStorageCustomerInternal is returned when an internal error occurs
	ErrStorageCustomerInternal = errors.New("internal storage error")
	// ErrStorageCustomerNotFound is returned when a customer is not found
	ErrStorageCustomerNotFound = errors.New("customer not found")
)