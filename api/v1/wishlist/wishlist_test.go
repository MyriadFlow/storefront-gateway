package wishlist

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"github.com/gin-gonic/gin"
	"github.com/MyriadFlow/storefront-gateway/util/testingcommon"
	"github.com/stretchr/testify/assert"
)

func Test_PostItemIdToUserWishlist(t *testing.T){
	testingcommon.InitializeEnvVars()
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	testItemId:="1111"
	t.Run("Should be able post User Wishlist item ", func(t *testing.T) {
		rr := httptest.NewRecorder()
		url := "/api/v1.0/wishlist"
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
		postItemIdToUserWishlist(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
		//delete test wishlist entry
		deleteFromUserWishlist(c)
	})

}

func Test_GetUserWishlist(t *testing.T){
	testingcommon.InitializeEnvVars()
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	// testItemId:="1111"
	t.Run("Should be able get wishlist on Wallet Address used ", func(t *testing.T) {

		rr := httptest.NewRecorder()
		url := "/api/v1.0/wishlist"
		req, err := http.NewRequest("GET", url,nil)
		req.Header.Add("Authorization", header)
		if err != nil {
			t.Fatal(err)
		}
		c, _ := gin.CreateTestContext(rr)
		// c.Params = []gin.Param{
		// 	{
		// 		Key: "itemId",
		// 		Value: testItemId,
		// 	},
		// }
		c.Request = req
		c.Set("walletAddress", testWallet.WalletAddress)
		getUserWishlist(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

}

func Test_DeleteFromUserWishlist(t *testing.T){
	testingcommon.InitializeEnvVars()
	testWallet := testingcommon.GenerateWallet()
	header := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
	testItemId:="1111"
	t.Run("Should be able delete Wishlist on ItemId passed", func(t *testing.T) {
		//insert a test like entry
		rr := httptest.NewRecorder()
		url := "/api/v1.0/wishlist"
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
		postItemIdToUserWishlist(c)

		//new request for delete 
		rr = httptest.NewRecorder()
		url = "/api/v1.0/wishlist"
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
		deleteFromUserWishlist(c)
		assert.Equal(t, http.StatusOK, rr.Result().StatusCode)
	})

}

