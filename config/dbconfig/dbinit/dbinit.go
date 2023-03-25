package dbinit

import (
	"log"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/config/envconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
	"github.com/MyriadFlow/storefront-gateway/models/Org"
)

func Init() error {
	db := dbconfig.GetDb()

	err := db.AutoMigrate(&models.User{}, &models.FlowId{}, &models.Highlights{}, &Org.Org{}, &models.Likes{}, &models.Wishlist{})
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

	err = Org.CreateOrg(
		Org.Org{
			Name:               envconfig.EnvVars.ORG_NAME,
			HomeTitle:          envconfig.EnvVars.HOME_TITLE,
			HomeDescription:    envconfig.EnvVars.HOME_DESCRIPTION,
			GraphqlMarketplace: envconfig.EnvVars.GRAPHQL_MARKETPLACE,
			MarketPlaceAddress: envconfig.EnvVars.MARKETPLACE_CONTRACT_ADDRESS,
			StoreFrontAddress:  envconfig.EnvVars.STOREFRONT_CONTRACT_ADDRESS,
			Footer:             envconfig.EnvVars.FOOTER,
			Contacts:           Org.OrgContacts{},
		},
	)

	if err != nil {
		return err
	}

	//Create flow id
	db.Exec(`
	DO $$ BEGIN
		CREATE TYPE flow_id_type AS ENUM (
			'AUTH',
			'ROLE');
	EXCEPTION
    	WHEN duplicate_object THEN null;
	END $$;`)

	return nil
}
