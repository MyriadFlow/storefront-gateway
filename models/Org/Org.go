package Org

import (
	"errors"

	"github.com/TheLazarusNetwork/marketplace-engine/config/dbconfig"
	"github.com/TheLazarusNetwork/marketplace-engine/config/envconfig"
	"github.com/lib/pq"
)

type Org struct {
	Name               string         `gorm:"primary_key" json:"name"`
	HomeTitle          string         `json:"homeTitle,omitempty"`
	HomeDescription    string         `json:"homeDescription,omitempty"`
	GraphUrl           string         `json:"graphurl,omitempty"`
	CreatifyAddress    string         `json:"creatifyAddress,omitempty"`
	MarketPlaceAddress string         `json:"marketPlaceAddress,omitempty"`
	Footer             string         `json:"footer,omitempty"`
	TopHighlights      pq.StringArray `gorm:"type:text[]" json:"topHighlights,omitempty"`
	Trendings          pq.StringArray `gorm:"type:text[]" json:"trendings,omitempty"`
	TopBids            pq.StringArray `gorm:"type:text[]" json:"topBids,omitempty"`
}

func CreateOrg(org Org) error {
	return dbconfig.GetDb().Create(&org).Error
}

func UpdateOrg(org Org) error {
	// Make sure org name is same as used while initiating
	org.Name = envconfig.EnvVars.ORG_NAME
	result := dbconfig.GetDb().Model(&org).Where("name = ?", org.Name).Updates(&org)
	if err := result.Error; err != nil {
		return err
	}
	if result.RowsAffected == 0 {
		err := errors.New("record not found")
		return err
	}
	return nil
}
func GetOrgDetails() (Org, error) {
	var org Org
	err := dbconfig.GetDb().First(&org).Error
	return org, err
}
