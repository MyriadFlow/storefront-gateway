package Org

import (
	"github.com/TheLazarusNetwork/marketplace-engine/config/dbconfig"
	"github.com/lib/pq"
)

type Org struct {
	Name               string         `gorm:"primary_key"`
	HomeTitle          string         `json:"homeTitle"`
	HomeDescription    string         `json:"homeDescription"`
	GraphUrl           string         `json:"graphurl"`
	CreatifyAddress    string         `json:"creatifyAddress"`
	MarketPlaceAddress string         `json:"marketPlaceAddress"`
	TopHighlights      pq.StringArray `gorm:"type:text[]" json:"topHighlights"`
	Trendings          pq.StringArray `gorm:"type:text[]" json:"trendings"`
	TopBids            pq.StringArray `gorm:"type:text[]" json:"topBids"`
}

func CreateOrg(org Org) error {
	return dbconfig.GetDb().Create(&org).Error
}
func GetOrgDetails() (error, Org) {
	var org Org
	err := dbconfig.GetDb().First(&org).Error
	return err, org
}
