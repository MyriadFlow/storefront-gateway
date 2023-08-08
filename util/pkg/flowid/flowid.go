package flowid

import (
	"fmt"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func GenerateFlowId(walletAddress string, flowIdType models.FlowIdType, relatedRoleId string) (string, error) {
	db := dbconfig.GetDb()
	flowId := uuid.NewString()
	var update bool
	update = true

	findResult := db.Model(&models.User{}).Find(&models.User{}, &models.User{WalletAddress: walletAddress})

	if err := findResult.Error; err != nil {
		err = fmt.Errorf("while finding user error occured, %s", err)
		logrus.Error(err)
		return "", err
	}

	rowsAffected := findResult.RowsAffected
	if rowsAffected == 0 {
		update = false
	}
	if update {
		// User exist so update
		findResult.Update("flow_ids", gorm.Expr("array_cat(tags, ?)", &models.FlowId{FlowIdType: flowIdType, FlowId: flowId, WalletAddress: walletAddress}))
	} else {
		// User doesn't exist so create
		id := uuid.New()
		newUser := &models.User{
			UserID:        id,
			WalletAddress: walletAddress,
			FlowIds: []models.FlowId{{
				FlowIdType: flowIdType, FlowId: flowId,
			}},
		}
		if err := db.Create(newUser).Error; err != nil {
			return "", err
		}

	}

	return flowId, nil
}
