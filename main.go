package main

import (
	order "go-restaurent/app/apis/order"
	helper "go-restaurent/app/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	router := gin.Default()

	v1 := router.Group("/v1")
	{
		v1.GET("/display-menu/:name", displayMenuEndpoint)
		v1.GET("/order-placement/:orderParam", orderPlacementEndpoint)
		v1.GET("/qrcode/:uuidv4/", uuidEndpoint)

	}
	router.Run()
}

func displayMenuEndpoint(c *gin.Context) {
	name := c.Param("name")
	response := order.DisplayMenu(name)

	// c.JSON(http.StatusOK, response)
	c.JSON(http.StatusOK, gin.H{
		"message": response,
	})
}

func orderPlacementEndpoint(c *gin.Context) {
	orderParam := c.Param("orderParam")
	response := order.PlaceOrder(orderParam)

	c.JSON(http.StatusOK, response)
}

func uuidEndpoint(c *gin.Context) {
	uuidv4 := c.Param("uuidv4")
	helper.SugarObj.Info(uuidv4)
	res := helper.GenQR(uuidv4)
	c.JSON(http.StatusOK, res)
}
