package highlights

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	//"fmt"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/highlights")
	{
		g.Use(paseto.PASETO)
		g.GET("", getHighlightsItemIds)
		g.DELETE("/:itemId", deleteHighlightsItemId)
		g.POST("", postHighlightsItemId)

	}
}

func postHighlightsItemId(c *gin.Context) {
	db := dbconfig.GetDb()
	var highlightItem models.Highlights
	err := c.BindJSON(&highlightItem)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}

	err = db.Model(&models.Highlights{}).Create(&highlightItem).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured : Unable to add Like")
		return
	}

	httphelper.SuccessResponse(c, "Highlights Details successfully updated", nil)

}

func getHighlightsItemIds(c *gin.Context) {
	db := dbconfig.GetDb()
	var list []string
	err := db.Model(&models.Highlights{}).Order("item_id").Select("item_id").Find(&list).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Highlights ItemIds List fetched successfully", list)
}

func deleteHighlightsItemId(c *gin.Context) {
	db := dbconfig.GetDb()
	id := c.Params.ByName("itemId")

	err := db.Where("item_id = ?", id).Delete(&models.Highlights{}).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Highlights ItemId deleted successfully", nil)
}
