package launchpad

type reqBody struct {
	ContractName      string         `json:"contractName"`
	ConstructorParams map[string]any `json:"constructorParams"`
	Network           string         `json:"network"`
	StorefrontId      string         `json:"storefrontId"`
	CollectionName    string         `json:"collectionName"`
	Thumbnail         string         `json:"thumbnail"`
	CoverImage        string         `json:"coverImage"`
}
type resBody struct {
	ChainId         int    `json:"chainId"`
	ContractAddress string `json:"contractAddress"`
	Verified        bool   `json:"verified"`
	BlockNumber     int    `json:"blockNumber"`
}
type data struct {
	ContractName      string         `json:"contractName"`
	ConstructorParams map[string]any `json:"constructorParams"`
}
type contractReqBody struct {
	Data    data   `json:"data"`
	Network string `json:"network"`
}
