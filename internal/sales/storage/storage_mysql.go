package storage

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// NewStorageSaleMySQL returns a new instance of StorageSaleMySQL
func NewStorageSaleMySQL(db *sql.DB) *StorageSaleMySQL {
	return &StorageSaleMySQL{db: db}
}

// SaleMySQL is a struct that represents a sale in MySQL
type SaleMySQL struct {
	Id        sql.NullInt32
	Quantity  sql.NullInt32
	ProductId sql.NullInt32
	InvoiceId sql.NullInt32
}

// StorageSaleMySQL is a struct that represents a sale storage in MySQL for StorageSale interface
type StorageSaleMySQL struct {
	db *sql.DB
}

// ReadAll returns all sales
func (s *StorageSaleMySQL) ReadAll() (ss []*Sale, err error) {
	// query
	query := "SELECT id, quantity, product_id, invoice_id FROM sales"

	// prepare statement
	var stmt *sql.Stmt
	stmt, err = s.db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageSaleInternal, err)
		return
	}
	defer stmt.Close()

	// execute query
	var rows *sql.Rows
	rows, err = stmt.Query()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageSaleInternal, err)
		return
	}

	// iterate rows
	for rows.Next() {
		// scan row
		var saMySQL SaleMySQL
		err = rows.Scan(&saMySQL.Id, &saMySQL.Quantity, &saMySQL.ProductId, &saMySQL.InvoiceId)
		if err != nil {
			err = fmt.Errorf("%w. %v", ErrStorageSaleInternal, err)
			return
		}

		// serialization
		var sa Sale
		sa.Id = int(saMySQL.Id.Int32)
		sa.Quantity = int(saMySQL.Quantity.Int32)
		sa.ProductId = int(saMySQL.ProductId.Int32)
		sa.InvoiceId = int(saMySQL.InvoiceId.Int32)

		ss = append(ss, &sa)
	}
	rows.Close()

	return
}

// Create inserts a new sale
func (s *StorageSaleMySQL) Create(sa *Sale) (err error) {
	// deserialization
	var saMySQL SaleMySQL
	if sa.Id != 0 {
		saMySQL.Id.Valid = true
		saMySQL.Id.Int32 = int32(sa.Id)
	}
	if sa.Quantity != 0 {
		saMySQL.Quantity.Valid = true
		saMySQL.Quantity.Int32 = int32(sa.Quantity)
	}
	if sa.ProductId != 0 {
		saMySQL.ProductId.Valid = true
		saMySQL.ProductId.Int32 = int32(sa.ProductId)
	}
	if sa.InvoiceId != 0 {
		saMySQL.InvoiceId.Valid = true
		saMySQL.InvoiceId.Int32 = int32(sa.InvoiceId)
	}

	// query
	query := "INSERT INTO sales (quantity, product_id, invoice_id) VALUES (?, ?, ?)"

	// prepare statement
	var stmt *sql.Stmt
	stmt, err = s.db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageSaleInternal, err)
		return
	}
	defer stmt.Close()

	// execute query
	var result sql.Result
	result, err = stmt.Exec(saMySQL.Quantity, saMySQL.ProductId, saMySQL.InvoiceId)
	if err != nil {
		if mysqlErr, ok := err.(*mysql.MySQLError); ok {
			switch mysqlErr.Number {
			case 1452:
				err = fmt.Errorf("%w. %v", ErrStorageSaleRelation, err)
			default:
				err = fmt.Errorf("%w. %v", ErrStorageSaleInternal, err)
			}

			return
		}

		err = fmt.Errorf("%w. %v", ErrStorageSaleInternal, err)
		return
	}

	// check rows affected
	var rowsAffected int64
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageSaleInternal, err)
		return		
	}
	if rowsAffected != 1 {
		err = fmt.Errorf("%w. %s", ErrStorageSaleInternal, "rows affected != 1")
		return
	}

	// get last insert id
	var lastInsertId int64
	lastInsertId, err = result.LastInsertId()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageSaleInternal, err)
		return		
	}

	(*sa).Id = int(lastInsertId)

	return
}