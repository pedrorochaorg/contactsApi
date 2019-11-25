package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/pedrorochaorg/contactsApi/obj"
)

func TestNewUserHandler(t *testing.T) {


	parsedTime, _ := time.Parse(time.RFC3339, "2019-11-22T10:00:00Z")

	userList := []obj.User{
		obj.User{1, "John", "Cena", parsedTime, parsedTime, nil},
		obj.User{2, "João", "Cenas", parsedTime, parsedTime, nil},
		obj.User{3, "Pedro", "Costas", parsedTime, parsedTime, nil},
	}

	userHandler := NewUserHandler(
		&StubUserRepo{users: userList},
	)

	t.Run("fetch a user by id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d", 2), nil)
		response := httptest.NewRecorder()

		userHandler.ServeHTTP(response, req)

		assert.Equal(t, http.StatusOK, response.Code, "Status Code doesn't match")




		assert.Equal(t, userList[1], fetchedUser)
	})


}

type StubUserRepo struct {
	sync.Mutex
	users []obj.User
}

func (s *StubUserRepo) List(ctx context.Context) ([]obj.User, error) {
	return s.users, nil
}

func (s *StubUserRepo) Create(ctx context.Context, user *obj.User) (*obj.User, error) {
	newUser := obj.User{
		ID:        len(s.users),
		FirstName: user.FirstName,
		LastName:  user.LastName,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Contacts:  nil,
	}

	s.users = append(s.users, newUser)

	return &newUser, nil
}

func (s *StubUserRepo) Update(ctx context.Context, user *obj.User) (*obj.User, error) {
	var updatedUser *obj.User = nil

	for _, v := range s.users {
		if v.ID == user.ID {
			updatedUser = &v
			break
		}
	}

	if updatedUser == nil {
		return nil, fmt.Errorf("User not found %d", user.ID)
	}

	updatedUser.FirstName = user.FirstName
	updatedUser.LastName = user.LastName

	updatedUser.UpdatedAt = time.Now()

	return updatedUser, nil

}

func (s *StubUserRepo) Get(ctx context.Context, id int) (*obj.User, error) {
	var fetchedUser *obj.User = nil

	log.Println("I'm in here")

	for _, v := range s.users {
		log.Println("UserId:", v.ID)
		if v.ID == id {

			fetchedUser = &v
			break
		}
	}

	if fetchedUser == nil {
		return nil, fmt.Errorf("User not found %d", id)
	}

	return fetchedUser, nil
}

func (s *StubUserRepo) Delete(ctx context.Context, id int) (bool, error)  {
	var fetchedUser *obj.User = nil
	var fetchedUserIndex int = -1

	for i, v := range s.users {
		if v.ID == id {
			fetchedUser = &v
			fetchedUserIndex = i
			break
		}
	}

	if fetchedUser == nil {
		return false, fmt.Errorf("User not found %d", id)
	}

	s.users = append(s.users[:fetchedUserIndex], s.users[fetchedUserIndex+1:]...)

	return true, nil

}


func getUserFromResponse(response string) obj.User {
	responseData := Response{}

	err := json.Unmarshal([]byte(response), responseData)

	json.NewDecoder(bytes).Decode(&responseData)

	buffer := bytes.NewBuffer(nil)

	json.NewEncoder(buffer).Encode(responseData.Result)

	fetchedUser := obj.User{}

	json.NewDecoder(buffer).Decode(&fetchedUser)
}