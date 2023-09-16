package webapp

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/webapp")
	{
		g.GET("/:id", DeployWebapp)
		g.GET("/contracts/:id", GetContractAddresses)
	}
}
func DeployWebapp(c *gin.Context) {
	db := dbconfig.GetDb()
	storefrontId := c.Param("id")
	var storefront models.Storefront
	err := db.Model(&models.Storefront{}).Where("id = ?", storefrontId).First(&storefront).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	c.JSON(http.StatusOK, storefront)
}

func GetContractAddresses(c *gin.Context) {
	storefrontId := c.Param("storefrontId")
	db := dbconfig.GetDb()
	var contracts []Contract
	if result := db.Model(models.Contract{}).Where("storefront_id = ?", storefrontId).Find(&contracts); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}
	c.JSON(http.StatusOK, contracts)

}
