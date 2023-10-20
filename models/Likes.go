package models

type Likes struct {
	ItemId            int    `json:"itemId,omitempty"`
	ContractAddress   string `json:"contractAddress,omitempty"`
	UserWalletAddress string `json:"walletAddress"`
}
