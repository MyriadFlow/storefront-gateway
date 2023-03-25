package likes

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
	g := r.Group("/likes")
	{
		g.Use(paseto.PASETO)
		g.GET("/:itemId", getAllUsersLikesCount)
		g.POST("/:itemId", postUserLike)
		g.DELETE("/:itemId", deleteUserLike)

	}
}

func getAllUsersLikesCount(c *gin.Context) {
	db := dbconfig.GetDb()
	itemId := c.Params.ByName("itemId")
	if itemId==""{
		httphelper.ErrResponse(c, http.StatusForbidden, "no itemId passed")
		return
	}

	var count int64
	err := db.Model(&models.Likes{}).Where("item_id = ?", itemId).Count(&count).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "All User Likes fetched successfully", count)
}

func deleteUserLike(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	itemId := c.Params.ByName("itemId")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusForbidden, "Unable to Parse item id")
		return
	}
	err = db.Where("item_id = ? AND user_wallet_address = ?", id,walletAddress).Delete(&models.Likes{}).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Item Unliked successfully", nil)
}

func postUserLike(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	itemId := c.Params.ByName("itemId")
	id, err := strconv.Atoi(itemId)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusForbidden, "Unable to Parse item id")
		return
	}
	//check if item already liked 
	isLiked:=checkIfAlreadyLiked(c)
	if isLiked {
		httphelper.ErrResponse(c, http.StatusForbidden, "Item Liked Already")
		return
	}
	fmt.Println("isLiked found :",isLiked)
	addUserLike := models.Likes{
		ItemId: id,
		UserWalletAddress: walletAddress,
	}
	err = db.Model(&models.Likes{}).Create(&addUserLike).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured : Unable to add Like")
		return
	}

	httphelper.SuccessResponse(c, "Item Liked successfully", nil)
}

func checkIfAlreadyLiked(c *gin.Context) bool {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	itemId := c.Params.ByName("itemId")
	var count int64
	err := db.Model(&models.Likes{}).Where("item_id = ? AND user_wallet_address = ?", itemId,walletAddress).Count(&count).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

	}
	fmt.Println("already exited count :",count)
	if count!=1 {
		return false //not liked yet
	}
	return true //already liked once
}
