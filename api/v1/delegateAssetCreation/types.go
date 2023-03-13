package delegateartifactcreation

type DelegateAssetCreationRequest struct {
	CreatorAddress string `json:"creatorAddress" binding:"required,hexadecimal"`
	MetaDataHash   string `json:"metaDataHash" binding:"required"`
}

type DelegateAssetCreationPayload struct {
	TransactionHash string `json:"transactionHash"`
}
