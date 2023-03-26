package models

type Likes struct {
	ItemId           int   `json:"tokenId,omitempty"`
	UserWalletAddress  string	`json:"walletAddress"`
}