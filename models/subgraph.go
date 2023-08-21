package models

type Subgraph struct {
	SubgraphId      string `json:"subgraphId,omitempty"`
	Name            string `json:"name,omitempty"`
	Folder          string `json:"folder,omitempty"`
	NodeUrl         string `json:"nodeUrl,omitempty"`
	IpfsUrl         string `json:"ipfsUrl,omitempty"`
	ContractAddress string `json:"contractAddress,omitempty"`
	Network         string `json:"network,omitempty"`
	Protocol        string `json:"protocol,omitempty"`
	Tag             string `json:"tag,omitempty"`
	SubgraphUrl     string `json:"subgraphUrl,omitempty"`
	WalletAddress   string `json:"walletAddress,omitempty"`
	StorefrontId    string `json:"storefrontId,omitempty"`
}
