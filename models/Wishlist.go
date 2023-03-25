package models

type Wishlist struct {
	ItemId           int   `json:"tokenId,omitempty"`
	UserWalletAddress  string	`json:"walletAddress"`
}