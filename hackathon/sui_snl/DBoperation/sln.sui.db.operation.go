package sln_sui_dboperation

import (
	"fmt"
	"time"

	dbflow "github.com/MyriadFlow/storefront-gateway/hackathon/dbFlow"
)

var db = dbflow.ConnectHackDatabase()

// Define a struct to represent your table
type SlnSui struct {
	ID              uint      `gorm:"primaryKey"`
	GameID          int       `gorm:"column:game_id"`
	CreatedAt       time.Time `gorm:"column:created_at"`
	ContractAddress string    `gorm:"column:contract_address"`
}

// func ConnectHackDatabase() *gorm.DB {
// 	// Connect to the database
// 	dbHost := os.Getenv("SUPABASE_DB_HOST")
// 	dbUsername := os.Getenv("SUPABASE_DB_USERNAME")
// 	dbPassword := os.Getenv("SUPABASE_DB_PASSWORD")
// 	dbName := os.Getenv("SUPABASE_DB_NAME")
// 	dbPort := os.Getenv("SUPABASE_DB_PORT")

// 	// Construct the connection string
// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
// 		dbHost, dbUsername, dbPassword, dbName, dbPort)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		panic("failed to connect database")
// 	}

// 	// // Auto migrate the schema
// 	// db.AutoMigrate(&SlnSui{})
// 	return db

// }

func InsertSLN_SUI() {
	// Create a new record

	newRecord := SlnSui{
		CreatedAt:       time.Now(),
		ContractAddress: "0x1234567890abcdef",
	}
	db.Create(&newRecord)

	// Check for errors
	if db.Error != nil {
		panic(db.Error)
	}

	fmt.Println("Record inserted successfully!")
}
