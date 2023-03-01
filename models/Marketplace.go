package models

type Marketplace struct {
	ItemId              int   `gorm:"primary_key" json:"itemId"`
	NFT_Contract_Address     string   `json:"nft-contract-address"`
	TokenId            string   `json:"tokenId,omitempty"`
	MetaDataURI string   `json:"metaDataURI,omitempty"`
}
