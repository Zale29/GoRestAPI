package productcontroller

import (
	"net/http"

	"github.com/Zale29/GoRestAPI/models"
	"github.com/gin-gonic/gin"
)

// Show
func Get(c *gin.Context) {

	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"Products ": products})

}

// Create
func Show(c *gin.Context) {

}

func Update(c *gin.Context) {

}

// Update
func Create(c *gin.Context) {

}

// Delete
func Delete(c *gin.Context) {

}
