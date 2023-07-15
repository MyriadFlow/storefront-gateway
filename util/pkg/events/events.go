package events

import (
	"context"
	"fmt"
	"log"

	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ListenEvent() {
	client, err := ethclient.Dial(envconfig.EnvVars.POLYGON_RPC)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(envconfig.EnvVars.SUBSCRIPTION_CONTRACT_ADDRESS)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}
	logs := make(chan types.Log)

	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog)
		}
	}
}
