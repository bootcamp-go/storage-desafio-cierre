package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// NewStorageProductJSONFile returns a new StorageProductJSONFile
func NewStorageProductJSONFile(file *os.File) *StorageProductJSONFile {
	return &StorageProductJSONFile{file: file}
}

// ProductJSONFile is an implementation of StorageProduct
type StorageProductJSONFile struct {
	file *os.File
}

// ProductJSONFile is a struct that represents a product in the JSON file
type ProductJSONFile struct {
	Id          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
}

// ReadAll returns all products
func (s *StorageProductJSONFile) ReadAll() (ps []*Product, err error) {
	var productsJSONFile []*ProductJSONFile
	err = json.NewDecoder(s.file).Decode(&productsJSONFile)
	if err != nil {
		err = fmt.Errorf("%w: %v", ErrStorageProductInternal, err)
		return
	}

	// serialization
	ps = make([]*Product, len(productsJSONFile))
	for i, productJSONFile := range productsJSONFile {
		ps[i] = &Product{
			Id:          productJSONFile.Id,
			Description: productJSONFile.Description,
			Price:       productJSONFile.Price,
		}
	}

	return
}

// Create inserts a new product
func (s *StorageProductJSONFile) Create(p *Product) (err error) {
	return
}

// TopSelled returns the top selled products
func (s *StorageProductJSONFile) TopSelled() (ps []*ProductSells, err error) {
	return
}