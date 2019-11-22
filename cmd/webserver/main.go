package main

import (
	"database/sql"
	"log"
	"net/http"

	_ "github.com/lib/pq"

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
	db, err := sql.Open("postgres", database.ConnectionString())
	if err != nil {
		log.Fatalf("error starting database connection: %s", err)
	}

	defer db.Close()

	server := api.NewAPI(db)

	log.Println("Starting the webserver in port 3000")
	if err := http.ListenAndServe(":3000", server); err != nil {
		log.Fatalf("Error while starting the web server: %s", err)
	}

}