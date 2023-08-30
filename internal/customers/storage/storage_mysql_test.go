package storage

import (
	"database/sql"
	"os"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/require"
)

// Tests for StorageCustomerMySQL
func TestStorageCustomerMySQL_ConditionInfo(t *testing.T) {
	// test cases
	// type input struct {}
	type output struct { cs []*CustomerConditionInfo; err error; errMsg string }
	type testCase struct {
		name string
		// input input
		output output
		// set-up
		setUpCfgDB func (cfg *mysql.Config)
		setUpDatabase func (db *sql.DB) (err error)
	}

	cases := []testCase{
		// valid case
		{
			name: "valid case",
			output: output{
				cs: []*CustomerConditionInfo{
					{Condition: 0, Total: 1},
					{Condition: 1, Total: 2},
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
				query := "INSERT INTO customers (id, first_name, last_name, `condition`) VALUES (?, ?, ?, ?)"
				stmt, err := db.Prepare(query)
				if err != nil {
					return
				}
				defer stmt.Close()

				_, err = stmt.Exec(1, "John", "Doe", 0)
				if err != nil {
					return
				}
				_, err = stmt.Exec(2, "Jane", "Doe", 1)
				if err != nil {
					return
				}
				_, err = stmt.Exec(3, "John", "Smith", 1)
				if err != nil {
					return
				}

				return
			},
		},
		// valid case - no customers
		{
			name: "valid case - no customers",
			output: output{
				cs: []*CustomerConditionInfo{},
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
			s := NewStorageCustomerMySQL(db)

			// act
			cs, err := s.ConditionInfo()

			// assert
			require.Equal(t, c.output.cs, cs)
			require.ErrorIs(t, err, c.output.err)
			if err != nil {
				require.EqualError(t, err, c.output.errMsg)
			}
		})
	}
}

func TestStorageCustomerMySQL_TopActiveCustomers(t *testing.T) {
	// type input struct {}
	type output struct { cs []*CustomerAmountSpent; err error; errMsg string }
	type testCase struct {
		name string
		// input input
		output output
		// set-up
		setUpCfgDB func (cfg *mysql.Config)
		setUpDatabase func (db *sql.DB) (err error)
	}

	cases := []testCase{
		// valid case
		{
			name: "valid case",
			output: output{
				cs: []*CustomerAmountSpent{
					{FirstName: "John", LastName: "Doe", Amount: 100},
					{FirstName: "Jane", LastName: "Doe", Amount: 200},
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
				query := "INSERT INTO customers (id, first_name, last_name, `condition`) VALUES (?, ?, ?, ?)"
				stmt, err := db.Prepare(query)
				if err != nil {
					return
				}
				defer stmt.Close()

				_, err = stmt.Exec(1, "John", "Doe", 1)
				if err != nil {
					return
				}
				_, err = stmt.Exec(2, "Jane", "Doe", 1)
				if err != nil {
					return
				}

				// insert orders
				query = "INSERT INTO orders (id, customer_id, amount) VALUES (?, ?, ?)"
				stmt, err = db.Prepare(query)
				if err != nil {
					return
				}
				defer stmt.Close()

				_, err = stmt.Exec(1, 1, 100)
				if err != nil {
					return
				}
				_, err = stmt.Exec(2, 2, 200)
				if err != nil {
					return
				}

				return
			},
		},
		// valid case - no customers
		{
			name: "valid case - no customers",
			output: output{
				cs: []*CustomerAmountSpent{},
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
			s := NewStorageCustomerMySQL(db)

			// act
			cs, err := s.TopActiveCustomers()

			// assert
			require.Equal(t, c.output.cs, cs)
			require.ErrorIs(t, err, c.output.err)
			if err != nil {
				require.EqualError(t, err, c.output.errMsg)
			}
		})
	}
}