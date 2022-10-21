package routers

import (
	ctrl "assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func ServerOn() *gin.Engine {
	router := gin.Default()
	router.POST("/orders", ctrl.CreateOrder)
	router.GET("/orders/", ctrl.GetOrders)
	router.PUT("/orders/:OrderID", ctrl.UpdateOrder)
	router.DELETE("/orders/:OrderID", ctrl.DeleteOrder)
	return router
}
