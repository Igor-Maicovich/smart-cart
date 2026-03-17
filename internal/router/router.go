package router

import (
	"smart-cart/internal/cart"

	"github.com/gin-gonic/gin"
)

func SetupRouter(handler *cart.Handler) *gin.Engine {
	r := gin.Default()

	cartGroup := r.Group("/cart")
	{
		cartGroup.GET("", handler.GetCart)
		cartGroup.POST("/items", handler.AddItem)
		cartGroup.PATCH("/items/:id", handler.UpdateItem)
		cartGroup.DELETE("/items/:id", handler.DeleteItem)
	}

	return r
}
