package model

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v4"
)

var DB *pgx.Conn

func HelloWorld() string {

	var greeting string
	err := DB.QueryRow(context.Background(), "select 'Hello, world!'").Scan(&greeting)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	return greeting
}
