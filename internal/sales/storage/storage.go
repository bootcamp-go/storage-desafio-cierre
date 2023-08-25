package storage

import "errors"

// Sale is a struct that represents a sale
type Sale struct {
	Id		   int
	Quantity   int
	ProductId  int
	InvoiceId  int
}

// StorageSale is an interface that represents a sale storage
type StorageSale interface {
	// ReadAll returns all sales
	ReadAll() (ss []*Sale, err error)

	// Create inserts a new sale
	Create(s *Sale) (err error)
}

var (
	// ErrStorageSaleInternal is returned when an internal error occurs
	ErrStorageSaleInternal = errors.New("internal storage error")
	// ErrStorageSaleNotFound is returned when a sale is not found
	ErrStorageSaleNotFound = errors.New("sale not found")
	// ErrStorageSaleRelation is returned when a sale relation is not found
	ErrStorageSaleRelation = errors.New("sale relation not found")
)