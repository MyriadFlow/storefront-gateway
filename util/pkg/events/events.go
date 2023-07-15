package events

import (
	"fmt"
	"log"

	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	flow "github.com/MyriadFlow/storefront-gateway/generated/smartcontract/subscription"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func ListenEvent() {
	client, err := ethclient.Dial(envconfig.EnvVars.POLYGON_RPC)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(envconfig.EnvVars.SUBSCRIPTION_CONTRACT_ADDRESS)

	instance, err := flow.NewFlowSubscription(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}
	subscriptionIssuedChannel := make(chan *flow.FlowSubscriptionSubscriptionIssued)
	_, err = instance.WatchSubscriptionIssued(nil, subscriptionIssuedChannel, []common.Address{})
	if err != nil {
		log.Fatal("failed to watch subscriptionIssued, error: ", err.Error())
	}

	for e := range subscriptionIssuedChannel {
		fmt.Println(e.Owner)
	}
}
