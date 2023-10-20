package highlights

type HighlightRequest struct {
	ItemId          int    `json:"itemId,omitempty"`
	ContractAddress string `json:"contractAddress,omitempty"`
	StorefrontId    string `json:"storefrontId,omitempty"`
}
