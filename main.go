package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/inflame-ue/godo/internal/api"
	"github.com/inflame-ue/godo/internal/database"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("loading .env file: %v", err)
	}

	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_NAME")
	db, err := database.NewMongoClient(context.Background(), uri, dbName)
	if err != nil {
		log.Fatal(err)
	}

	port := os.Getenv("PORT")
	srv := api.NewAPI(db)
	serv := http.Server{
		Addr:              ":" + port,
		Handler:           srv,
		ReadHeaderTimeout: 30 * time.Second,
		WriteTimeout:      30 * time.Second,
	}
	log.Printf("server listening on port %s", port)
	log.Fatal(serv.ListenAndServe())
}
