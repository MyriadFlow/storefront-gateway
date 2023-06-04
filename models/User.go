package models

import "github.com/google/uuid"

type User struct {
	UserID         uuid.UUID `gorm:"type:uuid;primaryKey;default:gen_random_uuid()"`
	Email          string    `json:"email",omitempty`
	Name           string    `json:"name"`
	WalletAddress  string    ` json:"wallet_address"`
	FlowIds        []FlowId  `gorm:"foreignkey:WalletAddress" json:"-"`
	ProfilePicture string    `json:"profilePictureUrl"`
	Country        string    `json:"country"`
	FacebookId     string    `json:"facebook_id"`
	InstagramId    string    `json:"instagram_id"`
	TwitterId      string    `json:"twitter_id"`
	DiscordId      string    `json:"discord_id"`
	TelegramId     string    `json:"telegram_id"`
	UserType       string    `json:"user_type,omitempty"`
}

// gorm:"type:varchar(20);column:id;next:uuid
