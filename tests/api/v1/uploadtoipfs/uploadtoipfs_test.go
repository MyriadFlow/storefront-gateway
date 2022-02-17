package uploadtoipfs

import (
	"bytes"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/TheLazarusNetwork/marketplace-engine/app"
	"github.com/TheLazarusNetwork/marketplace-engine/util/testingcommon"
	"github.com/go-playground/assert/v2"
)

func Test_UploadToIpfs(t *testing.T) {
	app.Init("../../../../.env", "../../../../logs")
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)

	url := "/api/v1.0/uploadtoipfs"

	rr := httptest.NewRecorder()

	body := new(bytes.Buffer)
	mw := multipart.NewWriter(body)

	filesToUpload := []string{"cardata.json", "bikedata.yaml"}
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
	app.GinApp.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
}
