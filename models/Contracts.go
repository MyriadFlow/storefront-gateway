package models

type Contract struct {
	ContractName    string `json:"contractName"`
	ContractAddress string `json:"contractAddress"`
	ChainId         int    `json:"chainId"`
	Verified        bool   `json:"verified"`
	SubscriptionId  string `json:"subscriptionId"`
	ContractId      string `json:"contractId"`
}