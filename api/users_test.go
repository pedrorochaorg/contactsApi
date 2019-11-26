package api

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
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

		user, err := getUserFromResponse(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusOK, response.Code, "Status Code doesn't match")
		assert.Equal(t, userList[1], *user)
	})

	t.Run("fetch an unexisting user by id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d", 4), nil)
		response := httptest.NewRecorder()

		userHandler.ServeHTTP(response, req)

		message, err := getResponseMessage(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusNotFound, response.Code, "Status Code doesn't match")
		assert.Equal(t, UserNotFound, message)
	})

	t.Run("fetch an user using an invalid id format", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodGet, "/users/invalid", nil)
		response := httptest.NewRecorder()

		userHandler.ServeHTTP(response, req)

		message, err := getResponseMessage(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusBadRequest, response.Code, "Status Code doesn't match")
		assert.Equal(t, BadIdFormat, message)
	})


	t.Run("delete a user by id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%d", 2), nil)
		response := httptest.NewRecorder()

		userHandler.ServeHTTP(response, req)

		message, err := getResponseMessage(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusNoContent, response.Code, "Status Code doesn't match")
		assert.Equal(t, UserDeletedSuccessfully, message)

		// Check that the user was really deleted
		req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d", 2), nil)
		response = httptest.NewRecorder()

		userHandler.ServeHTTP(response, req)

		message, err = getResponseMessage(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusNotFound, response.Code, "Status Code doesn't match")
		assert.Equal(t, UserNotFound, message)

	})

	t.Run("delete an unexisting user by id", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, fmt.Sprintf("/users/%d", 4), nil)
		response := httptest.NewRecorder()

		userHandler.ServeHTTP(response, req)

		message, err := getResponseMessage(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusNotFound, response.Code, "Status Code doesn't match")
		assert.Equal(t, UserNotFound, message)
	})

	t.Run("delete an user using an invalid id format", func(t *testing.T) {
		req, _ := http.NewRequest(http.MethodDelete, "/users/invalid", nil)
		response := httptest.NewRecorder()

		userHandler.ServeHTTP(response, req)

		message, err := getResponseMessage(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusBadRequest, response.Code, "Status Code doesn't match")
		assert.Equal(t, BadIdFormat, message)
	})


	t.Run("create a new user", func(t *testing.T) {

		bodyData, err := json.Marshal(obj.User{
			FirstName: "José",
			LastName:  "Santos",
		})
		if err != nil {
			t.Fatalf("error marshling request body %s", err)
		}

		req, _ := http.NewRequest(
			http.MethodPost,
			"/users/",
			bytes.NewBuffer(bodyData),
		)
		response := httptest.NewRecorder()


		var nextId = getNextId(userList)


		userHandler.ServeHTTP(response, req)



		user, err := getUserFromResponse(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusCreated, response.Code, "Status Code doesn't match")
		assert.Equal(t, "José", user.FirstName)
		assert.Equal(t, "Santos", user.LastName)
		assert.Equal(t, nextId, user.ID)



		req, _ = http.NewRequest(http.MethodGet, fmt.Sprintf("/users/%d", user.ID), nil)
		response = httptest.NewRecorder()

		userHandler.ServeHTTP(response, req)

		storedUser, err := getUserFromResponse(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusOK, response.Code, "Status Code doesn't match")
		assert.Equal(t, *user, *storedUser)

	})

	t.Run("create a new user with an invalid request body", func(t *testing.T) {


		req, _ := http.NewRequest(
			http.MethodPost,
			"/users/",
			bytes.NewBuffer([]byte("some string")),
		)
		response := httptest.NewRecorder()


		userHandler.ServeHTTP(response, req)

		message, err := getResponseMessage(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusBadRequest, response.Code, "Status Code doesn't match")
		assert.Contains(t, message, "invalid character", "Message should contain the string 'invalid character'")

	})


	t.Run("update an existing user", func(t *testing.T) {

		originalUser := userList[0]

		bodyData, err := json.Marshal(obj.User{
			FirstName: "Mário",
			LastName:  "Figueira",
		})
		if err != nil {
			t.Fatalf("error marshling request body %s", err)
		}

		req, _ := http.NewRequest(
			http.MethodPut,
			"/users/1",
			bytes.NewBuffer(bodyData),
		)
		response := httptest.NewRecorder()

		userHandler.ServeHTTP(response, req)

		user, err := getUserFromResponse(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusAccepted, response.Code, "Status Code doesn't match")
		assert.Equal(t, "Mário", user.FirstName)
		assert.Equal(t, "Figueira", user.LastName)
		assert.NotEqual(t, originalUser.FirstName, user.FirstName)
		assert.NotEqual(t, originalUser.LastName, user.LastName)
		assert.Equal(t, originalUser.ID, user.ID)

	})

	t.Run( "update an unexisting user", func(t *testing.T) {
		bodyData, err := json.Marshal(obj.User{
			FirstName: "Mário",
			LastName:  "Figueira",
		})
		if err != nil {
			t.Fatalf("error marshling request body %s", err)
		}

		req, _ := http.NewRequest(
			http.MethodPut,
			"/users/52",
			bytes.NewBuffer(bodyData),
		)
		response := httptest.NewRecorder()

		userHandler.ServeHTTP(response, req)

		message, err := getResponseMessage(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusNotFound, response.Code, "Status Code doesn't match")
		assert.Equal(t, UserNotFound, message)
	})

	t.Run( "update an user using an invalid id format", func(t *testing.T) {
		bodyData, err := json.Marshal(obj.User{
			FirstName: "Mário",
			LastName:  "Figueira",
		})
		if err != nil {
			t.Fatalf("error marshling request body %s", err)
		}

		req, _ := http.NewRequest(
			http.MethodPut,
			"/users/invalid",
			bytes.NewBuffer(bodyData),
		)
		response := httptest.NewRecorder()

		userHandler.ServeHTTP(response, req)

		message, err := getResponseMessage(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusBadRequest, response.Code, "Status Code doesn't match")
		assert.Equal(t, BadIdFormat, message)
	})

	t.Run( "update an user with an invalid request body", func(t *testing.T) {

		req, _ := http.NewRequest(
			http.MethodPut,
			"/users/1",
			bytes.NewBuffer([]byte("something not usefull")),
		)
		response := httptest.NewRecorder()

		userHandler.ServeHTTP(response, req)

		message, err := getResponseMessage(response.Body)

		if err != nil {
			t.Fatalf("error while unmarshling the response body %s", err)
		}

		assert.Equal(t, http.StatusBadRequest, response.Code, "Status Code doesn't match")
		assert.Contains(t, message, "invalid character", "Message should contain the string 'invalid character'")
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
		ID:        getNextId(s.users),
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

	for _, v := range s.users {
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


func getNextId(users []obj.User) int {
	var highestIndex = 0

	for _, v := range users {
		if v.ID > highestIndex {
			highestIndex = v.ID
		}
	}

	return highestIndex + 1

}


func getUserFromResponse(response *bytes.Buffer) (*obj.User, error) {
	responseObject := Response{}

	err := json.NewDecoder(response).Decode(&responseObject)
	if err != nil {
		return nil, err
	}

	responseData, err := json.Marshal(responseObject.Result)
	if err != nil {
		return nil, err
	}

	user := obj.User{}
	err = json.Unmarshal(responseData, &user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}



func getResponseMessage(response *bytes.Buffer) (string, error) {
	responseObject := Response{}

	err := json.NewDecoder(response).Decode(&responseObject)
	if err != nil {
		return "", err
	}

	return responseObject.Message, nil


}