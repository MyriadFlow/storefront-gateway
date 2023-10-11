package website

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/website")
	{
		g.GET("/storefronts", GetStorefronts)
	}
}
func GetStorefronts(c *gin.Context) {
	db := dbconfig.GetDb()
	var storefronts []models.Storefront
	result := db.Find(&storefronts)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, storefronts)
}
