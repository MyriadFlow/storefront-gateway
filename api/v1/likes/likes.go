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
	//"fmt"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/likes")
	{
		g.Use(paseto.PASETO)
		g.GET("/:contractAddress/:itemId", getAllUsersLikesCount)
		g.GET("/isliked/:contractAddress/:itemId", itemIsLiked)
		g.POST("/:contractAddress/:itemId", postUserLike)
		g.DELETE("/:contractAddress/:itemId", deleteUserLike)
	}
}

func getAllUsersLikesCount(c *gin.Context) {
	contractAddress := c.Param("contractAddress")
	itemIdStr := c.Param("itemId")
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured : Unable to parse string to int")
		return
	}
	db := dbconfig.GetDb()
	var count int64
	err = db.Model(&models.Likes{}).Where("item_id = ? AND contract_address = ?", itemId, contractAddress).Count(&count).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "All User Likes fetched successfully", count)
}

func deleteUserLike(c *gin.Context) {
	contractAddress := c.Param("contractAddress")
	itemIdStr := c.Param("itemId")
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured : Unable to parse string to int")
		return
	}

	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	err = db.Model(&models.Likes{}).Where("item_id = ? AND wallet_address = ? AND contract_address = ?", itemId, walletAddress, contractAddress).Delete(&models.Likes{}).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}

	httphelper.SuccessResponse(c, "Item Unliked successfully", nil)
}

func postUserLike(c *gin.Context) {
	contractAddress := c.Param("contractAddress")
	itemIdStr := c.Param("itemId")
	itemId, err := strconv.Atoi(itemIdStr)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured : Unable to parse string to int")
		return
	}

	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	//check if item already liked
	isLiked, err := checkIfAlreadyLiked(c)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Error checking already liked item")
		return
	}
	if isLiked {
		httphelper.ErrResponse(c, http.StatusForbidden, "Item Liked Already")
		return
	}

	addUserLike := models.Likes{
		ItemId:          itemId,
		ContractAddress: contractAddress,
		WalletAddress:   walletAddress,
	}
	err = db.Model(&models.Likes{}).Create(&addUserLike).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured : Unable to add Like")
		return
	}

	httphelper.SuccessResponse(c, "Item Liked successfully", nil)
}

func checkIfAlreadyLiked(c *gin.Context) (bool, error) {
	var req LikeReqeust
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var count int64
	err := db.Model(&models.Likes{}).Where("item_id = ? AND wallet_address = ? AND contract_address = ?", req.ItemId, walletAddress, req.ContractAddress).Count(&count).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return false, err
	}
	return count == 1, nil
}

func itemIsLiked(c *gin.Context) {
	var req LikeReqeust
	if err := c.BindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var count int64
	err := db.Model(&models.Likes{}).Where("item_id = ? AND wallet_address = ? AND contract_address = ?", req.ItemId, walletAddress, req.ContractAddress).Count(&count).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	if count == 1 {
		httphelper.SuccessResponse(c, "Item already liked", gin.H{
			"liked": true,
		})
	} else {
		httphelper.SuccessResponse(c, "Item not liked", gin.H{
			"liked": false,
		})
	}
}
