package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// NewStorageSaleJSONFile returns a new StorageSaleJSONFile
func NewStorageSaleJSONFile(file *os.File) *StorageSaleJSONFile {
	return &StorageSaleJSONFile{file: file}
}

// StorageSaleJSONFile is an implementation of StorageSale
type StorageSaleJSONFile struct {
	file *os.File
}

// SaleJSONFile is a struct that represents a sale in the JSON file
type SaleJSONFile struct {
	Id         int `json:"id"`
	Quantity   int `json:"quantity"`
	ProductId  int `json:"product_id"`
	InvoiceId  int `json:"invoice_id"`
}

// ReadAll returns all sales
func (s *StorageSaleJSONFile) ReadAll() (ss []*Sale, err error) {
	var salesJSONFile []*SaleJSONFile
	err = json.NewDecoder(s.file).Decode(&salesJSONFile)
	if err != nil {
		err = fmt.Errorf("%w: %v", ErrStorageSaleInternal, err)
		return
	}

	// serialization
	ss = make([]*Sale, len(salesJSONFile))
	for i, saleJSONFile := range salesJSONFile {
		ss[i] = &Sale{
			Id:         saleJSONFile.Id,
			Quantity:   saleJSONFile.Quantity,
			ProductId:  saleJSONFile.ProductId,
			InvoiceId:  saleJSONFile.InvoiceId,
		}
	}

	return
}

// Create inserts a new sale
func (s *StorageSaleJSONFile) Create(sa *Sale) (err error) {
	return
}