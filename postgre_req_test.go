package main

import "testing"

func TestAddNewOrder(t *testing.T) {
	//new_orders := []Order{
	//	{2, 3, "Banana", "2", 2.00, "paypal",
	//		"none", "27.03.2022", "26.02.2020", "00:00"},
	//	{3, 1, "Banana", "2", 1.80, "paypal",
	//		"2%", "27.03.2022", "26.02.2020", "00:01"},
	//}
	//
	//for _, order := range new_orders {
	//	AddNewOrder(order)
	//}

	testpair := []OrderInfoTestpair{
		{"2", Order{0, 3, "Banana", "2", 2.00, "paypal",
			"none", "27.03.2022", "26.02.2020", "00:00"}, ""},
		{"4", Order{}, "error"},
	}

	for _, pair := range testpair {
		err, order := GetOrderInfo(pair.Test_orderid)
		if err.Status != pair.Err_status || order != pair.Test_order {
			t.Error("Expect: ", pair.Err_status, pair.Test_order,
				"Got: ", err.Status, order)
		}
	}
}
