package dbinit

import (
	"log"

	"github.com/MyriadFlow/storefront-gateway/config/dbconfig"
	"github.com/MyriadFlow/storefront-gateway/models"
)

func Init() error {
	db := dbconfig.GetDb()

	err := db.AutoMigrate(&models.User{}, &models.FlowId{}, &models.Highlights{}, &models.Likes{}, &models.Wishlist{}, models.Contract{}, models.Subscription{})
	if err != nil {
		log.Fatal(err)
		return err
	}

	//Create flow id
	db.Exec(`
	DO $$ BEGIN
		CREATE TYPE flow_id_type AS ENUM (
			'AUTH'
			);
	EXCEPTION
    	WHEN duplicate_object THEN null;
	END $$;`)

	db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp";`)

	return nil
}
