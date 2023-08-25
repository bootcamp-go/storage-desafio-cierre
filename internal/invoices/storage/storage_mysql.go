package storage

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

// NewStorageInvoiceMySQL returns a new instance of StorageInvoiceMySQL
func NewStorageInvoiceMySQL(db *sql.DB) *StorageInvoiceMySQL {
	return &StorageInvoiceMySQL{db: db}
}

// InvoiceMySQL is a struct that represents a invoice in MySQL
type InvoiceMySQL struct {
	Id         sql.NullInt32
	Datetime   sql.NullTime
	Total      sql.NullFloat64
	CustomerId sql.NullInt32
}

// StorageInvoiceMySQL is a struct that represents a invoice storage in MySQL for StorageInvoice interface
type StorageInvoiceMySQL struct {
	db *sql.DB
}

// ReadAll returns all invoices
func (s *StorageInvoiceMySQL) ReadAll() (is []*Invoice, err error) {
	// query
	query := "SELECT id, datetime, total, customer_id FROM invoices"

	// prepare statement
	var stmt *sql.Stmt
	stmt, err = s.db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageInvoiceInternal, err)
		return
	}
	defer stmt.Close()

	// execute query
	var rows *sql.Rows
	rows, err = stmt.Query()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageInvoiceInternal, err)
		return
	}

	// iterate rows
	for rows.Next() {
		// scan row
		var inMySQL InvoiceMySQL
		err = rows.Scan(&inMySQL.Id, &inMySQL.Datetime, &inMySQL.Total, &inMySQL.CustomerId)
		if err != nil {
			err = fmt.Errorf("%w. %v", ErrStorageInvoiceInternal, err)
			return
		}

		// serialization
		i := new(Invoice)
		if inMySQL.Id.Valid {
			i.Id = int(inMySQL.Id.Int32)
		}
		if inMySQL.Datetime.Valid {
			i.Datetime = inMySQL.Datetime.Time
		}
		if inMySQL.Total.Valid {
			i.Total = inMySQL.Total.Float64
		}
		if inMySQL.CustomerId.Valid {
			i.CustomerId = int(inMySQL.CustomerId.Int32)
		}

		// append to slice
		is = append(is, i)
	}

	return
}

// Create inserts a new invoice
func (s *StorageInvoiceMySQL) Create(i *Invoice) (err error) {
	// deserialization
	var inMySQL InvoiceMySQL
	if i.Datetime != (Invoice{}).Datetime {
		inMySQL.Datetime.Valid = true
		inMySQL.Datetime.Time = i.Datetime
	}
	if i.Total != (Invoice{}).Total {
		inMySQL.Total.Valid = true
		inMySQL.Total.Float64 = i.Total
	}
	if i.CustomerId != (Invoice{}).CustomerId {
		inMySQL.CustomerId.Valid = true
		inMySQL.CustomerId.Int32 = int32(i.CustomerId)
	}

	// query
	query := "INSERT INTO invoices (datetime, total, customer_id) VALUES (?, ?, ?)"

	// prepare statement
	var stmt *sql.Stmt
	stmt, err = s.db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageInvoiceInternal, err)
		return
	}
	defer stmt.Close()

	// execute query
	var result sql.Result
	result, err = stmt.Exec(inMySQL.Datetime, inMySQL.Total, inMySQL.CustomerId)
	if err != nil {
		errMySQL, ok := err.(*mysql.MySQLError)
		if ok {
			switch errMySQL.Number {
			case 1452:
				err = fmt.Errorf("%w. %v", ErrStorageInvoiceRelation, err)
			default:
				err = fmt.Errorf("%w. %v", ErrStorageInvoiceInternal, err)
			}
			return
		}

		err = fmt.Errorf("%w. %v", ErrStorageInvoiceInternal, err)
		return
	}

	// check rows affected
	var rowsAffected int64
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageInvoiceInternal, err)
		return
	}
	if rowsAffected != 1 {
		err = fmt.Errorf("%w. %s", ErrStorageInvoiceInternal, "rows affected != 1")
		return
	}

	// get last insert id
	var lastInsertId int64
	lastInsertId, err = result.LastInsertId()
	if err != nil {
		err = ErrStorageInvoiceInternal
		return
	}

	// set id
	i.Id = int(lastInsertId)

	return
}

