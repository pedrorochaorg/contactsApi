package obj_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pedrorochaorg/contactsApi/obj"
)

func TestContact_String(t *testing.T) {
	t.Run("empty object should return string with zero value fields", func(t *testing.T) {
		user := obj.Contact{}

		assert.Equal(t, "ID=0 UserID=0 FirstName= LastName= Email= Phone= CreatedAt=0001-01-01 00:00:00 +0000 UTC" +
			" UpdatedAt=0001-01-01 00:00:00 +0000 UTC", user.String(),
			"Strings don't match")
	})

	t.Run("non empty object should return a string with all field values", func(t *testing.T) {
		parsedTime, _ := time.Parse(time.RFC3339, "2019-11-22T10:00:00Z")

		user := obj.Contact{
			ID: 1,
			UserID: 1,
			FirstName: "Pedro",
			LastName: "Cenas",
			Email: "example@example.com",
			Phone: "919236587",
			CreatedAt: parsedTime,
			UpdatedAt: parsedTime,
		}

		assert.Equal(t, "ID=1 UserID=1 FirstName=Pedro LastName=Cenas Email=example@example." +
			"com Phone=919236587 CreatedAt=2019-11-22 10:00:00 +0000 UTC UpdatedAt=2019-11-22 10:00:00 +0000 UTC",
			user.String(),"Strings don't match")
	})
}
