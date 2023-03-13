//go:build !race
// +build !race

package claimrole

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/MyriadFlow/storefront-gateway/api/types"
	roleid "github.com/MyriadFlow/storefront-gateway/api/v1/roleId"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig/dbinit"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/config/smartcontract"
	"github.com/MyriadFlow/storefront-gateway/config/smartcontract/auth"
	storefront "github.com/MyriadFlow/storefront-gateway/config/storefront"
	smartcontractstorefront "github.com/MyriadFlow/storefront-gateway/generated/smartcontract/storefront"
	"github.com/MyriadFlow/storefront-gateway/global"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
	"github.com/MyriadFlow/storefront-gateway/util/testingcommon"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func Test_PostClaimRole(t *testing.T) {
	defer time.Sleep(4 * time.Second)
	envconfig.InitEnvVars()
	logwrapper.Init()
	dbinit.Init()
	global.InitGlobal()
	t.Cleanup(testingcommon.DeleteCreatedEntities())
	gin.SetMode(gin.TestMode)

	t.Run("Should claim role if signature is correct", func(t *testing.T) {
		testWallet := testingcommon.GenerateWallet()
		headers := testingcommon.PrepareAndGetAuthHeader(t, testWallet.WalletAddress)
		url := "/api/v1.0/claimrole"
		rr := httptest.NewRecorder()
		requestRoleRes := requestRole(t, headers, testWallet.WalletAddress)
		signature := getSignature(requestRoleRes.Eula, requestRoleRes.FlowId, testWallet.PrivateKey)
		reqBody := ClaimRoleRequest{
			Signature: signature, FlowId: requestRoleRes.FlowId,
		}
		jsonBytes, _ := json.Marshal(reqBody)
		req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonBytes))
		if err != nil {
			t.Fatal(err)
		}
		req.Header.Add("Authorization", headers)
		c, _ := gin.CreateTestContext(rr)
		c.Request = req
		c.Set("walletAddress", testWallet.WalletAddress)
		postClaimRole(c)
		ok := assert.Equal(t, http.StatusOK, rr.Result().StatusCode, rr.Body.String())
		if !ok {
			t.FailNow()
		}
		client, err := smartcontract.GetClient()
		if err != nil {
			t.Fatal(err)
		}
		instance, err := storefront.GetInstance(client)
		if err != nil {
			t.Fatalf("failed to get instance for %v , error: %v", "STOREFRONT", err.Error())
		}
		creatorRole, err := storefront.GetRole(storefront.CREATOR_ROLE)
		if err != nil {
			t.Fatalf("failed to get role id for %v , error: %v", "CREATOR ROLE", err.Error())
		}
		addr := common.HexToAddress(testWallet.WalletAddress)
		roleGrantedChannel := make(chan *smartcontractstorefront.StoreRoleGranted, 10)

		authBindOpts, err := auth.GetAuth(client)

		if err != nil {
			t.Fatalf("failed to get auth, error: %v", err.Error())
		}
		subs, err := instance.WatchRoleGranted(nil, roleGrantedChannel, [][32]byte{creatorRole}, []common.Address{addr}, []common.Address{authBindOpts.From})
		if err != nil {
			t.Fatalf("failed to listen to an event %v, error: %v", "RoleGranted", err.Error())
		}

		//Check if role trasaction is successfull
		hasRole, err := instance.HasRole(nil, creatorRole, addr)
		if err != nil {
			t.Fatalf("failed to call %v smart contract function HasRole , error: %v", "STOREFRONT", err.Error())
		}
		success := false
		if !hasRole {
			go failAfter(t, &success, 10*time.Second, roleGrantedChannel)
			event := <-roleGrantedChannel
			subs.Unsubscribe()
			if event != nil && event.Account.String() != addr.String() {
				log.Fatal("user doesn't have role in blockchain")
			} else {
				success = true
			}
		}
	})

}

func failAfter(t *testing.T, success *bool, duration time.Duration, ch chan *smartcontractstorefront.StorefrontRoleGranted) {
	time.Sleep(duration)
	if !*success {
		close(ch)
		t.Errorf("didn't got any response from %v after %v", "RoleGranted", duration)
	}
}
func requestRole(t *testing.T, headers string, walletAddres string) roleid.GetRoleIdPayload {
	creatorRole, err := storefront.GetRole(storefront.CREATOR_ROLE)
	if err != nil {
		t.Fatalf("failed to get role id for %v , error: %v", "CREATOR ROLE", err.Error())
	}

	url := "/api/v1.0/roleId/" + hexutil.Encode(creatorRole[:])
	rr := httptest.NewRecorder()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Add("Authorization", headers)
	c, _ := gin.CreateTestContext(rr)
	c.Params = gin.Params{{Key: "roleId", Value: hexutil.Encode(creatorRole[:])}}
	c.Request = req
	c.Set("walletAddress", walletAddres)
	roleid.GetRoleId(c)
	if rr.Result().StatusCode != 200 {
		t.Fatalf("failed to fetch flowId for role request, error: %v", rr.Body.String())
	}
	var res types.ApiResponse
	json.NewDecoder(rr.Result().Body).Decode(&res)
	var getRoleIdPayload roleid.GetRoleIdPayload
	testingcommon.ExtractPayload(&res, &getRoleIdPayload)
	return getRoleIdPayload
}
func getSignature(eula string, flowId string, hexPrivateKey string) string {
	message := eula + flowId

	newMsg := fmt.Sprintf("\x19Ethereum Signed Message:\n%v%v", len(message), message)

	privateKey, err := crypto.HexToECDSA(hexPrivateKey)
	if err != nil {
		log.Fatal("HexToECDSA failed ", err)
	}

	// keccak256 hash of the data
	dataBytes := []byte(newMsg)
	hashData := crypto.Keccak256Hash(dataBytes)

	signatureBytes, err := crypto.Sign(hashData.Bytes(), privateKey)
	if err != nil {
		log.Fatal("len", err)
	}

	signature := hexutil.Encode(signatureBytes)

	return signature
}
