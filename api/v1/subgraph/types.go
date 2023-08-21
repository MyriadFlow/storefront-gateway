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
	Name            string `json:"name,omitempty"`
	Folder          string `json:"folder,omitempty"`
	NodeUrl         string `json:"nodeUrl,omitempty"`
	IpfsUrl         string `json:"ipfsUrl,omitempty"`
	ContractAddress string `json:"contractAddress,omitempty"`
	Network         string `json:"network,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	Tag             string `json:"tag,omitempty"`
	SubgraphUrl     string `json:"subgraphUrl,omitempty"`
	SubgraphId      string `json:"subgraphId,omitempty"`
}
