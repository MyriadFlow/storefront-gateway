package storefront

import (
	"github.com/google/uuid"
)

type StorefrontRequest struct {
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	CreatedBy   string `json:"createdBy"`
	UpdatedBy   string `json:"updatedBy"`
	Image       string `json:"image"`
	Headline    string `json:"headline"`
	Description string `json:"description"`
	Blockchain  string `json:"blockchain"`
}

type UpdateStorefrontRequest struct {
	Id                  uuid.UUID `json:"id"`
	Name                string    `json:"name"`
	UpdatedBy           string    `json:"updatedBy"`
	Headline            string    `json:"headline"`
	Description         string    `json:"description"`
	Image               string    `json:"image"`
	ProfileImage        string    `json:"profileImage,omitempty"`
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
}

type DeployStorefrontRequest struct {
	Name                string `json:"name,omitempty"`
	StorefrontId        string `json:"storefrontId,omitempty"`
	Network             string `json:"network,omitempty"`
	Protocol            string `json:"protocol,omitempty"`
	Tag                 string `json:"tag,omitempty"`
	StorefrontName      string `json:"storefrontName,omitempty"`
	Headline            string `json:"headline,omitempty"`
	Description         string `json:"description,omitempty"`
	ProfileImage        string `json:"profileImage,omitempty"`
	CoverImage          string `json:"coverImage,omitempty"`
	AssetName           string `json:"assetName,omitempty"`
	AssetDescription    string `json:"assetDescription,omitempty"`
	PersonalInformation string `json:"personalTagline,omitempty"`
	PersonalDescription string `json:"personalDescription,omitempty"`
	RelevantImage       string `json:"relevantImage,omitempty"`
	MailId              string `json:"mailId,omitempty"`
	Twitter             string `json:"twitter,omitempty"`
	Discord             string `json:"discord,omitempty"`
	Instagram           string `json:"instagram,omitempty"`
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
