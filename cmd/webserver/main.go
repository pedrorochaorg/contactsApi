package main

import (
	"log"
	"net/http"

	"github.com/pedrorochaorg/contactsApi/api"
	"github.com/pedrorochaorg/contactsApi/db"
)

func main() {
	database := db.NewDatabaseConnection(
		db.WithUsername("contacts"),
		db.WithSslMode("disable"),
		db.WithDatabase("contacts"),
		db.WithHost("localhost"),
		db.WithPort("5435"),
		db.WithPassword("TwE5]>*Gm^sk_eq)"),
	)

	if ok, err := database.OpenConn(); !ok || err !=nil {
		log.Fatalf("Error while opening a connection to the database: %s", err)
		return
	}

	defer database.CloseConn()

	server := api.NewAPI(database)

	log.Println("Starting the webserver in port 3000")
	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatalf("Error while starting the web server: %s", err)
	}

}