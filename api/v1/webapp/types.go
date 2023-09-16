package webapp

type WebappPayload struct {
	SubgraphId string `json:"subgraphId,omitempty"`
}

type WebappResponse struct {
	SubgraphUrl string `json:"subgraphUrl,omitempty"`
}

type Contract struct {
	ContractName    string `json:"contractName"`
	ContractAddress string `json:"contractAddress"`
	CollectionName  string `json:"collectionName"`
}
