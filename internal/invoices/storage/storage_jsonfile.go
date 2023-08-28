package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

// NewStorageInvoiceJSONFile returns a new StorageInvoiceJSONFile
func NewStorageInvoiceJSONFile(file *os.File) *StorageInvoiceJSONFile {
	return &StorageInvoiceJSONFile{file: file}
}

// StorageInvoiceJSONFile is an implementation of StorageInvoice
type StorageInvoiceJSONFile struct {
	file *os.File
}

type InvoiceJSONFile struct {
	Id		   int		 `json:"id"`
	Datetime   time.Time `json:"datetime"`
	Total	   float64	 `json:"total"`
	CustomerId int		 `json:"customer_id"`
}

// ReadAll returns all invoices
func (s *StorageInvoiceJSONFile) ReadAll() (is []*Invoice, err error) {
	var invoicesJSONFile []*InvoiceJSONFile
	err = json.NewDecoder(s.file).Decode(&invoicesJSONFile)
	if err != nil {
		err = fmt.Errorf("%w: %v", ErrStorageInvoiceInternal, err)
		return
	}

	// serialization
	is = make([]*Invoice, len(invoicesJSONFile))
	for i, invoiceJSONFile := range invoicesJSONFile {
		is[i] = &Invoice{
			Id:		   invoiceJSONFile.Id,
			Datetime:   invoiceJSONFile.Datetime,
			Total:	   invoiceJSONFile.Total,
			CustomerId: invoiceJSONFile.CustomerId,
		}
	}

	return
}

// Create inserts a new invoice
func (s *StorageInvoiceJSONFile) Create(i *Invoice) (err error) {
	return
}