package models

type Contract struct {
	ContractName    string `json:"contractName"`
	ContractAddress string `json:"contractAddress"`
	WalletAddress   string `json:"walletAddress"`
	ChainId         int    `json:"chainId"`
	Verified        bool   `json:"verified"`
	StorefrontId    string `json:"storefrontId"`
	ContractId      string `json:"contractId"`
}
