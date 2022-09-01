package uploadtoipfs

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/TheLazarusNetwork/marketplace-engine/config/dbconfig/dbinit"
	"github.com/TheLazarusNetwork/marketplace-engine/util/pkg/logwrapper"
	"github.com/TheLazarusNetwork/marketplace-engine/util/testingcommon"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func Test_UploadToIpfs(t *testing.T) {
	logwrapper.Init()
	dbinit.Init()
	t.Cleanup(testingcommon.DeleteCreatedEntities())

	t.Run("Should be able to upload file to ipfs", func(t *testing.T) {
		testWallet := testingcommon.GenerateWallet()
		header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)

		url := "/api/v1.0/uploadtoipfs"

		rr := httptest.NewRecorder()

		body := new(bytes.Buffer)
		mw := multipart.NewWriter(body)

		filesToUpload := []string{"test/cardata.json", "test/bikedata.yaml"}
		for _, fileToUpload := range filesToUpload {
			file, err := os.Open(fileToUpload)
			if err != nil {
				t.Fatal(err)
			}
			w, err := mw.CreateFormFile("file", file.Name())
			if err != nil {
				t.Fatal(err)
			}
			if _, err := io.Copy(w, file); err != nil {
				t.Fatal(err)
			}
		}

		// close the writer before making the request
		mw.Close()

		req := httptest.NewRequest(http.MethodPost, url, body)
		req.Header.Add("Authorization", header)

		req.Header.Add("Content-Type", mw.FormDataContentType())
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		uploadtoipfs(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

}
