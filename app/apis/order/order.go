package order

import (
	"fmt"
	"go-restaurent/app/helper"
)

type OrderStruct struct {
	OrderNumber string `json:"orderNumber"`
	OrderStatus string `json:"orderStatus"`
}

type DisplayMenuStruct struct {
	ItemName  string  `json:"itemName"`
	ItemImage string  `json:"itemImage"`
	ItemPrice float32 `json:"itemPrice"`
}

func DisplayMenu(name string) []*DisplayMenuStruct {

	var menuList []*DisplayMenuStruct

	for i := 0; i < 10; i++ {
		displayMenu := &DisplayMenuStruct{
			ItemImage: "https://static.toiimg.com/photo/76425511.cms",
			ItemName:  "Samosa",
			ItemPrice: float32(i),
		}

		menuList = append(menuList, displayMenu)
	}

	return menuList
}

func PlaceOrder(order string) *OrderStruct {
	fmt.Println("Place Order: ", order)

	res := helper.PaymentConfirmation()

	if res {

		go sendNoticationToCustomer()
		go sendOrderToKot()

		orderResponse := &OrderStruct{
			OrderNumber: "30",
			OrderStatus: "Success",
		}
		return orderResponse
	}

	return &OrderStruct{
		OrderNumber: "00",
		OrderStatus: "Failed",
	}
}

func sendOrderToKot() {

}

func sendNoticationToCustomer() {

}
