package order

import (
	"fmt"
)

type OrderStruct struct {
	OrderNumber string `json:"orderNumber"`
	OrderStatus string `json:"orderStatus"`
}

func PlaceOrder(order string) *OrderStruct {
	fmt.Println("Place Order: ", order)

	orderResponse := &OrderStruct{
		OrderNumber: "30",
		OrderStatus: "Success",
	}
	return orderResponse
}
