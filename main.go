package main

import (
	"github.com/Zale29/GoRestAPI/controllers/productcontroller"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/products", productcontroller.Delete)

	r.Run()

}
