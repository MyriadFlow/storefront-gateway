package uploadtonfts

import (
	"github.com/MyriadFlow/storefront_gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront_gateway/config/envconfig"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/httphelper"
	"github.com/gin-gonic/gin"
	client "github.com/nftstorage/go-client"
	"context"
	"os"
	"encoding/json"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/uploadtonfts")
	{
		g.Use(paseto.PASETO)
		g.POST("", uploadtonfts)
	}
}


func uploadtonfts(c *gin.Context) {

	token:=envconfig.EnvVars.NFT_API_KEY
	configuration := client.NewConfiguration()
	ctx := context.WithValue(context.Background(), client.ContextAccessToken, token)
	api_client := client.NewAPIClient(configuration)
	
	form, err := c.MultipartForm()
	if err != nil {
		httphelper.NewInternalServerError(c, "failed to parse multipart form", "failed to parse multipart form, error: %v", err.Error())
		return
	}

	values:=form.Value
	file_names:=values["file"]

	responsePayload := make([]UploadToNftsPayload, 0)

	for _, file := range file_names {

		fO, err :=os.Open(file)

		if err != nil {
			httphelper.NewInternalServerError(c, "failed to load file", "failed to open file, error: %v", err.Error())
			return
		}

		resp, r, err := api_client.NFTStorageAPI.Store(ctx).Body(fO).Execute()
		if err != nil {
			httphelper.NewInternalServerError(c, "failed to upload ", "Error when calling `NFTStorageAPI.Store``: error: %v", err.Error())
			httphelper.NewInternalServerError(c, "failed to upload ", "Full HTTP response, error: %v", r)
			return
		}
		fO.Close()
		cid, _ := json.Marshal(resp.Value.Cid)

		responsePayload=append(responsePayload,UploadToNftsPayload{file,string(cid)})


	}

	httphelper.SuccessResponse(c, "file successfully uploaded to nft storage", responsePayload)
}

