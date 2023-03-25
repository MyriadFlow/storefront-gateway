package wishlist

import (
	"net/http"
	"strconv"
	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"fmt"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/wishlist")
	{
		g.Use(paseto.PASETO)
		g.GET("", getUserWishlist)
		g.POST("/:itemId", postItemIdToUserWishlist)
		g.DELETE("/:itemId", deleteFromUserWishlist)

	}
}

//likes api
func getUserWishlist(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var response []string
	err := db.Model(&models.Wishlist{}).Order("item_id").Select("item_id").Where("user_wallet_address = ?", walletAddress).Find(&response).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Wishlist fetched successfully", response)
}

func deleteFromUserWishlist(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	itemId := c.Params.ByName("itemId")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusForbidden, "Unable to Parse item id")
		return
	}
	err = db.Where("item_id = ? AND user_wallet_address = ?", id,walletAddress).Delete(&models.Wishlist{}).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Item removed from Wishlist successfully", nil)
}

func postItemIdToUserWishlist(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	itemId := c.Params.ByName("itemId")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusForbidden, "Unable to Parse item id")
		return
	}
	//check if item already wishlisted
	isWishlisted:=checkIfAlreadyWishlisted(c)
	if isWishlisted {
		httphelper.ErrResponse(c, http.StatusForbidden, "Item Liked Already")
		return
	}
	addToWishlist := models.Wishlist{
		ItemId: id,
		UserWalletAddress: walletAddress,
	}
	err = db.Model(&models.Wishlist{}).Create(&addToWishlist).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured : Unable to create Wishlist")
		return
	}

	httphelper.SuccessResponse(c, "Item Wishlisted successfully", nil)
}

func checkIfAlreadyWishlisted(c *gin.Context) bool {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	itemId := c.Params.ByName("itemId")
	var count int64
	err := db.Model(&models.Wishlist{}).Where("item_id = ? AND user_wallet_address = ?", itemId,walletAddress).Count(&count).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
	}

	fmt.Println("already exited count :",count)
	if count!=1 {
		return false //not wishlisted yet
	}
	return true //already wishlisted once
}