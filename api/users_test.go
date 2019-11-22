package api

import (
	"context"
	"testing"

	"github.com/pedrorochaorg/contactsApi/obj"
)

func TestNewUserHandler(t *testing.T) {

	userHandler := NewUserHandler()


}

type StubUserRepo struct {
	users []obj.User
}

func (s *StubUserRepo) List(ctx context.Context) ([]obj.User, error) {

}

func (s *StubUserRepo) Create(ctx context.Context, user *obj.User) (*obj.User, error) {

}

func (s *StubUserRepo) Update(ctx context.Context, user *obj.User) (*obj.User, error) {

}

func (s *StubUserRepo) Get(ctx context.Context, id int) (*obj.User, error) {

}

func (s *StubUserRepo) Delete(ctx context.Context, id int) (bool, error)  {

}