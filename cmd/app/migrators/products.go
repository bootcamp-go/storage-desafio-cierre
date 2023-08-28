package migrators

import "app/internal/products/storage"

// NewMigrationProduct returns a new MigrationProduct
func NewMigrationProduct(storageDb storage.StorageProduct, storageFile storage.StorageProduct) *MigrationProduct {
	return &MigrationProduct{
		storageDb:   storageDb,
		storageFile: storageFile,
	}
}

// MigrationProduct is a controller to migrate products from one storage to another
type MigrationProduct struct {
	storageDb   storage.StorageProduct
	storageFile storage.StorageProduct
}

// Migrate migrates products from one storage to another
func (m *MigrationProduct) Migrate() (err error) {
	var products []*storage.Product

	// read all products from file
	products, err = m.storageFile.ReadAll()
	if err != nil {
		return
	}

	// create products in db
	for _, product := range products {
		err = m.storageDb.Create(product)
		if err != nil {
			return
		}
	}

	return
}