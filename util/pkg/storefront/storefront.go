package storefrontUtil

import (
	"time"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/google/uuid"
)

func CreateStorefront(name string, owner string, walletAdress string, createdBy string, updatedBy string, image string, headline string, description string, blockchain string) error {
	db := dbconfig.GetDb()
	storefront := models.Storefront{
		Id:           uuid.New(),
		Name:         name,
		Owner:        owner,
		WalletAdress: walletAdress,
		Status:       "active",
		Validity:     time.Now().AddDate(1, 0, 0),
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
		CreatedBy:    createdBy,
		UpdatedBy:    updatedBy,
		Image:        image,
		Headline:     headline,
		Description:  description,
		Blockchain:   blockchain,
	}
	result := db.Create(&storefront)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
