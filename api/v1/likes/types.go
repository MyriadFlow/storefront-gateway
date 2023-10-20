package likes

type LikeReqeust struct {
	ItemId          int    `json:"itemId,omitempty"`
	ContractAddress string `json:"contractAddress,omitempty"`
}
