package repos_test

import (
	"context"
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/pedrorochaorg/contactsApi/db"
	"github.com/pedrorochaorg/contactsApi/obj"
	"github.com/pedrorochaorg/contactsApi/repos"
)

func TestNewUserRepository(t *testing.T) {
	mockDatabase := repos.NewUserRepository(&StubUsersDatabase{
		conOpen: false,
		data:    map[int]obj.User{
			1: {
				ID:        1,
				FirstName: "John",
				LastName:  "Cena",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				Contacts:  nil,
			},
			2: {
				ID:        2,
				FirstName: "Cena",
				LastName:  "Men",
				CreatedAt: time.Time{},
				UpdatedAt: time.Time{},
				Contacts:  nil,
			},
		},
		columns: []string{"id", "firstName", "lastName", "updated_at", "created_at"},
	})

}


// StubUsersDatabase stub database to simulate operations in the database executed by the user repository
type StubUsersDatabase struct {
	conOpen bool
	*sql.DB
	mock sqlmock.Sqlmock
	data map[int]obj.User
	columns []string
}

// OpenConn simulate a database connection opening procedure
func (s *StubUsersDatabase) OpenConn() (bool, error) {
	s.DB, s.mock, _ = sqlmock.New()

	s.conOpen = true

	return true, nil
}

// CloseConn simulate a database connection closing procedure
func (s *StubUsersDatabase) CloseConn() (bool, error) {
	_ = s.Close()
	return true, nil
}


func (s *StubUsersDatabase) FetchAll(ctx context.Context, query string, args ...interface{}) (*db.DatabaseRows, error) {
	listAllUsers := sqlmock.NewRows(s.columns)

	for _, v := range s.data {
		listAllUsers = listAllUsers.AddRow(v.ID, v.FirstName, v.LastName, v.UpdatedAt, v.CreatedAt)
	}
	s.mock.ExpectQuery("SELECT * FROM \"contactsApi\".\"users\"").WillReturnRows(listAllUsers)

	rs, _ := s.QueryContext(ctx, query, args)

	return &db.DatabaseRows{Rows: rs}, nil
}


func (s *StubUsersDatabase) FetchOne(ctx context.Context, query string, args ...interface{}) (*db.DatabaseRows, error) {
	fetchUser := sqlmock.NewRows([]string{"id", "firstName", "lastName", "updated_at", "created_at"})

	var id int

	if id, ok := args[0].(int); ok {
		if user, ok := s.data[id]; ok {
			fetchUser = fetchUser.AddRow(user.ID, user.FirstName, user.LastName, user.UpdatedAt, user.CreatedAt)
		}
	}

	s.mock.ExpectQuery("SELECT * FROM \"contactsApi\".\"users\" WHERE id = $1").
		WithArgs(id).WillReturnRows(fetchUser)

	rs, _ := s.QueryContext(ctx, query, args)

	return &db.DatabaseRows{Rows: rs}, nil
}


func (s *StubUsersDatabase) Delete(ctx context.Context, query string, args ...interface{}) (bool, error) {
	id := args[0].(int)

	result := sqlmock.NewResult(0, 1)

	s.mock.ExpectExec("DELETE FROM \"contactsApi\".\"users\" WHERE id = $1").
		WithArgs(id).WillReturnResult(result)


	_, _ = s.ExecContext(ctx, query, args)

	return true, nil
}

func (s *StubUsersDatabase) Insert(ctx context.Context, query string, args ...interface{}) (*db.DatabaseRows, error) {


	nextId := len(s.data)

	user := obj.User{
		ID:        nextId,
		FirstName: args[0].(string),
		LastName:  args[1].(string),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Contacts:  nil,
	}

	s.data[user.ID] = user

	rows := sqlmock.NewRows(s.columns).
		AddRow(user.ID, user.FirstName, user.LastName, user.UpdatedAt, user.CreatedAt)

	s.mock.ExpectQuery("INSERT INTO \"contactsApi\".\"users\"(\"firstName\", " +
		"\"lastName\") VALUES($1,$2) RETURNING *").WithArgs(args[0], args[1]).WillReturnRows(rows)

	rs, _ := s.QueryContext(ctx, query, args)

	return &db.DatabaseRows{Rows: rs}, nil
}

func (s *StubUsersDatabase) Update(ctx context.Context, query string, args ...interface{}) (*db.DatabaseRows, error) {

	id := args[2].(int)

	user := s.data[id]

	user.FirstName = args[0].(string)
	user.LastName = args[1].(string)
	user.UpdatedAt = time.Now()

	s.data[id] = user

	rows := sqlmock.NewRows(s.columns).
		AddRow(user.ID, user.FirstName, user.LastName, user.UpdatedAt, user.CreatedAt)

	s.mock.ExpectQuery("UPDATE \"contactsApi\".\"users\" SET \"firstName\" = $1, " +
		"\"lastName\" = $2 WHERE id = $3 RETURNING *").
		WithArgs(args[0], args[1], args[2]).WillReturnRows(rows)

	rs, _ := s.QueryContext(ctx, query, args)

	return &db.DatabaseRows{Rows: rs}, nil
}


func (s *StubUsersDatabase) Execute(ctx context.Context, query string) (bool, error) {
	return true, nil
}

func (s *StubUsersDatabase) GetConnectionString() string {
	return ""
}
