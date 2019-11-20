package repos

import (
	"context"
	"fmt"

	"github.com/pedrorochaorg/contactsApi/db"
	"github.com/pedrorochaorg/contactsApi/obj"
)


type UserRepo interface {
	List(ctx context.Context) ([]obj.User, error)
	Create(ctx context.Context, user *obj.User) (*obj.User, error)
	Update(ctx context.Context, user *obj.User) (*obj.User, error)
	Get(ctx context.Context, id int64) (*obj.User, error)
	Delete(ctx context.Context, id int64) (bool, error)
}


type UserRepository struct {
	db db.DatabaseConnection
}

// NewUserRepository instantiates a new user repository injecting the database connection interface as a dependency
func NewUserRepository(db db.DatabaseConnection) UserRepository {
	return UserRepository{db}
}

// List return a set of users from database
func (u *UserRepository) List(ctx context.Context) ([]obj.User, error) {
	rows, err := u.db.FetchAll(ctx, "SELECT * FROM \"contactsApi\".\"users\"")
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users from database: %s", err)
	}

	users := []obj.User{}

	defer rows.Close()
	for rows.Next() {
		user := obj.User{}
		err = rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.UpdatedAt,
			&user.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to map row to user: %s", err)
		}
		users = append(users, user)
	}
	return users, nil
}

// Creates a user in database
func (u *UserRepository) Create(ctx context.Context, user *obj.User) (*obj.User, error) {
	rows, err := u.db.Insert(ctx, "INSERT INTO \"contactsApi\".\"users\"(\"firstName\", " +
		"\"lastName\") VALUES($1," +
		"$2) RETURNING *",
		user.FirstName, user.LastName)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users from database: %s", err)
	}

	defer rows.Close()
	rows.Next()
	err = rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UpdatedAt,
		&user.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to map row to user: %s", err)
	}

	return user, nil
}

// Update
func (u *UserRepository) Update(ctx context.Context, user *obj.User) (*obj.User, error) {
	rows, err := u.db.Update(ctx, "UPDATE \"contactsApi\".\"users\" SET \"firstName\" = $1, " +
		"\"lastName\" = $2 WHERE id = $3 RETURNING *",
		user.FirstName, user.LastName, user.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users from database: %s", err)
	}

	defer rows.Close()
	rows.Next()
	err = rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UpdatedAt,
		&user.CreatedAt)


	if err != nil {
		return nil, fmt.Errorf("failed to map row to user: %s", err)
	}

	return user, nil
}

// Get
func (u *UserRepository) Get(ctx context.Context, id int64) (*obj.User, error) {
	rows, err := u.db.FetchOne(ctx, "SELECT * FROM \"contactsApi\".\"users\" WHERE id = $1", id)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch users from database: %s", err)
	}

	user := &obj.User{}

	defer rows.Close()
	rows.Next()
	err = rows.Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.UpdatedAt,
		&user.CreatedAt)

	if err != nil {
		return nil, fmt.Errorf("failed to map row to user: %s", err)
	}

	return user, nil
}

// Delete
func (u *UserRepository) Delete(ctx context.Context, id int64) (bool, error) {
	rows, err := u.db.Delete(ctx, "DELETE FROM \"contactsApi\".\"users\" WHERE id = $1", id)
	if err != nil {
		return false, fmt.Errorf("failed to fetch users from database: %s", err)
	}

	if !rows {
		return false, fmt.Errorf("failed to delete user: %d", id)
	}

	return true, nil
}

