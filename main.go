package main

import (
	"github.com/Zale29/gorestapi/controller/productcontroller"
)

func main() {
	r := git.Default()

	r.GET("/api/products", productcontroller.Index)
	r.GET("/api/products:id", productcontroller.Show)
	r.POST("/api/products", productcontroller.Create)
	r.PUT("/api/products/:id", productcontroller.Update)
	r.DELETE("/api/products", productcontroller.Delete)

	r.Run()

}
