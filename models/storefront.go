package models

import (
	"time"

	"github.com/google/uuid"
)

type Storefront struct {
	Id                    uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name                  string    `json:"name"`
	WalletAddress         string    `json:"walletAddress"`
	CreatedAt             time.Time `json:"createdAt"`
	UpdatedAt             time.Time `json:"updatedAt"`
	Image                 string    `json:"Image"`
	Headline              string    `json:"headline"`
	Description           string    `json:"description"`
	Blockchain            string    `json:"blockchain"`
	Network               string    `json:"network"`
	StorefrontName        string    `json:"storefrontName"`
	StorefrontDescription string    `json:"storefrontDescription"`
	StorefrontHeadline    string    `json:"storefrontHeadline"`
	ProfileImage          string    `json:"profileimage"`
	StorefrontImage       string    `json:"storefrontImage"`
	PersonalTagline       string    `json:"personalTagline"`
	PersonalDescription   string    `json:"personalDescription"`
	RelevantImage         string    `json:"relevantImage"`
	MailId                string    `json:"mailId"`
	Twitter               string    `json:"twitter"`
	Discord               string    `json:"discord"`
	Instagram             string    `json:"instagram"`
	WebappUrl             string    `json:"webappUrl"`
	SubgraphUrl           string    `json:"subgraphUrl"`
	Deployed              bool      `json:"deployed"`
}
