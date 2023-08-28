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
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Validity    string    `json:"validity"`
	UpdatedBy   string    `json:"updatedBy"`
	Headline    string    `json:"headline"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
}

type DeployStorefrontRequest struct {
	Name         string `json:"name"`
	NodeUrl      string `json:"nodeUrl"`
	StorefrontId string `json:"storefrontId"`
	Network      string `json:"network"`
	Protocol     string `json:"protocol"`
	Tag          string `json:"tag"`
	NodectlUrl   string `json:"nodectlUrl"`
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
	Name    string `json:"name"`
	Address string `json:"address"`
}

type NodectlRequest struct {
	StorefrontName string `json:"storefrontName"`
	StorefrontId   string `json:"storefrontId"`
}
type NodectlResponse struct {
	StorefrontUrl string `json:"storefrontUrl"`
}
