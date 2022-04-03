package delegateartifactcreation

type DelegateArtifactCreationRequest struct {
	CreatorAddress string `json:"creatorAddress" binding:"required,hexadecimal"`
	MetaDataHash   string `json:"metaDataHash" binding:"required"`
}

type DelegateArtifactCreationPayload struct {
	TransactionHash string `json:"transactionHash"`
}
