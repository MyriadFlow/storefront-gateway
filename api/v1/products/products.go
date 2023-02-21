package products

import (
	"net/http"
	"github.com/MyriadFlow/storefront_gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	"github.com/MyriadFlow/storefront_gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront_gateway/models"
	"github.com/sirupsen/logrus"
)


// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/products")
	{
		g.Use(paseto.PASETO)
		g.PATCH("/", patchProduct)
		g.GET("/itemIds", getItemIds)
		g.GET("/itemId/:id", getItemById)
		g.GET("/trending", getTrending)
		g.GET("/highlights", getHighlights)
		g.GET("/marketplaceCounts", getMarketPlaceCounts)
		g.GET("/marketplaceInfo", getmarketplaceInfo)
	}
}
var productIds []string
func patchProduct(c *gin.Context) {
	db := dbconfig.GetDb()
	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}

	itemId := product.ItemId
	result := db.Model(&models.Product{}).Where("item_id = ?", itemId).Updates(product)

	if result.Error != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

		return
	}
	if result.RowsAffected == 0 {
		httphelper.ErrResponse(c, http.StatusNotFound, "Record not found")

		return
	}
	httphelper.SuccessResponse(c, "Profile successfully updated", nil)

}

func getProductsIds(c *gin.Context){
	db := dbconfig.GetDb()
	err := db.Table("products").Order("item_id").Select("item_id").Find(&productIds).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}	
}

func getItemIds(c *gin.Context) {
	db := dbconfig.GetDb()
	var list []string
	err := db.Table("products").Order("item_id").Select("item_id").Find(&list).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	productIds=list
	httphelper.SuccessResponse(c, "Profile fetched successfully", list)
}
func getItemById(c *gin.Context) {
	db := dbconfig.GetDb()
	var product models.Product
	id := c.Params.ByName("id")
	err := db.Table("products").Where("item_id = ?", id).First(&product).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Profile fetched successfully", product)
}
func getTrending(c *gin.Context) {
	getProductsIds(c)
	httphelper.SuccessResponse(c, "Profile fetched successfully", productIds)
}

func getHighlights(c *gin.Context) {
	getProductsIds(c)
	httphelper.SuccessResponse(c, "Profile fetched successfully", productIds)
}

func getMarketPlaceCounts(c *gin.Context) {
	payload:=[]MarketPlaceCounts{{7878,4545,5546}}
	httphelper.SuccessResponse(c, "Profile fetched successfully", payload)
}

func getmarketplaceInfo(c *gin.Context) {
	payload:=[]MarketPlaceInfo{{"MyriadFlow","An innovative platform to explore & launch NFT Experiences.","@0xMyriadFlow"}}
	httphelper.SuccessResponse(c, "Profile fetched successfully", payload)
}