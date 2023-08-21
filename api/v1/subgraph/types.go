package subgraph

type SubgraphPayload struct {
	Name            string `json:"name"`
	Folder          string `json:"folder"`
	NodeUrl         string `json:"nodeUrl"`
	IpfsUrl         string `json:"ipfsUrl"`
	ContractAddress string `json:"contractAddress"`
	Network         string `json:"network"`
	Protocol        string `json:"protocol"`
	Tag             string `json:"tag"`
}

type SubgraphResponse struct {
	SubgraphUrl string `json:"subgraphUrl,omitempty"`
	SubgraphId  string `json:"subgraphId,omitempty"`
}
