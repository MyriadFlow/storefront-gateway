package subgraph

type SubgraphPayload struct {
	Name            string `json:"name"`
	Folder          string `json:"folder"`
	NodeUrl         string `json:"nodeUrl"`
	IpfsUrl         string `json:"ipfsUrl"`
	ContractName    string `json:"contractName"`
	ContractAddress string `json:"contractAddress"`
	Network         string `json:"network"`
	Protocol        string `json:"protocol"`
	Tag             string `json:"tag"`
}

type SubgraphResponse struct {
	Name            string `json:"name"`
	Folder          string `json:"folder"`
	NodeUrl         string `json:"nodeUrl"`
	IpfsUrl         string `json:"ipfsUrl"`
	ContractName    string `json:"contractName"`
	ContractAddress string `json:"contractAddress"`
	Network         string `json:"network"`
	Protocol        string `json:"protocol"`
	Tag             string `json:"tag"`
	SubgraphUrl     string `json:"subgraphUrl"`
	SubgraphId      string `json:"subgraphId"`
}
