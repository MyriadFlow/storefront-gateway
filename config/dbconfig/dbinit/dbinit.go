package dbinit

import (
	"log"
	// "errors"
	// "encoding/json"
	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	//"github.com/MyriadFlow/storefront-gateway/config/storefront"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/models/Org"
	//"github.com/MyriadFlow/storefront-gateway/util/pkg/logwrapper"
	//"github.com/ethereum/go-ethereum/common/hexutil"
)

func Init() error {
	db := dbconfig.GetDb()
	//err := db.AutoMigrate(&models.User{}, &models.FlowId{}, &models.Role{}, &Org.Org{},&models.Product{})
	err := db.AutoMigrate(&models.User{}, &models.FlowId{}, &models.Role{}, &models.Marketplace{})
	if err != nil {
		log.Fatal(err)
		return err

	}
	//create org table
	db.Exec(`create table if not exists orgs (
		name text,
		home_title text,
		home_description text,
		graphql_marketpace text,
		store_front_address text,
		market_place_address text,
		footer text,
		top_highlights text[],
		trendings text[],
		contacts jsonb,
		unique (name)
		)`)

	contacts := Org.OrgContacts{
		DiscordId:   "-",
		InstagramId: "-",
		TelegramId:  "-",
		TwitterId:   "@0xMyriadFlow",
	}

	err = Org.CreateOrg(
		Org.Org{
			Name:               envconfig.EnvVars.ORG_NAME,
			HomeTitle:          envconfig.EnvVars.HOME_TITLE,
			HomeDescription:    envconfig.EnvVars.HOME_DESCRIPTION,
			GraphUrl:           envconfig.EnvVars.GRAPHQL_MARKETPLACE,
			MarketPlaceAddress: envconfig.EnvVars.MARKETPLACE_CONTRACT_ADDRESS,
			StoreFrontAddress:  envconfig.EnvVars.STOREFRONT_CONTRACT_ADDRESS,
			Footer:             envconfig.EnvVars.FOOTER,
			TopHighlights:      envconfig.EnvVars.TOP_HIGHLIGHTS,
			Trendings:          envconfig.EnvVars.TRENDINGS,
			Contacts:           contacts,
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

	// creatorRoleId, err := storefront.GetRole(storefront.CREATOR_ROLE)
	// if err != nil {
	// 	logwrapper.Fatal(err)
	// }

	// creatorEula := envconfig.EnvVars.CREATOR_EULA

	// // TODO: create role only if they does not exist
	// rolesToBeAdded := []models.Role{
	// 	{Name: "Creator Role", RoleId: hexutil.Encode(creatorRoleId[:]), Eula: creatorEula}}
	// for _, role := range rolesToBeAdded {
	// 	if err := db.Model(&models.Role{}).FirstOrCreate(&role).Error; err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	//add dummy marketplace details
	demoProduct := []models.Marketplace{
		{ItemId: 1, NFT_Contract_Address: "nftcontractaddr1", TokenId: "tokenid1", MetaDataURI: "metadatauri1"},
		{ItemId: 2, NFT_Contract_Address: "nftcontractaddr2", TokenId: "tokenid2", MetaDataURI: "metadatauri2"},
	}
	for _, product := range demoProduct {
		if err := db.Model(&models.Marketplace{}).Create(&product).Error; err != nil {
			log.Fatal(err)
		}
	}
	return nil
}
