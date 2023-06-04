package storefront

import (
	"github.com/MyriadFlow/storefront-gateway/generated/smartcontract/storefront"
	"github.com/ethereum/go-ethereum/ethclient"
)

var instance *storefront.Storefront

var (
// errEnvVariableNotDefined = errors.New("environment variable STOREFRONT_CONTRACT_ADDRESS is required")
)

func GetInstance(client *ethclient.Client) (*storefront.Storefront, error) {
	if instance != nil {
		return instance, nil
	}
	// envContractAddress := envconfig.EnvVars.STOREFRONT_CONTRACT_ADDRESS
	/* if envContractAddress == "" {
		logwrapper.Errorf("environment variable %v is required", "STOREFRONT_CONTRACT_ADDRESS")
		return nil, errEnvVariableNotDefined
	}
	addr := common.HexToAddress(envContractAddress)
	var err error
	instance, err = storefront.NewStorefront(addr, client)
	if err != nil {
		logwrapper.Errorf("failed to load storefront contract at address %v, error: %v", envContractAddress, err.Error())
		return nil, err
	} */
	return instance, nil
}
