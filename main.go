package main

import (
	"context"
	"log"
	"os"

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
	_, err = database.NewMongoClient(context.Background(), uri, dbName)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("successfully initialized and pinged deployment")
}
