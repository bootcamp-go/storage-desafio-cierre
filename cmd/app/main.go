package main

import (
	"app/cmd/app/migrators"
	storageCustomers "app/internal/customers/storage"
	storageInvoices "app/internal/invoices/storage"
	storageProducts "app/internal/products/storage"
	storageSales "app/internal/sales/storage"
	"database/sql"
	"os"

	"github.com/go-sql-driver/mysql"
)

func main() {
	// env
	// ...

	// dependencies
	// -> database
	cfgMySQL := mysql.Config{
		User:                 os.Getenv("DB_MYSQL_USER"),
		Passwd:               os.Getenv("DB_MYSQL_PASS"),
		Addr:                 os.Getenv("DB_MYSQL_HOST") + ":" + os.Getenv("DB_MYSQL_PORT"),
		DBName:               os.Getenv("DB_MYSQL_NAME"),
		Net:                  "tcp",
		AllowNativePasswords: true,
	}
	db, err := sql.Open("mysql", cfgMySQL.FormatDSN())
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}

	// -> files
	fileCustomers, err := os.Open(os.Getenv("STORAGE_FILE_CUSTOMER"))
	if err != nil {
		panic(err)
	}
	defer fileCustomers.Close()

	fileInvoices, err := os.Open(os.Getenv("STORAGE_FILE_INVOICE"))
	if err != nil {
		panic(err)
	}
	defer fileInvoices.Close()

	fileProducts, err := os.Open(os.Getenv("STORAGE_FILE_PRODUCT"))
	if err != nil {
		panic(err)
	}
	defer fileProducts.Close()

	fileSales, err := os.Open(os.Getenv("STORAGE_FILE_SALE"))
	if err != nil {
		panic(err)
	}
	defer fileSales.Close()

	// -> storages
	stFileCustomers := storageCustomers.NewStorageCustomerFile(fileCustomers)
	stDbCustomers := storageCustomers.NewStorageCustomerMySQL(db)

	stFileInvoices := storageInvoices.NewStorageInvoiceJSONFile(fileInvoices)
	stDbInvoices := storageInvoices.NewStorageInvoiceMySQL(db)

	stFileProducts := storageProducts.NewStorageProductJSONFile(fileProducts)
	stDbProducts := storageProducts.NewStorageProductMySQL(db)

	stFileSales := storageSales.NewStorageSaleJSONFile(fileSales)
	stDbSales := storageSales.NewStorageSaleMySQL(db)

	// -> migrators
	migratorCustomers := migrators.NewMigrationCustomer(stDbCustomers, stFileCustomers)
	migratorInvoices := migrators.NewMigrationInvoice(stDbInvoices, stFileInvoices)
	migratorProducts := migrators.NewMigrationProduct(stDbProducts, stFileProducts)
	migratorSales := migrators.NewMigrationSale(stDbSales, stFileSales)
	
	// run
	// -> customers
	if err := migratorCustomers.Migrate(); err != nil {
		panic(err)
	}

	// -> invoices
	if err := migratorInvoices.Migrate(); err != nil {
		panic(err)
	}

	// -> products
	if err := migratorProducts.Migrate(); err != nil {
		panic(err)
	}

	// -> sales
	if err := migratorSales.Migrate(); err != nil {
		panic(err)
	}
}