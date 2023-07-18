package subscription

import (
	"net/http"
	"time"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/subscription"
	"github.com/gin-gonic/gin"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/subscription")
	{
		g.POST("", Subscribe)
		g.PUT("", Update)
		g.GET("", GetSubscriptions)
	}
}

func Subscribe(c *gin.Context) {
	var subRequest subscriptionRequest
	if err := c.BindJSON(&subRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	err := subscription.CreateSubscription(subRequest.Name, subRequest.Owner, subRequest.Plan, subRequest.Cost, subRequest.Currency, subRequest.CreatedBy, subRequest.UpdatedBy)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subscription created successfully"})
}

func Update(c *gin.Context) {
	db := dbconfig.GetDb()
	var updateRequest UpdateSubscriptionRequest
	if err := c.BindJSON(&updateRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	var subscription models.Subscription
	result := db.Where("id = ?", updateRequest.Id).First(&subscription)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	date, err := time.Parse("2006-01-02", updateRequest.Validity)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	subscription.Status = updateRequest.Status
	subscription.Validity = date
	subscription.UpdatedBy = updateRequest.UpdatedBy
	subscription.UpdatedAt = time.Now()
	result = db.Save(&subscription)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subscription updated successfully"})
}

func GetSubscriptions(c *gin.Context) {
	db := dbconfig.GetDb()
	var subscriptions []models.Subscription
	result := db.Find(&subscriptions)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, subscriptions)
}
