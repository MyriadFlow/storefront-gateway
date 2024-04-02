package delegateassetcreation

import (
	"math/big"
)

type DelegateAssetCreationRequest struct {
	CreatorAddress           string   `json:"creatorAddress" binding:"required,hexadecimal"`
	MetaDataHash             string   `json:"metaDataHash" binding:"required"`
	RoyaltyPercentBasisPoint *big.Int `json:"royaltyPercentBasisPoint" binding:"required"`
	ContractAddress          string   `json:"contractAddress" binding:"required"`
}

type DelegateAssetCreationPayload struct {
	TransactionHash string `json:"transactionHash"`
}

type AssetStoreRequest struct {
	MetaDataHash             string `json:"metaDataHash" binding:"required"`
	RoyaltyPercentBasisPoint int    `json:"royaltyPercentBasisPoint" binding:"required"`
	StorefrontId             string `json:"storefrontId" binding:"required"`
	ContractAddress          string `json:"contractAddress" binding:"required"`
}

type AssetCreateRequest struct {
	MetaDataHash             string   `json:"metaDataHash" binding:"required"`
	RoyaltyPercentBasisPoint *big.Int `json:"royaltyPercentBasisPoint" binding:"required"`
	ContractAddress          string   `json:"contractAddress" binding:"required"`
}

type AssetTransferRequest struct {
	FromAddress     string   `json:"fromAddress" binding:"required,hexadecimal"`
	ToAddress       string   `json:"toAddress" binding:"required,hexadecimal"`
	TokenId         *big.Int `json:"tokenId" binding:"required"`
	ContractAddress string   `json:"contractAddress" binding:"required"`
}
