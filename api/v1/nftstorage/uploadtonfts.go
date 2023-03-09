package nftstorage

import (
	"context"
	"encoding/json"
	"os"
	"net/http"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/gin-gonic/gin"
	client "github.com/nftstorage/go-client"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/nftstorage")
	{
		g.Use(paseto.PASETO)
		g.POST("/upload", uploadtonftstorage)
		g.POST("/like", updatenftlikes)
		g.GET("/likes/:cid", getnftlikes)
	}
}
func insertcidtodb(cid string)(error){
	//insert record with like count equals to 1
	nftlikes := models.NftLikes{
		Cid: cid,
		Likes: 1,
	}
	db := dbconfig.GetDb()
	err := db.Model(&models.NftLikes{}).Create(&nftlikes).Error
	return err
}
func uploadtonftstorage(c *gin.Context) {

	token := envconfig.EnvVars.NFT_STORAGE_API_KEY
	configuration := client.NewConfiguration()
	ctx := context.WithValue(context.Background(), client.ContextAccessToken, token)
	api_client := client.NewAPIClient(configuration)

	form, err := c.MultipartForm()
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to parse multipart form", "failed to parse multipart form, error: %v", err.Error())
		return
	}

	responsePayload := make([]NftStorageUploadResponse, 0)
	files := form.File["file"]

	for _, file := range files {
		//tempporarily storing multipart file and then read as os file
		filename := "./" + file.Filename
		err := c.SaveUploadedFile(file, filename)
		if err != nil {
			httphelper.NewInternalServerError(c, "failed to SaveUpload file", "failed to open file, error: %v", err.Error())
			return
		}
		logrus.Info("\n File Saved at :", filename)
		bO, err := os.Open(filename)
		if err != nil {
			httphelper.NewInternalServerError(c, "failed to load file", "failed to open file, error: %v", err.Error())
			return
		}

		resp, r, err := api_client.NFTStorageAPI.Store(ctx).Body(bO).Execute()
		if err != nil {
			httphelper.NewInternalServerError(c, "failed to upload ", "Error when calling `NFTStorageAPI.Store``: error: %v", err.Error())
			httphelper.NewInternalServerError(c, "failed to upload ", "Full HTTP response, error: %v", r)
			return
		}
		bO.Close()
		err = os.Remove(filename)
		if err != nil {
			httphelper.NewInternalServerError(c, "failed to clear temporary file stored", "failed to clear temporary file stored, error: %v", err.Error())
			return
		}
		cid, _ := json.Marshal(resp.Value.Cid)
		//remove ' " ' from extrem ends
		cid = cid[1:]
		cid = cid[:len(cid)-2]
		
		//add the cid for likes count
		err = insertcidtodb(string(cid))
		if err != nil {
			httphelper.ErrResponse(c, http.StatusInternalServerError, "Fail to insert NFT likes record")
			return 
		}

		responsePayload = append(responsePayload, NftStorageUploadResponse{file.Filename, string(cid)})

	}

	httphelper.SuccessResponse(c, "file successfully uploaded to nft storage", responsePayload)
}

func updatenftlikes(c *gin.Context) {

	db := dbconfig.GetDb()
	var nftcid LikeNft
	
	err := c.BindJSON(&nftcid)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}
	
	result := db.Model(&models.NftLikes{}).Where("cid = ?", nftcid.CID).Update("likes", gorm.Expr("likes + ?", 1))
	if result.Error != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured : Unable to update Likes")
		return
	}
	if result.RowsAffected == 0 {
		httphelper.ErrResponse(c, http.StatusNotFound, "Record not found")
		return
	}

	httphelper.SuccessResponse(c, "Like successfully updated", nil)

}

func getnftlikes(c *gin.Context) {
	db := dbconfig.GetDb()
	var nftlikes models.NftLikes
	cid := c.Params.ByName("cid")

	err := db.Model(&models.NftLikes{}).Where("cid = ?", cid).First(&nftlikes).Error
	if err != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	httphelper.SuccessResponse(c, "Likes fetched successfully", nftlikes.Likes)

}
