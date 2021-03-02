package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/KushagraMehta/Basic-PostgreSQL-Go-connection/model"
	"github.com/jackc/pgx/v4"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
		os.Exit(1)
	}

	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"),
		os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGDATABASE"))

	model.DB, err = pgx.Connect(context.Background(), databaseURL)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer model.DB.Close(context.Background())

	fmt.Println(model.HelloWorld())
}
