package profile

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
	g := r.Group("/profile")
	{
		g.Use(paseto.PASETO)
		g.PATCH("", patchProfile)
		g.GET("", getProfile)
		g.PATCH("/verify", verifySocial)
		g.PATCH("/subscribe", BasicSubscription)
	}
}

func patchProfile(c *gin.Context) {
	db := dbconfig.GetDb()
	var user models.User
	err := c.BindJSON(&user)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}
	walletAddress := c.GetString("walletAddress")
	user.WalletAddress = walletAddress
	result := db.Model(&models.User{}).Where("wallet_address = ?", walletAddress).Updates(user)

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

func getProfile(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var user models.User
	//err := db.Model(&models.User{}).Select("name, profile_picture_url,country, wallet_address").Where("wallet_address = ?", walletAddress).First(&user).Error
	err := db.Model(&models.User{}).Where("wallet_address = ?", walletAddress).First(&user).Error
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

		return
	}

	payload := GetProfilePayload{
		user.Name, user.WalletAddress, user.ProfilePicture, user.CoverPicture, user.Location, user.FacebookId, user.InstagramId, user.TwitterId, user.DiscordId, user.TelegramId, user.Email, user.Bio, user.InstagramVerified, user.FacebookVerified, user.TwitterVerified, user.DiscordVerified, user.TelegramVerified,
	}
	httphelper.SuccessResponse(c, "Profile fetched successfully", payload)
}

func verifySocial(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	var req verifySocialPayload
	err := db.Model(&models.User{}).Where("wallet_address = ?", walletAddress).Update(req.SocialName+"_verified", true)
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")

		return
	}
	httphelper.SuccessResponse(c, "verification for "+req.SocialName+" updated", nil)
}

func BasicSubscription(c *gin.Context) {
	db := dbconfig.GetDb()
	walletAddress := c.GetString("walletAddress")
	err := db.Model(&models.User{}).Where("wallet_address = ?", walletAddress).Update("plan", "basic")
	if err != nil {
		logrus.Error(err)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	httphelper.SuccessResponse(c, "Basic Plan Subscribed", nil)
}
