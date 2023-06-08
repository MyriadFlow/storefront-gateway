package models

type Highlights struct {
	ItemId           int    `gorm:"primary_key" json:"itemId"`
	Contract_Address string `json:"contractAddress"`
	TokenId          string `json:"tokenId,omitempty"`
	Metadata         string `json:"metadata,omitempty"`
}
