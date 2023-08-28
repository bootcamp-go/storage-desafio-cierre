// storage.go
package storage

import "errors"

// Product is a struct that represents a product
type Product struct {
	Id          int
	Description string
	Price       float64
}

// ProductSells is a struct that represents a product sells
type ProductSells struct {
	Description string
	Quantity	int
}

// StorageProduct is an interface that represents a product storage
type StorageProduct interface {
	// ReadAll returns all products
	ReadAll() (ps []*Product, err error)

	// Create inserts a new product
	Create(p *Product) (err error)

	// TopFiveSelled returns the top selled products
	TopSelled() (ps []*ProductSells, err error)
}

var (
	// ErrStorageProductInternal is returned when an internal error occurs
	ErrStorageProductInternal = errors.New("internal storage error")
	// ErrStorageProductNotFound is returned when a product is not found
	ErrStorageProductNotFound = errors.New("product not found")
)
