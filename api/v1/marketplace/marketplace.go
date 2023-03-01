package marketplace

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/marketplace")
	{
		g.Use(paseto.PASETO)
		g.PATCH("/itemId/:id", patchMarketplaceById)
		g.GET("/itemIds", getMarketplaceItemIds)
		g.GET("/itemId/:id", getMarketplaceById)

	}
}

func patchMarketplaceById(c *gin.Context) {
	db := dbconfig.GetDb()
	var product models.Marketplace
	err := c.BindJSON(&product)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}

	itemId := c.Params.ByName("id")
	result := db.Model(&models.Marketplace{}).Where("item_id = ?", itemId).Updates(product)

	if result.Error != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

		return
	}
	if result.RowsAffected == 0 {
		httphelper.ErrResponse(c, http.StatusNotFound, "Record not found")

		return
	}
	httphelper.SuccessResponse(c, "Marketplace Details successfully updated", nil)

}

func getMarketplaceItemIds(c *gin.Context) {
	db := dbconfig.GetDb()
	var list []string
	//err := db.Table("marketplace").Order("item_id").Select("item_id").Find(&list).Error
	err := db.Model(&models.Marketplace{}).Order("item_id").Select("item_id").Find(&list).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Marketplace ItemIds List fetched successfully", list)
}

func getMarketplaceById(c *gin.Context) {
	db := dbconfig.GetDb()
	var product models.Marketplace
	id := c.Params.ByName("id")
	//err := db.Table("marketplace").Where("item_id = ?", id).First(&product).Error
	err := db.Model(&models.Marketplace{}).Where("item_id = ?", id).First(&product).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Marketplace Details fetched successfully", product)
}
