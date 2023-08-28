package storage

import (
	"encoding/json"
	"fmt"
	"os"
)

// NewStorageCustomerFile creates a new StorageCustomerJSONFile
func NewStorageCustomerFile(file *os.File) *StorageCustomerJSONFile {
	return &StorageCustomerJSONFile{file: file}
}

// StorageCustomerJSONFile is an implementation of the Storage interface
type StorageCustomerJSONFile struct {
	file *os.File
}

// CustomerJSONFile is a struct that represents a customer in the json file
type CustomerJSONFile struct {
	Id        int	 `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Condition bool   `json:"condition"`
}

// Read reads the file
func (s *StorageCustomerJSONFile) ReadAll() (cs []*Customer, err error) {
	// read the file
	var csj []*CustomerJSONFile
	err = json.NewDecoder(s.file).Decode(&csj)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageCustomerInternal, err)
		return
	}

	// serialization
	cs = make([]*Customer, len(csj))
	for i, c := range csj {
		cs[i] = &Customer{
			Id:        c.Id,
			FirstName: c.FirstName,
			LastName:  c.LastName,
			Condition: c.Condition,
		}
	}

	return
}

// Create creates a new customer
func (s *StorageCustomerJSONFile) Create(c *Customer) (err error) {
	return nil
}