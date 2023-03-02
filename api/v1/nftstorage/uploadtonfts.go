package nftstorage

import (
	"context"
	"encoding/json"
	"os"

	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	client "github.com/nftstorage/go-client"
	"github.com/sirupsen/logrus"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/nftstorage")
	{
		g.Use(paseto.PASETO)
		g.POST("/upload", uploadtonftstorage)
	}
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

		responsePayload = append(responsePayload, NftStorageUploadResponse{file.Filename, string(cid)})

	}
	logrus.Info("\n responsePayload :", responsePayload)
	httphelper.SuccessResponse(c, "file successfully uploaded to nft storage", responsePayload)
}
