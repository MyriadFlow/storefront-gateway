package models

type Likes struct {
	NFT_Contract_Address     string   `json:"nft-contract-address"`
	TokenId            string   `json:"tokenId,omitempty"`
	UserWalletAddress  string	`json:"walletAddress"`
}