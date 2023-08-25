package storage

import (
	"errors"
	"time"
)

// Invoice is a struct that represents a invoice
type Invoice struct {
	Id		   int
	Datetime   time.Time
	Total	   float64
	CustomerId int
}

// StorageInvoice is an interface that represents a invoice storage
type StorageInvoice interface {
	// ReadAll returns all invoices
	ReadAll() (is []*Invoice, err error)

	// Create inserts a new invoice
	Create(i *Invoice) (err error)
}

var (
	// ErrStorageInvoiceInternal is returned when an internal error occurs
	ErrStorageInvoiceInternal = errors.New("internal storage error")
	// ErrStorageInvoiceNotFound is returned when a invoice is not found
	ErrStorageInvoiceNotFound = errors.New("invoice not found")
	// ErrStorageInvoiceRelation is returned when a invoice relation is not found
	ErrStorageInvoiceRelation = errors.New("invoice relation not found")
)