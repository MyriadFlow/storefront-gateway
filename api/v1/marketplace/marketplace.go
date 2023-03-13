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
		//like apis
		g.GET("/userLikes/:contractaddress/:tokenid", getUserLikes)
		g.GET("/likes/:contractaddress/:tokenid", getAllUsersLike)
		g.DELETE("/userLikes/:contractaddress/:tokenid", deleteUserLike)
		g.POST("/userLike", postUserLike)

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


//likes api

func getUserLikes(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	contractaddress := c.Params.ByName("contractaddress")
	if contractaddress==""{
		httphelper.ErrResponse(c, http.StatusForbidden, "payload has no contractaddress")
		return
	}
	tokenid := c.Params.ByName("tokenid")
	if tokenid==""{
		httphelper.ErrResponse(c, http.StatusForbidden, "payload has no tokenid")
		return
	}
	var count int64
	err := db.Model(&models.Likes{}).Where("nft_contract_address = ? AND token_id = ? AND user_wallet_address = ?", contractaddress,tokenid,walletAddress).Count(&count).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "User Likes fetched successfully", count)
}

func getAllUsersLike(c *gin.Context) {
	db := dbconfig.GetDb()
	contractaddress := c.Params.ByName("contractaddress")
	if contractaddress==""{
		httphelper.ErrResponse(c, http.StatusForbidden, "payload has no contractaddress")
		return
	}
	tokenid := c.Params.ByName("tokenid")
	if tokenid==""{
		httphelper.ErrResponse(c, http.StatusForbidden, "payload has no tokenid")
		return
	}

	var count int64
	err := db.Model(&models.Likes{}).Where("nft_contract_address = ? AND token_id = ?", contractaddress,tokenid).Count(&count).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "All User Likes fetched successfully", count)
}

func deleteUserLike(c *gin.Context) {
	db := dbconfig.GetDb()
	contractaddress := c.Params.ByName("contractaddress")
	if contractaddress==""{
		httphelper.ErrResponse(c, http.StatusForbidden, "payload has no contractaddress")
		return
	}
	tokenid := c.Params.ByName("tokenid")
	if tokenid==""{
		httphelper.ErrResponse(c, http.StatusForbidden, "payload has no tokenid")
		return
	}
	walletAddress := c.GetString("walletAddress")
	err := db.Where("nft_contract_address = ? AND token_id = ? AND user_wallet_address = ?", contractaddress,tokenid,walletAddress).Delete(&models.Likes{}).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Unlike done successfully", nil)
}

func postUserLike(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var payload LikesQueryPayload
	err := c.BindJSON(&payload)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}
	addLike := models.Likes{
		NFT_Contract_Address: payload.NFT_Contract_Address,
		TokenId: payload.TokenId,
		UserWalletAddress:walletAddress,
	}
	err = db.Model(&models.Likes{}).Create(&addLike).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured : Unable to create Like")
		return
	}

	httphelper.SuccessResponse(c, "Like done successfully", nil)
}