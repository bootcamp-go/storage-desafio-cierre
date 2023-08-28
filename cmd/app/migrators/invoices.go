package migrators

import "app/internal/invoices/storage"

// NewMigrationInvoice returns a new MigrationInvoice
func NewMigrationInvoice(storageDb storage.StorageInvoice, storageFile storage.StorageInvoice) *MigrationInvoice {
	return &MigrationInvoice{
		storageDb:   storageDb,
		storageFile: storageFile,
	}
}

// MigrationInvoice is a controller to migrate invoices from one storage to another
type MigrationInvoice struct {
	storageDb   storage.StorageInvoice
	storageFile storage.StorageInvoice
}

// Migrate migrates invoices from one storage to another
func (m *MigrationInvoice) Migrate() (err error) {
	var invoices []*storage.Invoice

	// read all invoices from file
	invoices, err = m.storageFile.ReadAll()
	if err != nil {
		return
	}

	// create invoices in db
	for _, invoice := range invoices {
		err = m.storageDb.Create(invoice)
		if err != nil {
			return
		}
	}

	return
}