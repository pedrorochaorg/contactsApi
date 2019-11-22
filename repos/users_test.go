package repos_test

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"

	"github.com/pedrorochaorg/contactsApi/obj"
	"github.com/pedrorochaorg/contactsApi/repos"
)

func TestUserRepository_List(t *testing.T) {

	storedUsers := []obj.User{
		{
			ID:        1,
			FirstName: "John",
			LastName:  "Cena",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Contacts:  nil,
		},
		{
			ID:        2,
			FirstName: "Cena",
			LastName:  "Men",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Contacts:  nil,
		},
	}

	t.Run("test that we are able to obtain a list of stored users", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error while opening a new database connection")
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "firstName", "lastName", "updated_at", "created_at"}).
			AddRow(storedUsers[0].ID, storedUsers[0].FirstName, storedUsers[0].LastName, storedUsers[0].UpdatedAt,
				storedUsers[0].CreatedAt).
			AddRow(storedUsers[1].ID, storedUsers[1].FirstName, storedUsers[1].LastName, storedUsers[1].UpdatedAt,
				storedUsers[1].CreatedAt)

		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		userRepo := repos.NewUserRepository(db)

		users, _ := userRepo.List(context.Background())

		assert.ElementsMatch(t, users, storedUsers, "List's don't match")

	})

	t.Run("test that we are able to handle errors returned by the method", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error while opening a new database connection")
		}
		defer db.Close()

		mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("error"))

		userRepo := repos.NewUserRepository(db)

		_, err = userRepo.List(context.Background())

		assert.Error(t, err, "should have returned an error")
		assert.Contains(t, err.Error(), "error", "Error message doesn't match")

	})

	t.Run("test that we are able to handle errors returned by the method scan call", func(t *testing.T) {
		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error while opening a new database connection")
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "firstName", "lastName", "updated_at", "created_at"}).
			AddRow(nil, storedUsers[0].FirstName, storedUsers[0].LastName, storedUsers[0].UpdatedAt,
				storedUsers[0].CreatedAt)

		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		userRepo := repos.NewUserRepository(db)

		_, err = userRepo.List(context.Background())

		log.Println(err)

		assert.Error(t, err, "should have returned an error")
		assert.Contains(t, err.Error(), "failed to map row to user", "Error message doesn't match")

	})

}

func TestUserRepository_Create(t *testing.T) {


	t.Run("test that we are able to call the method Create", func(t *testing.T) {

		storedUser := obj.User{
			ID:        1,
			FirstName: "John",
			LastName:  "Cena",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Contacts:  nil,
		}

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error while opening a new database connection")
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "firstName", "lastName", "updated_at", "created_at"}).
			AddRow(1, storedUser.FirstName, storedUser.LastName, storedUser.UpdatedAt,
			storedUser.CreatedAt)

		mock.ExpectQuery("INSERT").WillReturnRows(rows)

		userRepo := repos.NewUserRepository(db)

		users, _ := userRepo.Create(context.Background(), &obj.User{FirstName:"John", LastName:"Cena"})

		assert.Equal(t, *users, storedUser, "Results's don't match")

	})

	t.Run("test that we are able to handle a errors returned by the query execution", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error while opening a new database connection")
		}
		defer db.Close()

		mock.ExpectQuery("INSERT").WillReturnError(fmt.Errorf("error"))

		userRepo := repos.NewUserRepository(db)

		_, err = userRepo.Create(context.Background(), &obj.User{FirstName:"John", LastName:"Cena"})

		assert.Error(t, err, "should have returned an error")
		assert.Contains(t, err.Error(), "error", "Error message doesn't match")

	})

	t.Run("test that we are able to handle a error while mapping the query result into an struct", func(t *testing.T) {

		storedUser := obj.User{
			ID:        1,
			FirstName: "John",
			LastName:  "Cena",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Contacts:  nil,
		}

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error while opening a new database connection")
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "firstName", "lastName", "updated_at", "created_at"}).
			AddRow(nil, storedUser.FirstName, storedUser.LastName, storedUser.UpdatedAt,
				storedUser.CreatedAt)

		mock.ExpectQuery("INSERT").WillReturnRows(rows)

		userRepo := repos.NewUserRepository(db)

		_, err = userRepo.Create(context.Background(), &obj.User{FirstName:"John", LastName:"Cena"})

		assert.Error(t, err, "should have returned an error")
		assert.Contains(t, err.Error(), "failed to map row to user", "Error message doesn't match")

	})

}

func TestUserRepository_Update(t *testing.T) {



	t.Run("test that we are able to call the method Update", func(t *testing.T) {

		storedUser := obj.User{
			ID:        1,
			FirstName: "John",
			LastName:  "Cena",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Contacts:  nil,
		}

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error while opening a new database connection")
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "firstName", "lastName", "updated_at", "created_at"}).
			AddRow(1, storedUser.FirstName, storedUser.LastName, storedUser.UpdatedAt,
				storedUser.CreatedAt)

		mock.ExpectQuery("UPDATE").WillReturnRows(rows)

		userRepo := repos.NewUserRepository(db)

		users, _ := userRepo.Update(context.Background(), &obj.User{ID: 1, FirstName:"John", LastName:"Cena"})

		assert.Equal(t, *users, storedUser, "Results's don't match")

	})

	t.Run("test that we are able to handle a errors returned by the query execution", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error while opening a new database connection")
		}
		defer db.Close()

		mock.ExpectQuery("UPDATE").WillReturnError(fmt.Errorf("error"))

		userRepo := repos.NewUserRepository(db)

		_, err = userRepo.Update(context.Background(), &obj.User{ID: 1, FirstName:"John", LastName:"Cena"})

		assert.Error(t, err, "should have returned an error")
		assert.Contains(t, err.Error(), "error", "Error message doesn't match")

	})

	t.Run("test that we are able to handle a error while mapping the query result into an struct", func(t *testing.T) {

			storedUser := obj.User{

				ID:        1,
				FirstName: "John",
				LastName:  "Cena",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Contacts:  nil,
			}

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error while opening a new database connection")
			}
			defer db.Close()

			rows := sqlmock.NewRows([]string{"id", "firstName", "lastName", "updated_at", "created_at"}).
				AddRow(nil, storedUser.FirstName, storedUser.LastName, storedUser.UpdatedAt,
					storedUser.CreatedAt)

			mock.ExpectQuery("UPDATE").WillReturnRows(rows)

			userRepo := repos.NewUserRepository(db)

			_, err = userRepo.Update(context.Background(), &obj.User{ID: 1, FirstName:"John", LastName:"Cena"})

			assert.Error(t, err, "should have returned an error")
			assert.Contains(t, err.Error(), "failed to map row to user", "Error message doesn't match")

		})

}

func TestUserRepository_Get(t *testing.T) {



	t.Run("test that we are able to call the method Get", func(t *testing.T) {

		storedUser := obj.User{
			ID:        1,
			FirstName: "John",
			LastName:  "Cena",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Contacts:  nil,
		}

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error while opening a new database connection")
		}
		defer db.Close()

		rows := sqlmock.NewRows([]string{"id", "firstName", "lastName", "updated_at", "created_at"}).
			AddRow(1, storedUser.FirstName, storedUser.LastName, storedUser.UpdatedAt,
				storedUser.CreatedAt)


		mock.ExpectQuery("SELECT").WillReturnRows(rows)

		userRepo := repos.NewUserRepository(db)

		users, _ := userRepo.Get(context.Background(), 1)

		assert.Equal(t, *users, storedUser, "Results's don't match")

	})

	t.Run("test that we are able to handle a error while executing the update query", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error while opening a new database connection")
		}
		defer db.Close()

		mock.ExpectQuery("SELECT").WillReturnError(fmt.Errorf("error"))

		userRepo := repos.NewUserRepository(db)

		_, err = userRepo.Get(context.Background(), 1)

		assert.Error(t, err, "should have returned an error")
		assert.Contains(t, err.Error(), "error", "Error message doesn't match")

	})

	t.Run("test that we are able to handle a error while mapping the query result into an struct", func(t *testing.T) {

			storedUser := obj.User{
				ID:        1,
				FirstName: "John",
				LastName:  "Cena",
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Contacts:  nil,
			}

			db, mock, err := sqlmock.New()
			if err != nil {
				t.Fatalf("error while opening a new database connection")
			}
			defer db.Close()

			rows := sqlmock.NewRows([]string{"id", "firstName", "lastName", "updated_at", "created_at"}).
				AddRow(nil, storedUser.FirstName, storedUser.LastName, storedUser.UpdatedAt,
					storedUser.CreatedAt)


			mock.ExpectQuery("SELECT").WillReturnRows(rows)

			userRepo := repos.NewUserRepository(db)

			_, err = userRepo.Get(context.Background(), 1)

			assert.Error(t, err, "should have returned an error")
			assert.Contains(t, err.Error(), "failed to map row to user", "Error message doesn't match")

		})

}



func TestUserRepository_Delete(t *testing.T) {



	t.Run("call the method Delete", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error while opening a new database connection")
		}
		defer db.Close()

		result := sqlmock.NewResult(0, 0)


		mock.ExpectExec("DELETE").WithArgs(1).WillReturnResult(result)

		userRepo := repos.NewUserRepository(db)

		users, _ := userRepo.Delete(context.Background(), 1)

		assert.True(t, users, "Shoudl have returned a value of true")

	})

	t.Run("handle a error while executing the delete query", func(t *testing.T) {

		db, mock, err := sqlmock.New()
		if err != nil {
			t.Fatalf("error while opening a new database connection")
		}
		defer db.Close()

		mock.ExpectExec("DELETE").WithArgs(1).WillReturnError(fmt.Errorf("error"))

		userRepo := repos.NewUserRepository(db)

		_, err = userRepo.Delete(context.Background(), 1)

		assert.Error(t, err, "should have returned an error")
		assert.Contains(t, err.Error(), "error", "Error message doesn't match")

	})


}
