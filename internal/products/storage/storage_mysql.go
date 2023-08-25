package storage

import (
	"database/sql"
	"fmt"
)

// NewStorageProductMySQL returns a new instance of StorageProductMySQL
func NewStorageProductMySQL(db *sql.DB) *StorageProductMySQL {
	return &StorageProductMySQL{db}
}

// ProductMySQL is a struct that represents a product in MySQL
type ProductMySQL struct {
	Id          sql.NullInt32
	Description sql.NullString
	Price       sql.NullFloat64
}

// StorageProductMySQL is a struct that represents a product storage in MySQL for StorageProduct interface
type StorageProductMySQL struct {
	db *sql.DB
}

// ReadAll returns all products
func (s *StorageProductMySQL) ReadAll() (ps []*Product, err error) {
	// query
	query := "SELECT id, description, price FROM products"

	// prepare statement
	var stmt *sql.Stmt
	stmt, err = s.db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageProductInternal, err)
		return
	}
	defer stmt.Close()

	// execute query
	var rows *sql.Rows
	rows, err = stmt.Query()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageProductInternal, err)
		return
	}

	// iterate rows
	for rows.Next() {
		// scan row
		var psMySQL ProductMySQL
		err = rows.Scan(&psMySQL.Id, &psMySQL.Description, &psMySQL.Price)
		if err != nil {
			err = fmt.Errorf("%w. %v", ErrStorageProductInternal, err)
			return
		}

		// serialization
		p := new(Product)
		if psMySQL.Id.Valid {
			p.Id = int(psMySQL.Id.Int32)
		}
		if psMySQL.Description.Valid {
			p.Description = psMySQL.Description.String
		}
		if psMySQL.Price.Valid {
			p.Price = psMySQL.Price.Float64
		}

		// append to list
		ps = append(ps, p)
	}

	return
}

// Create inserts a new product
func (s *StorageProductMySQL) Create(p *Product) (err error) {
	// deserialization
	var psMySQL ProductMySQL
	if p.Description != "" {
		psMySQL.Description.Valid = true
		psMySQL.Description.String = p.Description
	}
	if p.Price != 0 {
		psMySQL.Price.Valid = true
		psMySQL.Price.Float64 = p.Price
	}

	// query
	query := "INSERT INTO products (description, price) VALUES (?, ?)"

	// prepare statement
	var stmt *sql.Stmt
	stmt, err = s.db.Prepare(query)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageProductInternal, err)
		return
	}
	defer stmt.Close()

	// execute query
	var res sql.Result
	res, err = stmt.Exec(psMySQL.Description, psMySQL.Price)
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageProductInternal, err)
		return
	}

	// check rows affected
	var rowsAffected int64
	rowsAffected, err = res.RowsAffected()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageProductInternal, err)
		return
	}
	if rowsAffected != 1 {
		err = fmt.Errorf("%w. %v", ErrStorageProductInternal, "rows affected != 1")
		return
	}

	// get last insert id
	var id int64
	id, err = res.LastInsertId()
	if err != nil {
		err = fmt.Errorf("%w. %v", ErrStorageProductInternal, err)
		return
	}

	// update product id
	p.Id = int(id)

	return
}