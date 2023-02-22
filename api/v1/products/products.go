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
		g.PATCH("/itemId/:id", patchProductById)
		g.GET("/itemIds", getProductItemIds)
		g.GET("/itemId/:id", getProductById)

	}
}

func patchProductById(c *gin.Context) {
	db := dbconfig.GetDb()
	var product models.Product
	err := c.BindJSON(&product)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}

	itemId := c.Params.ByName("id")
	result := db.Model(&models.Product{}).Where("item_id = ?", itemId).Updates(product)

	if result.Error != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

		return
	}
	if result.RowsAffected == 0 {
		httphelper.ErrResponse(c, http.StatusNotFound, "Record not found")

		return
	}
	httphelper.SuccessResponse(c, "Product Details successfully updated", nil)

}

func getProductItemIds(c *gin.Context) {
	db := dbconfig.GetDb()
	var list []string
	err := db.Table("products").Order("item_id").Select("item_id").Find(&list).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Product ItemIds List fetched successfully", list)
}

func getProductById(c *gin.Context) {
	db := dbconfig.GetDb()
	var product models.Product
	id := c.Params.ByName("id")
	err := db.Table("products").Where("item_id = ?", id).First(&product).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Product Details fetched successfully", product)
}
