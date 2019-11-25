package api

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/pedrorochaorg/contactsApi/obj"
	"github.com/pedrorochaorg/contactsApi/repos"
)

const (
	UserCreatedSuccessfully  = "User successfully created!"
	UserUpdatedSuccessfully  = "User successfully updated!"
	UserDeletedSuccessfully  = "User successfully deleted!"
	BadIdFormat  = "Bad id format!"
	UserNotFound  = "User not found!"
)

type UserHandler struct {
	repo repos.UserRepo
	*http.ServeMux
	handlers Handlers
}

func NewUserHandler(db repos.UserRepo) *UserHandler {
	handler := new(UserHandler)

	handler.repo = db

	handler.handlers = Handlers{}

	handler.handlers.Add("", http.MethodGet, handler.listUsers)
	handler.handlers.Add("", http.MethodPost, handler.createUser)
	handler.handlers.Add("/{id}", http.MethodGet, handler.getUser)
	handler.handlers.Add("/{id}", http.MethodPut, handler.updateUser)
	handler.handlers.Add("/{id}", http.MethodDelete, handler.deleteUser)

	return handler
}

func (u *UserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	subPath := r.URL.Path[len("/users/"):]

	if subPath[len(subPath)-1:] == "/" {
		subPath = subPath[:len(subPath)-1]
	}

	httpMethod := r.Method

	handler, err := u.handlers.GetByMethodAndType(subPath, httpMethod)

	if err != nil {
		FailureReply(err, w, r)
		return
	}

	handler.H.handler(w, UrlRequest{R: r, Vars: handler.Vars})

}

func (u *UserHandler) listUsers(w http.ResponseWriter, r UrlRequest) {

	users, err := u.repo.List(r.R.Context())
	if err != nil {
		FailureReply(&Error{msg: err.Error(), status: 500}, w, r.R)
		return
	}

	SuccessReply(
		&Data{status: 200, message: ContentReady, data: users},
		w,
		r.R,
	)

}

func (u *UserHandler) createUser(w http.ResponseWriter, r UrlRequest) {
	user := obj.User{}
	err := json.NewDecoder(r.R.Body).Decode(&user)
	if err != nil {
		FailureReply(&Error{msg: err.Error(), status: 500}, w, r.R)
		return
	}

	finalUser, err := u.repo.Create(r.R.Context(), &user)
	if err != nil {
		FailureReply(&Error{msg: err.Error(), status: 500}, w, r.R)
		return
	}

	SuccessReply(
		&Data{status: http.StatusCreated, message: UserCreatedSuccessfully, data: finalUser},
		w,
		r.R,
	)

}

func (u *UserHandler) getUser(w http.ResponseWriter, r UrlRequest) {

	userId, err := strconv.Atoi(r.Vars["id"])
	if err != nil {
			FailureReply(&Error{msg: BadIdFormat, status: 400}, w, r.R)
			return
	}

	user, err := u.repo.Get(r.R.Context(), userId)
	if err != nil {
		FailureReply(&Error{msg: UserNotFound, status: 404}, w, r.R)
		return
	}

	SuccessReply(
		&Data{status: http.StatusOK, message: ContentReady, data: user},
		w,
		r.R,
	)
}


func (u *UserHandler) updateUser(w http.ResponseWriter, r UrlRequest) {

	userId, err := strconv.Atoi(r.Vars["id"])
	if err != nil {
		FailureReply(&Error{msg: BadIdFormat, status: 400}, w, r.R)
		return
	}

	user, err := u.repo.Get(r.R.Context(), int(userId))
	if err != nil {
		FailureReply(&Error{msg: UserNotFound, status: 404}, w, r.R)
		return
	}

	updatedUser := obj.User{}
	err = json.NewDecoder(r.R.Body).Decode(&updatedUser)
	if err != nil {
		FailureReply(&Error{msg: err.Error(), status: 500}, w, r.R)
		return
	}

	user.FirstName = updatedUser.FirstName
	user.LastName = updatedUser.LastName

	finalUser, err := u.repo.Update(r.R.Context(), user)
	if err != nil {
		FailureReply(&Error{msg: err.Error(), status: 500}, w, r.R)
		return
	}

	SuccessReply(
		&Data{status: http.StatusAccepted, message: UserUpdatedSuccessfully, data: finalUser},
		w,
		r.R,
	)

}


func (u *UserHandler) deleteUser(w http.ResponseWriter, r UrlRequest) {

	userId, err := strconv.Atoi(r.Vars["id"])
	if err != nil {
		FailureReply(&Error{msg: BadIdFormat, status: 400}, w, r.R)
		return
	}

	_, err = u.repo.Get(r.R.Context(), userId)
	if err != nil {
		FailureReply(&Error{msg: UserNotFound, status: 404}, w, r.R)
		return
	}

	_, err = u.repo.Delete(r.R.Context(), userId)
	if err != nil {
		FailureReply(&Error{msg: err.Error(), status: 400}, w, r.R)
		return
	}

	SuccessReply(
		&Data{status: http.StatusOK, message: UserDeletedSuccessfully, data: nil},
		w,
		r.R,
	)

}


func (u *UserHandler) getUserContacts(w http.ResponseWriter, r UrlRequest) {

	userId, err := strconv.Atoi(r.Vars["userId"])
	if err != nil {
		FailureReply(&Error{msg: BadIdFormat, status: 400}, w, r.R)
		return
	}

	contactId, err := strconv.Atoi(r.Vars["id"])
	if err != nil {
		FailureReply(&Error{msg: BadIdFormat, status: 400}, w, r.R)
		return
	}

	log.Println("ContactId", contactId)

	user, err := u.repo.Get(r.R.Context(), userId)
	if err != nil {
		FailureReply(&Error{msg: UserNotFound, status: 404}, w, r.R)
		return
	}

	SuccessReply(
		&Data{status: http.StatusOK, message: ContentReady, data: user},
		w,
		r.R,
	)
}

