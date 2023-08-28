package migrators

import "app/internal/sales/storage"

// NewMigrationSale returns a new MigrationSale
func NewMigrationSale(storageDb storage.StorageSale, storageFile storage.StorageSale) *MigrationSale {
	return &MigrationSale{
		storageDb:   storageDb,
		storageFile: storageFile,
	}
}

// MigrationSale is a controller to migrate sales from one storage to another
type MigrationSale struct {
	storageDb   storage.StorageSale
	storageFile storage.StorageSale
}

// Migrate migrates sales from one storage to another
func (m *MigrationSale) Migrate() (err error) {
	var sales []*storage.Sale

	// read all sales from file
	sales, err = m.storageFile.ReadAll()
	if err != nil {
		return
	}

	// create sales in db
	for _, sale := range sales {
		err = m.storageDb.Create(sale)
		if err != nil {
			return
		}
	}

	return
}