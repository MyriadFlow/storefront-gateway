package likes

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/MyriadFlow/storefront-gateway/util/testingcommon"
	"github.com/stretchr/testify/assert"
)

func Test_PostUserLike(t *testing.T){
	testingcommon.InitializeEnvVars()
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	testItemId:="1111"
	t.Run("Should be able post User Like ", func(t *testing.T) {
		rr := httptest.NewRecorder()
		url := "/api/v1.0/likes"
		req, err := http.NewRequest("POST", url, nil)
		req.Header.Add("Authorization", header)
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)

		c.Request = req
		c.Params = []gin.Param{
						{
							Key: "itemId",
							Value: testItemId,
						},
					}
		c.Set("walletAddress", testWallet.WalletAddress)
		postUserLike(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
		//delete test like entry
		deleteUserLike(c)
	})

}

func Test_GetAllUsersLikesCount(t *testing.T){
	testingcommon.InitializeEnvVars()
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	testItemId:="1111"
	t.Run("Should be able get Likes Count on ItemId passed ", func(t *testing.T) {

		rr := httptest.NewRecorder()
		url := "/api/v1.0/likes"
		req, err := http.NewRequest("GET", url,nil)
		req.Header.Add("Authorization", header)
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)
		c.Params = []gin.Param{
			{
				Key: "itemId",
				Value: testItemId,
			},
		}
		c.Request = req
		c.Set("walletAddress", testWallet.WalletAddress)
		getAllUsersLikesCount(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

}

func Test_DeleteUserLike(t *testing.T){
	testingcommon.InitializeEnvVars()
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	testItemId:="1111"
	t.Run("Should be able delete Like on ItemId passed", func(t *testing.T) {
		//insert a test like entry
		rr := httptest.NewRecorder()
		url := "/api/v1.0/likes"
		req, err := http.NewRequest("POST", url, nil)
		req.Header.Add("Authorization", header)
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		c.Params = []gin.Param{
						{
							Key: "itemId",
							Value: testItemId,
						},
					}
		c.Set("walletAddress", testWallet.WalletAddress)
		postUserLike(c)

		//new request for delete 
		rr = httptest.NewRecorder()
		url = "/api/v1.0/likes"
		req, err = http.NewRequest("DELETE", url,nil)
		req.Header.Add("Authorization", header)
		if err != nil {
			t.Fatal(err)
		}
		c, _ = gin.CreateTestContext(rr)

		c.Request = req
		c.Params = []gin.Param{
			{
				Key: "itemId",
				Value: testItemId,
			},
		}
		c.Set("walletAddress", testWallet.WalletAddress)
		deleteUserLike(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

}

