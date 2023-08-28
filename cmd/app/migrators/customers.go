package migrators

import "app/internal/customers/storage"

// NewMigrationCustomer returns a new MigrationCustomer
func NewMigrationCustomer(storageDb storage.StorageCustomer, storageFile storage.StorageCustomer) *MigrationCustomer {
	return &MigrationCustomer{
		storageDb:   storageDb,
		storageFile: storageFile,
	}
}

// MigrationCustomer is a controller to migrate customers from one storage to another
type MigrationCustomer struct {
	storageDb   storage.StorageCustomer
	storageFile storage.StorageCustomer
}

// Migrate migrates customers from one storage to another
func (m *MigrationCustomer) Migrate() (err error) {
	var customers []*storage.Customer

	// read all customers from file
	customers, err = m.storageFile.ReadAll()
	if err != nil {
		return
	}

	// create customers in db
	for _, customer := range customers {
		err = m.storageDb.Create(customer)
		if err != nil {
			return
		}
	}

	return
}