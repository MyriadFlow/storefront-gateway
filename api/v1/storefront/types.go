package storefront

import (
	"github.com/google/uuid"
)

type StorefrontRequest struct {
	Name        string `json:"name"`
	Blockchain  string `json:"blockchain"`
	Headline    string `json:"headline"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Network     string `json:"network"`
}

type UpdateStorefrontRequest struct {
	Id                  uuid.UUID `json:"id"`
	Name                string    `json:"name"`
	Headline            string    `json:"headline"`
	Description         string    `json:"description"`
	ProfileImage        string    `json:"profileImage"`
	StorefrontImage     string    `json:"storefrontImage"`
	PersonalTagline     string    `json:"personalTagline"`
	PersonalDescription string    `json:"personalDescription"`
	RelevantImage       string    `json:"relevantImage"`
	MailId              string    `json:"mailId"`
	Twitter             string    `json:"twitter"`
	Discord             string    `json:"discord"`
	Instagram           string    `json:"instagram"`
}

type DeployStorefrontRequest struct {
	Id                  uuid.UUID `json:"id"`
	Name                string    `json:"name"`
	Headline            string    `json:"headline"`
	Description         string    `json:"description"`
	ProfileImage        string    `json:"profileImage"`
	StorefrontImage     string    `json:"storefrontImage"`
	PersonalTagline     string    `json:"personalTagline"`
	PersonalDescription string    `json:"personalDescription"`
	RelevantImage       string    `json:"relevantImage"`
	MailId              string    `json:"mailId"`
	Twitter             string    `json:"twitter"`
	Discord             string    `json:"discord"`
	Instagram           string    `json:"instagram"`
	Network             string    `json:"network,omitempty"`
	Protocol            string    `json:"protocol,omitempty"`
	Tag                 string    `json:"tag,omitempty"`
}

type GraphRequest struct {
	Name      string     `json:"name"`
	Folder    string     `json:"folder"`
	NodeURL   string     `json:"nodeUrl"`
	IpfsURL   string     `json:"ipfsUrl"`
	Contracts []Contract `json:"contracts"`
	Network   string     `json:"network"`
	Protocol  string     `json:"protocol"`
	Tag       string     `json:"tag"`
}

type Contract struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	BlockNumber int    `json:"blockNumber"`
}

type NodectlRequest struct {
	StorefrontName string `json:"storefrontName"`
	StorefrontId   string `json:"storefrontId"`
}
type NodectlResponse struct {
	StorefrontUrl string `json:"storefrontUrl"`
}
