package storage

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// Tests for StorageProductMySQL
func TestStorageProductMySQL_TopSelled(t *testing.T) {
	// type input struct {}
	type output struct { ps []*ProductSells; err error; errMsg string }
	type testCase struct {
		name string
		// input input
		output output
		// set-up
		setUpCfgDB func (cfg *mysql.Config)
		setUpDatabase func (db *sql.DB) (err error)
	}

	// test cases
	cases := []testCase{
		// valid case
		{
			name: "valid case",
			output: output{
				ps: []*ProductSells{
					{Description: "Product 1", Quantity: 2},
					{Description: "Product 2", Quantity: 1},
				},
				err: nil,
				errMsg: "",
			},
			setUpCfgDB: func (cfg *mysql.Config) {
				cfg.User = os.Getenv("DB_MYSQL_USER")
				cfg.Passwd = os.Getenv("DB_MYSQL_PASSWORD")
				cfg.Net = "tcp"
				cfg.Addr = os.Getenv("DB_MYSQL_ADDR")
				cfg.DBName = os.Getenv("DB_MYSQL_DATABASE")
				cfg.ParseTime = true
			},
			setUpDatabase: func (db *sql.DB) (err error) {
				// insert customers
				query := "INSERT INTO customers (first_name, last_name, `condition`) VALUES (?, ?, ?)"
				stmt, err := db.Prepare(query)
				if err != nil {
					return
				}
				defer stmt.Close()

				_, err = stmt.Exec("John", "Doe", 0)
				if err != nil {
					return
				}
				stmt.Close()

				// insert invoices
				query = "INSERT INTO invoices (`datetime`, total, customer_id) VALUES (?, ?, ?)"
				stmt, err = db.Prepare(query)
				if err != nil {
					return
				}
				defer stmt.Close()

				_, err = stmt.Exec("2021-01-01 00:00:00", 1.0, 1)
				if err != nil {
					return
				}
				stmt.Close()

				// insert products
				query = "INSERT INTO products (id, description, price) VALUES (?, ?, ?)"
				stmt, err = db.Prepare(query)
				if err != nil {
					return
				}
				defer stmt.Close()

				_, err = stmt.Exec(1, "Product 1", 1.0)
				if err != nil {
					return
				}
				_, err = stmt.Exec(2, "Product 2", 1.0)
				if err != nil {
					return
				}
				stmt.Close()

				// insert sales
				query = "INSERT INTO sales (quantity, product_id, invoice_id) VALUES (?, ?, ?)"
				stmt, err = db.Prepare(query)
				if err != nil {
					return
				}
				defer stmt.Close()

				_, err = stmt.Exec(2, 1, 1)
				if err != nil {
					return
				}
				_, err = stmt.Exec(1, 2, 1)
				if err != nil {
					return
				}
				stmt.Close()

				return
			},
		},
		// valid case - no sells
		{
			name: "valid case - no sells",
			output: output{
				ps: []*ProductSells{},
				err: nil,
				errMsg: "",
			},
			setUpCfgDB: func (cfg *mysql.Config) {
				cfg.User = os.Getenv("DB_MYSQL_USER")
				cfg.Passwd = os.Getenv("DB_MYSQL_PASSWORD")
				cfg.Net = "tcp"
				cfg.Addr = os.Getenv("DB_MYSQL_ADDR")
				cfg.DBName = os.Getenv("DB_MYSQL_DATABASE")
				cfg.ParseTime = true
			},
			setUpDatabase: func (db *sql.DB) (err error) {
				return
			},
		},
	}

	// run tests
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			// arrange
			// -> database - connection
			cfg := mysql.Config{}
			c.setUpCfgDB(&cfg)
			db, err := sql.Open("mysql", cfg.FormatDSN())
			require.NoError(t, err)
			defer db.Close()

			err = db.Ping()
			require.NoError(t, err)

			// -> database - transaction
			tx, err := db.Begin()
			require.NoError(t, err)
			defer tx.Rollback()

			// -> database - set-up
			err = c.setUpDatabase(db)
			require.NoError(t, err)

			// -> storage
			s := NewStorageProductMySQL(db)

			// act
			ps, err := s.TopSelled()

			// assert
			require.Equal(t, c.output.ps, ps)
			require.ErrorIs(t, err, c.output.err)
			if err != nil {
				require.EqualError(t, err, c.output.errMsg)
			}
		})
	}
}