package productcontroller

import (
	"github.com/Zale29/GoRestAPI/GoRestAPI/controllers/productcontroller"
	"github.com/Zale29/GoRestAPI/GoRestAPI/models"
	"github.com/gin-gonic/gin"
)

// Show
func Index(c *gin Context) {
	
	var products []models.Product

	models.DB.Find(&products)
	c.JSON(http.StatusOK, gin.H{"Products :" products})

}

// Create
func Show(c *gin Context) {

}

// Update
func Create(c *gin Context) {

}

// Delete
func Delete(c *gin Context) {

}
