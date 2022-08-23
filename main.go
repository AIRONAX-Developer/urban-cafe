package main

import (
	home "go-restaurent/services/apis/home"
	order "go-restaurent/services/apis/order"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	//  Display Menu for restaurent
	router.GET("/display-menu/:name", func(c *gin.Context) {
		name := c.Param("name")
		response := home.DisplayMenu(name)

		c.JSON(http.StatusOK, response)
	})

	// Order Placement for restaurent

	router.GET("/order-placement/:orderParam", func(c *gin.Context) {
		orderParam := c.Param("orderParam")
		response := order.PlaceOrder(orderParam)

		c.JSON(http.StatusOK, response)
	})

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	router.Run()
	// router.Run(":3000") for a hard coded port
}
