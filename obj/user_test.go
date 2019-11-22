package obj_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pedrorochaorg/contactsApi/obj"
)

func TestUser_String(t *testing.T) {
	t.Run("empty object should return string with zero value fields", func(t *testing.T) {
		user := obj.User{}

		assert.Equal(t, "ID=0 FirstName= LastName= CreatedAt=0001-01-01 00:00:00 +0000 UTC UpdatedAt=0001-01-01 00:00" +
			":00 +0000 UTC Contacts=[]", user.String(),
			"Strings don't match")
	})

	t.Run("empty object should return string with zero value fields", func(t *testing.T) {
		parsedTime, _ := time.Parse(time.RFC3339, "2019-11-22T10:00:00Z")

		user := obj.User{
			ID: 1,
			FirstName: "Pedro",
			LastName: "Cenas",
			CreatedAt: parsedTime,
			UpdatedAt: parsedTime,
		}

		assert.Equal(t, "ID=1 FirstName=Pedro LastName=Cenas CreatedAt=2019-11-22 10:00:00 +0000 UTC UpdatedAt=2019" +
			"-11-22 10:00:00 +0000 UTC Contacts=[]", user.String(),
			"Strings don't match")
	})
}
