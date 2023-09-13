package models

import (
	"time"

	"github.com/google/uuid"
)

type Storefront struct {
	Id                  uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name                string    `json:"name"`
	Owner               string    `json:"owner"`
	WalletAddress       string    `json:"walletAddress"`
	Status              string    `json:"status"`
	Validity            time.Time `json:"validity"`
	CreatedAt           time.Time `json:"createdAt"`
	UpdatedAt           time.Time `json:"updatedAt"`
	CreatedBy           string    `json:"createdBy"`
	UpdatedBy           string    `json:"updatedBy"`
	Image               string    `json:"image"`
	Headline            string    `json:"headline"`
	Description         string    `json:"description"`
	Blockchain          string    `json:"blockchain"`
	ProfileImage        string    `json:"Profileimage,omitempty"`
	CoverImage          string    `json:"coverImage,omitempty"`
	AssetName           string    `json:"assetName,omitempty"`
	AssetDescription    string    `json:"assetDescription,omitempty"`
	PersonalInformation string    `json:"personalTagline,omitempty"`
	PersonalDescription string    `json:"personalDescription,omitempty"`
	RelevantImage       string    `json:"relevantImage,omitempty"`
	MailId              string    `json:"mailId,omitempty"`
	Twitter             string    `json:"twitter,omitempty"`
	Discord             string    `json:"discord,omitempty"`
	Instagram           string    `json:"instagram,omitempty"`
	WebappUrl           string    `json:"webappUrl,omitempty"`
	SubgraphUrl         string    `json:"subgraphUrl,omitempty"`
}
