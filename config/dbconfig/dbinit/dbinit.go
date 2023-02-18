package dbinit

import (
	"log"

	"github.com/MyriadFlow/storefront_gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront_gateway/config/envconfig"
	"github.com/MyriadFlow/storefront_gateway/config/storefront"
	"github.com/MyriadFlow/storefront_gateway/models"
	"github.com/MyriadFlow/storefront_gateway/models/Org"
	"github.com/MyriadFlow/storefront_gateway/util/pkg/logwrapper"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

func Init() error {
	db := dbconfig.GetDb()
	err := db.AutoMigrate(&models.User{}, &models.FlowId{}, &models.Role{}, &Org.Org{},&models.Product{})
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = Org.CreateOrg(
		Org.Org{
			Name:               envconfig.EnvVars.ORG_NAME,
			HomeTitle:          envconfig.EnvVars.HOME_TITLE,
			HomeDescription:    envconfig.EnvVars.HOME_DESCRIPTION,
			GraphUrl:           envconfig.EnvVars.GRAPH_URL,
			MarketPlaceAddress: envconfig.EnvVars.MARKETPLACE_CONTRACT_ADDRESS,
			StoreFrontAddress:  envconfig.EnvVars.STOREFRONT_CONTRACT_ADDRESS,
			Footer:             envconfig.EnvVars.FOOTER,
			TopHighlights:      envconfig.EnvVars.TOP_HIGHLIGHTS,
			Trendings:          envconfig.EnvVars.TRENDINGS,
			TopBids:            envconfig.EnvVars.TOP_BIDS,
		},
	)

	if err != nil {
		return err
	}

	//Create user_roles table
	db.Exec(`create table if not exists user_roles (
			wallet_address text,
			role_id text,
			unique (wallet_address,role_id)
			)`)

	//Create flow id
	db.Exec(`
	DO $$ BEGIN
		CREATE TYPE flow_id_type AS ENUM (
			'AUTH',
			'ROLE');
	EXCEPTION
    	WHEN duplicate_object THEN null;
	END $$;`)

	creatorRoleId, err := storefront.GetRole(storefront.CREATOR_ROLE)
	if err != nil {
		logwrapper.Fatal(err)
	}

	creatorEula := envconfig.EnvVars.CREATOR_EULA

	// TODO: create role only if they does not exist
	rolesToBeAdded := []models.Role{
		{Name: "Creator Role", RoleId: hexutil.Encode(creatorRoleId[:]), Eula: creatorEula}}
	for _, role := range rolesToBeAdded {
		if err := db.Model(&models.Role{}).FirstOrCreate(&role).Error; err != nil {
			log.Fatal(err)
		}
	}

	//add dummy products details
	demoProduct := []models.Product{
		{ItemId:1,NFT_Contract_Address:"nftcontractaddr1",TokenId:"tokenid1",MetaDataURI:"metadatauri1"},
		{ItemId:2,NFT_Contract_Address:"nftcontractaddr2",TokenId:"tokenid2",MetaDataURI:"metadatauri2"},
	}
	for _, product := range demoProduct {
		if err := db.Model(&models.Product{}).Create(&product).Error; err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
