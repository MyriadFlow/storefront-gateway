package models

type Likes struct {
	ItemId          int    `json:"itemId,omitempty"`
	ContractAddress string `json:"contractAddress,omitempty"`
	WalletAddress   string `json:"walletAddress"`
}
