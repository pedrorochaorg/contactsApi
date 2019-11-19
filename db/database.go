package db

import (
	"context"
	"database/sql"
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

// Database struct that will would our database connection and the options that were used while connecting to this
// database.
type Database struct {
	*sql.DB
	mu sync.Mutex
	host string
	port string
	username string
	password string
	database string
	sslmode string
}

type DatabaseRows struct {
	*sql.Rows
}


// ConnectionString returns a string reperesentation of the databaseOptions object and all it's property values,
// this string is used to establish the connection with the database server
func (d *Database) connectionString() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", d.host, d.port, d.username,
		d.password, d.database, d.sslmode)
}

// DatabaseOpts type func used to populate the databaseoptions struct with each property value implementing the
// Functional Options pattern
type DatabaseOpts func(d *Database)

// WithHost set's the value of the 'host' property value of the 'databaseOptions' struct
func WithHost(host string) DatabaseOpts {
	return func(d *Database) {
		d.host = host
	}
}

// WithPort set's the value of the 'host' property value of the 'databaseOptions' struct
func WithPort(port string) DatabaseOpts {
	return func(d *Database) {
		d.port = port
	}
}

// WithUsername set's the value of the 'username' property value of the 'databaseOptions' struct
func WithUsername(username string) DatabaseOpts {
	return func(d *Database) {
		d.username = username
	}
}

// WithPassword set's the value of the 'password' property value of the 'databaseOptions' struct
func WithPassword(password string) DatabaseOpts {
	return func(d *Database) {
		d.password = password
	}
}

// WithDatabase set's the value of the 'database' property value of the 'databaseOptions' struct
func WithDatabase(database string) DatabaseOpts {
	return func(d *Database) {
		d.database = database
	}
}

// WithSslMode set's the value of the 'sslmode' property value of the 'databaseOptions' struct
func WithSslMode(sslmode string) DatabaseOpts {
	return func(d *Database) {
		d.sslmode = sslmode
	}
}

// DatabaseConnection interface that defines a new database connection and the methods that it should implement
type DatabaseConnection interface {
	OpenConn() (bool, error)
	CloseConn() (bool, error)
	Query(ctx context.Context, query string) (*DatabaseRows, error)
	QueryWithArgs(ctx context.Context, query string, args ...interface{}) (*DatabaseRows, error)
	Delete(ctx context.Context, query string, args ...interface{}) (bool, error)
	Execute(ctx context.Context, query string) (bool, error)
	GetConnectionString() string
}

// OpenConn opens the database connection
func (d *Database) OpenConn() (bool, error) {
	d.mu.Lock()

	db, err := sql.Open("postgres", d.GetConnectionString())
	if err != nil {
		return false, fmt.Errorf("an error ocurred while opening the database connection: %s", err)
	}

	d.DB = db
	return true, nil
}


// CloseConn closes the existing database connection
func (d *Database) CloseConn() (bool, error) {

	if err := d.Close(); err != nil {
		return false, fmt.Errorf("an error ocurred while closing the database connection: %s", err)
	}

	d.mu.Unlock()
	return true, nil
}

// Query
func (d *Database) Query(ctx context.Context, query string) (*DatabaseRows, error) {

	rows, err := d.QueryContext(ctx, query)
	if err != nil {
		return nil, fmt.Errorf("an error ocurred while executing the query: %s", err)
	}

	return &DatabaseRows{rows}, nil

}

func (d *Database) QueryWithArgs(ctx context.Context, query string, args ...interface{}) (*DatabaseRows, error) {

	rows, err := d.QueryContext(ctx, query, args ...)
	if err != nil {
		return nil, fmt.Errorf("an error ocurred while executing the query: %s", err)
	}

	return &DatabaseRows{rows}, nil

}


func (d *Database) Delete(ctx context.Context, query string, args ...interface{}) (bool, error) {
	result, err := d.Exec(query, args...)
	if err != nil {
		return false, fmt.Errorf("an error ocurred while executing the query: %s", err)
	}

	_, err = result.RowsAffected()
	if err != nil {
		return false, fmt.Errorf("an error ocurred while obtaining the number of affected rows: %s", err)
	}

	return true, nil
}

func (d *Database) Execute(ctx context.Context, query string) (bool, error) {
	_, err := d.Exec(query)
	if err != nil {
		return false, fmt.Errorf("an error ocurred while executing the query: %s , %s", query, err)
	}

	return true, nil
}

func (d *Database) GetConnectionString() string {
	return d.connectionString()
}

func NewDatabaseConnection(opts ...DatabaseOpts) DatabaseConnection {
	dbInstance := &Database{}
	for _, opt := range opts {
		opt(dbInstance)
	}

	return dbInstance
}