package db_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/pedrorochaorg/contactsApi/db"
)

func TestNewDatabase(t *testing.T) {

	t.Run("test that we get a blank connection string when sending nil opts", func(t *testing.T) {
		database := db.NewDatabaseConnection()

		assert.Equal(
			t,
			database.GetConnectionString(),
			"host= port= user= password= dbname= sslmode=",
			"Connection strings don't match",
		)

	})

	t.Run("test that we get the right connection string when initiating with options", func(t *testing.T) {
		database := db.NewDatabaseConnection(
			db.WithHost("localhost"),
			db.WithPassword("123456"),
			db.WithPort("5432"),
			db.WithDatabase("contacts_exercise"),
			db.WithUsername("user"),
			db.WithSslMode("disable"),
		)

		assert.Equal(
			t,
			database.GetConnectionString(),
			"host=localhost port=5432 user=user password=123456 dbname=contacts_exercise sslmode=disable",
			"Connection strings don't match",
		)

	})

}
