package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"os"
)

var Connect *pgx.Conn

func init() {
	connection, err := pgx.Connect(context.Background(), "postgres://postgres:postgres@localhost:5432/test_task")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connection to database: %v\n", err)
	} else {
		fmt.Println("Connected to PSQL!")
	}

	Connect = connection
}

func AddNewOrder(new_order Order) {
	_, err := Connect.Exec(context.Background(), "INSERT INTO public.orders VALUES($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)",
		new_order.Idorder, new_order.Uuiduser, new_order.Product_name, new_order.Product_id,
		new_order.Price, new_order.Payment, new_order.Discount, new_order.Shipping_delivery,
		new_order.Record_time, new_order.Record_date)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
}
