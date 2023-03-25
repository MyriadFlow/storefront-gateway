package creatorrole

import (
	"net/http"

	"github.com/MyriadFlow/storefront-gateway/api/middleware/auth/paseto"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/smartcontract"
	"github.com/MyriadFlow/storefront-gateway/config/smartcontract/rawtransaction"
	storefront_instance "github.com/MyriadFlow/storefront-gateway/config/storefront"
	storefront "github.com/MyriadFlow/storefront-gateway/generated/smartcontract/storefront"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/httphelper"
	"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	"github.com/gin-gonic/gin"
)

// ApplyRoutes applies router to gin Router
func ApplyRoutes(r *gin.RouterGroup) {
	g := r.Group("/creatorrole")
	{
		g.Use(paseto.PASETO)
		g.POST("", postGrantRole)
	}
}

func postGrantRole(c *gin.Context) {
	var req CreatorRoleRequest
	err := c.BindJSON(&req)
	if err != nil {
		httphelper.ErrResponse(c, http.StatusForbidden, "payload is invalid")
		return
	}

	walletAddress := c.GetString("walletAddress")
	client, err := smartcontract.GetClient()
	if err != nil {
		httphelper.InternalServerError(c)
		c.Abort()
		return
	}

	instance, err := storefront_instance.GetInstance(client)
	if err != nil {
		logwrapper.Errorf("failed to get instance for %v , error: %v", "STOREFRONT", err.Error())
		httphelper.InternalServerError(c)
		c.Abort()
		return
	}

	operatorRole, err := instance.STOREFRONTOPERATORROLE(nil)
	if err != nil {
		logwrapper.Errorf("Failed to get %v, error: %v", "STOREFRONTOPERATORROLE", err.Error())
		httphelper.InternalServerError(c)
		c.Abort()
		return
	}
	hasRole, err := instance.HasRole(nil, operatorRole, common.HexToAddress(walletAddress))
	if err != nil {
		logwrapper.Errorf("failed to call %v smart contract function HasRole , error: %v", "STOREFRONT", err.Error())
		httphelper.InternalServerError(c)
		c.Abort()
		return
	}

	if !hasRole {
		httphelper.ErrResponse(c, http.StatusForbidden, "only operator can access this API")
		c.Abort()
		return
	}

	roleIdBytesSlice, err := hexutil.Decode(req.RoleId)
	if err != nil {
		logwrapper.Warnf("failed to decode hex string : %v, for role for wallet address %v", req.RoleId, walletAddress)
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		return
	}
	walletAddressHex := common.HexToAddress(walletAddress)
	var roleIdBytes [32]byte
	copy(roleIdBytes[:], roleIdBytesSlice)

	tx, err := rawtransaction.SendRawTransaction(storefront.StorefrontABI, "grantRole", roleIdBytes, walletAddressHex)

	if err != nil {
		httphelper.ErrResponse(c, http.StatusInternalServerError, "Unexpected error occured")
		logwrapper.Warnf("failed to grant role to user with walletaddress %v, error: %v", walletAddress, err.Error())
		return
	}
	transactionHash := tx.Hash().String()
	logwrapper.Infof("trasaction hash is %v", transactionHash)
	payload := CreatorRolePayload{
		TransactionHash: transactionHash,
	}
	httphelper.SuccessResponse(c, "role grant transaction has been broadcasted", payload)

}

func getRoleByFlowId(flowId string) (models.Role, error) {
	db := dbconfig.GetDb()
	var flowIdRecord models.FlowId
	err := db.Model(&models.FlowId{}).Where("flow_id = ?", flowId).First(&flowIdRecord).Error
	if err != nil {
		return models.Role{}, err
	}

	var role models.Role
	err = db.Model(&models.Role{}).Where("role_id = ?", flowIdRecord.RelatedRoleId).First(&role).Error
	if err != nil {
		return models.Role{}, err
	}
	return role, nil
}
