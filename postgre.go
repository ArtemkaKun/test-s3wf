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
		new_order.Record_date, new_order.Record_time)
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
	}
}

func GetOrderInfo(order_id string) (LogErr, Order) {
	var log_err LogErr
	var order_info Order

	err := Connect.QueryRow(context.Background(), "SELECT uuiduser, product_name, product_id, price, payment, discount, shipping, record_date, record_time FROM orders WHERE idorder=$1", order_id).Scan(&order_info.Uuiduser,
		&order_info.Product_name, &order_info.Product_id, &order_info.Price, &order_info.Payment,
		&order_info.Discount, &order_info.Shipping_delivery, &order_info.Record_date, &order_info.Record_time)
	if err != nil {
		log_err = LogErr{Status: "error", Message: fmt.Sprintf("%v", err.Error())}
		return log_err, order_info
	}

	return log_err, order_info
}
