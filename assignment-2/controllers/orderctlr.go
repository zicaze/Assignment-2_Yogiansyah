package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	dt "assignment-2/structs"

	"github.com/gin-gonic/gin"
)

var orders []dt.Order

// create new data to the database
func CreateOrder(c *gin.Context) {
	var newOrder dt.Order
	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	orders = append(orders, newOrder)

	c.JSON(http.StatusCreated, gin.H{
		"orders": newOrder,
		"status": http.StatusCreated,
	})
}

// get all data order
func GetOrders(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{
		"orders": orders,
		"status": fmt.Sprintf("%d", http.StatusOK),
	})
}

// update data with {id} as query
func UpdateOrder(c *gin.Context) {
	var newOrder dt.Order

	if err := c.ShouldBindJSON(&newOrder); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	orderId := c.Param("OrderID")
	convertOrderID, _ := strconv.Atoi(orderId)

	var cond bool

	for i, ordr := range orders {
		if uint(convertOrderID) == ordr.OrderID {
			cond = true
			orders[i] = newOrder
		}
	}

	if !cond {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": fmt.Sprintf("%d", http.StatusOK),
	})

}

// delete data with {id}
func DeleteOrder(c *gin.Context) {

	orderId := c.Param("OrderID")
	convertOrderID, _ := strconv.Atoi(orderId)
	var indexorder int

	var iExist bool

	for i, ordr := range orders {
		if uint(convertOrderID) == ordr.OrderID {
			iExist = true
			indexorder = i
			break
		}
	}

	if !iExist {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status": "Data Not Found",
		})
		return
	}

	copy(orders[indexorder:], orders[indexorder+1:])
	orders[len(orders)-1] = dt.Order{}
	orders = orders[:len(orders)-1]

	c.JSON(http.StatusOK, gin.H{
		"status": fmt.Sprintf("%d", http.StatusOK),
	})

}
