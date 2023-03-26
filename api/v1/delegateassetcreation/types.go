package delegateassetcreation

import (
	"math/big"
)

type DelegateAssetCreationRequest struct {
	CreatorAddress           string   `json:"creatorAddress" binding:"required,hexadecimal"`
	MetaDataHash             string   `json:"metaDataHash" binding:"required"`
	RoyaltyPercentBasisPoint *big.Int `json:"royaltyPercentBasisPoint" binding:"required"`
}

type DelegateAssetCreationPayload struct {
	TransactionHash string `json:"transactionHash"`
}
