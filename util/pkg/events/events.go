package events

import (
	"fmt"
	"log"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	flow "github.com/MyriadFlow/storefront-gateway/generated/smartcontract/subscription"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

func ListenEvent() {
	client, err := ethclient.Dial(envconfig.EnvVars.POLYGON_RPC_WSS)
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
	} else {
		fmt.Println("watching subscriptionIssued event")
	}

	for e := range subscriptionIssuedChannel {
		db := dbconfig.GetDb()
		result := db.Model(&models.User{}).Where("wallet_address = ?", e.Owner.String()).Update("plan", "pro")
		if result.Error != nil {
			logrus.Error("Unable to subscribe. Error: ")
			return
		}
	}
}
