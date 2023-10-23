package productcontroller

import (
	"net/http"

	"github.com/Zale29/GoRestAPI/models"
	"github.com/gin-gonic/gin"
)

// Show
func Index(c *gin.Context) {

	var Products []models.Product

	models.DB.Find(&Products)
	c.JSON(http.StatusOK, gin.H{"Products ": Products})

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
