package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

func init() {
	_, err := pgx.Connect(context.Background(), "postgres://tester:alpha-omega@localhost:5432/test_task")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
	} else {
		fmt.Println("Connected to PSQL!")
	}
}
