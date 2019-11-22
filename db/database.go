package db

import (
	"fmt"
	"sync"

	_ "github.com/lib/pq"
)

// Database struct that will would our database connection and the options that were used while connecting to this
// database.
type Database struct {
	mu sync.Mutex
	host string
	port string
	username string
	password string
	database string
	sslmode string
}


// ConnectionString returns a string reperesentation of the databaseOptions object and all it's property values,
// this string is used to establish the connection with the database server
func (d *Database) ConnectionString() string {
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


func NewDatabaseConnection(opts ...DatabaseOpts) *Database {
	dbInstance := &Database{}

	for _, opt := range opts {
		opt(dbInstance)
	}

	return dbInstance
}