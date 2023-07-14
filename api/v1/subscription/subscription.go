package subscription

import (
	"net/http"
	"time"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/subscription")
	{
		g.POST("/new", Subscribe)
	}
}

func Subscribe(c *gin.Context) {
	db := dbconfig.GetDb()
	var subRequest subscriptionRequest
	if err := c.BindJSON(&subRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	subscription := models.Subscription{
		Id:        uuid.New(),
		Name:      subRequest.Name,
		Owner:     subRequest.Owner,
		Plan:      subRequest.Plan,
		Cost:      subRequest.Cost,
		Currency:  subRequest.Currency,
		Status:    "active",
		Validity:  time.Now().AddDate(1, 0, 0),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		CreatedBy: subRequest.CreatedBy,
		UpdatedBy: subRequest.UpdatedBy,
	}
	result := db.Create(&subscription)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Subscription created successfully"})
}
