package Org

import (
	"errors"
	"encoding/json"
	"github.com/MyriadFlow/storefront_gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront_gateway/config/envconfig"
	"github.com/lib/pq"
	"database/sql/driver"
)

type Contacts struct {	
	DiscordId	string		`json:"discordId,omitempty"`
	InstagramId	string		`json:"instagramId,omitempty"`
	TelegramId	string		`json:"telegramId,omitempty"`
	TwitterId	string		`json:"twitterId,omitempty"`
}

type Org struct {
	Name               string         `gorm:"primary_key" json:"name"`
	HomeTitle          string         `json:"homeTitle,omitempty"`
	HomeDescription    string         `json:"homeDescription,omitempty"`
	GraphUrl           string         `json:"graphurl,omitempty"`
	StoreFrontAddress  string         `json:"storeFrontAddress,omitempty"`
	MarketPlaceAddress string         `json:"marketplaceAddress,omitempty"`
	Footer             string         `json:"footer,omitempty"`
	TopHighlights      pq.StringArray `gorm:"type:text[]" json:"topHighlights,omitempty"`
	Trendings          pq.StringArray `gorm:"type:text[]" json:"trendings,omitempty"`

	Contact				Contacts		`json:"contact_json_object,omitempty"`      
}
func (contact Contacts) Value() (driver.Value, error) {
	return json.Marshal(contact)
}

func (contact *Contacts) Scan(value interface{}) error {
	parse, ok := value.([]byte)
    if !ok {
        return errors.New("type assertion to []byte failed")
    }
    return json.Unmarshal(parse, &contact)
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
