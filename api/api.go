package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/pedrorochaorg/contactsApi/db"
	"github.com/pedrorochaorg/contactsApi/repos"
)

const (
	ContentReady     = string("Content Ready")
	ErrNotFound     = string("Page not found")
	JsonContentType = "application/json"
)

type API struct {
	db db.DatabaseConnection
	http.Handler
}

func NewAPI(db db.DatabaseConnection) *API {
	handler := new(API)

	handler.db = db

	initDB(db)

	router := http.NewServeMux()

	repository := repos.NewUserRepository(db)
	router.Handle("/users/", NewUserHandler(&repository))


	handler.Handler = router
	return handler
}

func initDB(database db.DatabaseConnection) {

	log.Println("Initializing database")
	for _, stmt := range db.InitStatements {
		_, err := database.Execute(context.Background(), stmt)
		if err != nil {
			log.Fatalf("failed to execute statement in database: %s, %s", stmt, err)
		}
	}

}


// FailureReply method that encodes a notFoundReply to
func FailureReply(er *Error ,w http.ResponseWriter, r *http.Request) {
	log.Printf("Path: %s, Method: %s, Msg: %s, Status: %d", r.URL.Path, r.Method, er.msg, er.status)

	w.Header().Set("content-type", JsonContentType)
	w.WriteHeader(er.status)

	response := Response{
		Status:  false,
		Message: er.msg,
		Result:  nil,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Failed to encode object %v", response)
		return
	}
}


func SuccessReply(data *Data, w http.ResponseWriter, r *http.Request) {
	log.Printf("Path: %s, Method: %s, Msg: %s, Status: %d", r.URL.Path, r.Method, data.message, data.status)

	w.Header().Set("content-type", JsonContentType)
	w.WriteHeader(data.status)

	response := Response{
		Status:  true,
		Message: data.message,
		Result:  data.data,
	}

	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Failed to encode object %v", response)
		return
	}
}
