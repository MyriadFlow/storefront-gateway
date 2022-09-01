package uploadtoipfs

import (
	jwtMiddleWare "github.com/TheLazarusNetwork/marketplace-engine/api/middleware/auth/jwt"
	"github.com/TheLazarusNetwork/marketplace-engine/config/envconfig"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/httphelper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"

	"github.com/gin-gonic/gin"
	ipfsGateway "github.com/ipfs/go-ipfs-api"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/uploadtoipfs")
	{
		g.Use(jwtMiddleWare.JWT)
		g.POST("", uploadtoipfs)
	}
}

func uploadtoipfs(c *gin.Context) {
	ig := ipfsGateway.NewShell(envconfig.EnvVars.IPFS_NODE_URL)

	form, e := c.MultipartForm()
	if e != nil {
		httphelper.NewInternalServerError(c, "failed to parse multipart form", "failed to parse multipart form, error: %v", e.Error())
		return
	}
	files := form.File["file"]

	responsePayload := make([]UploadToIpfsPayload, 0)
	for _, file := range files {
		fO, e := file.Open()

		if e != nil {
			httphelper.NewInternalServerError(c, "failed to load file", "failed to open file, error: %v", e.Error())
			return
		}
		//Uploads file to ipfs and returns metahash
		hash, err := ig.Add(fO)
		if err != nil {
			httphelper.NewInternalServerError(c, "failed to add file in ipfs", "failed to add file to ipfs, error: %v", e.Error())
			return
		}
		responsePayload = append(responsePayload, UploadToIpfsPayload{
			Name: file.Filename,
			Hash: hash,
		})
		err = fO.Close()
		if err != nil {
			logwrapper.Warnf("failed to close file: %v", file.Filename)
		}
	}

	httphelper.SuccessResponse(c, "file successfully added to ipfs", responsePayload)
}
