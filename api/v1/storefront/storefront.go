package storefront

import (
	"net/http"
	"time"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	storefrontUtil "github.com/MyriadFlow/storefront-gateway/util/pkg/storefront"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/storefront")
	{
		g.POST("", NewStorefront)
		g.PUT("", UpdateStorefront)
		g.GET("", GetStorefronts)
	}
}

func NewStorefront(c *gin.Context) {
	var StorefrontRequest StorefrontRequest
	if err := c.BindJSON(&StorefrontRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := storefrontUtil.CreateStorefront(StorefrontRequest.Name, StorefrontRequest.Owner, StorefrontRequest.WalletAdress, StorefrontRequest.Plan, StorefrontRequest.Cost, StorefrontRequest.Currency, StorefrontRequest.CreatedBy, StorefrontRequest.UpdatedBy, StorefrontRequest.Image, StorefrontRequest.Headline, StorefrontRequest.Description, StorefrontRequest.Blockchain)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Storefront created successfully"})
}

func UpdateStorefront(c *gin.Context) {
	db := dbconfig.GetDb()
	var updateRequest UpdateStorefrontRequest
	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var storefront models.Storefront
	result := db.Where("id = ?", updateRequest.Id).First(&storefront)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	date, err := time.Parse("2006-01-02", updateRequest.Validity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	storefront.Status = updateRequest.Status
	storefront.Validity = date
	storefront.UpdatedBy = updateRequest.UpdatedBy
	storefront.UpdatedAt = time.Now()
	result = db.Save(&storefront)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Storefront updated successfully"})
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