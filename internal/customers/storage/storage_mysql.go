package storage

import (
	"database/sql"
	"fmt"
)

// NewStorageCustomerMySQL returns a new instance of StorageCustomerMySQL
func NewStorageCustomerMySQL(db *sql.DB) *StorageCustomerMySQL {
	return &StorageCustomerMySQL{db}
}

// CustomerMySQL is a struct that represents a customer in MySQL
type CustomerMySQL struct {
	Id			sql.NullInt32
	FirstName	sql.NullString
	LastName	sql.NullString
	Condition	sql.NullBool
}

// StorageCustomerMySQL is a struct that represents a customer storage in MySQL for StorageCustomer interface
type StorageCustomerMySQL struct {
	db *sql.DB
}

// ReadAll returns all customers
func (s *StorageCustomerMySQL) ReadAll() (cs []*Customer, err error) {
	// query
	query := "SELECT id, first_name, last_name, condition FROM customers"
	
	// prepare statement
	var stmt *sql.Stmt
	stmt, err = s.db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageCustomerInternal, err)
		return
	}
	defer stmt.Close()

	// execute query
	var rows *sql.Rows
	rows, err = stmt.Query()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageCustomerInternal, err)
		return
	}

	// iterate rows
	for rows.Next() {
		// scan row
		var csMySQL CustomerMySQL
		err = rows.Scan(&csMySQL.Id, &csMySQL.FirstName, &csMySQL.LastName, &csMySQL.Condition)
		if err != nil {
			err = fmt.Errorf("%w. %v", ErrStorageCustomerInternal, err)
			return
		}

		// serialization
		c := new(Customer)
		if csMySQL.Id.Valid {
			c.Id = int(csMySQL.Id.Int32)
		}
		if csMySQL.FirstName.Valid {
			c.FirstName = csMySQL.FirstName.String
		}
		if csMySQL.LastName.Valid {
			c.LastName = csMySQL.LastName.String
		}
		if csMySQL.Condition.Valid {
			c.Condition = csMySQL.Condition.Bool
		}

		// append customer
		cs = append(cs, c)
	}

	return
}

// Create inserts a new customer
func (s *StorageCustomerMySQL) Create(c *Customer) (err error) {
	// deserialization
	var csMySQL CustomerMySQL
	if c.FirstName != "" {
		csMySQL.FirstName.Valid = true
		csMySQL.FirstName.String = c.FirstName
	}
	if c.LastName != "" {
		csMySQL.LastName.Valid = true
		csMySQL.LastName.String = c.LastName
	}
	if c.Condition {
		csMySQL.Condition.Valid = true
		csMySQL.Condition.Bool = c.Condition
	}

	// query
	query := "INSERT INTO customers (first_name, last_name, condition) VALUES (?, ?, ?)"

	// prepare statement
	var stmt *sql.Stmt
	stmt, err = s.db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageCustomerInternal, err)
		return
	}
	defer stmt.Close()

	// execute query
	var result sql.Result
	result, err = stmt.Exec(csMySQL.FirstName, csMySQL.LastName, csMySQL.Condition)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageCustomerInternal, err)
		return
	}

	// check rows affected
	var rowsAffected int64
	rowsAffected, err = result.RowsAffected()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageCustomerInternal, err)
		return
	}
	if rowsAffected != 1 {
		err = fmt.Errorf("%w. %s", ErrStorageCustomerInternal, "rows affected != 1")
		return
	}

	// get last insert id
	var lastInsertId int64
	lastInsertId, err = result.LastInsertId()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageCustomerInternal, err)
		return
	}

	// set last insert id
	c.Id = int(lastInsertId)

	return
}
