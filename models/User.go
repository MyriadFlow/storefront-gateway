package models

type User struct {
	Name              string   `json:"name,omitempty"`
	WalletAddress     string   `gorm:"primary_key" json:"walletAddress"`
	FlowIds           []FlowId `gorm:"foreignkey:WalletAddress" json:"-"`
	ProfilePictureUrl string   `json:"profilePictureUrl,omitempty"`
	Country           string   `json:"country,omitempty"`
	FacebookId           string   `json:"facebook_id,omitempty"`
	InstagramId           string   `json:"instagram_id,omitempty"`
	TwitterId           string   `json:"twitter_id,omitempty"`
	DiscordId           string   `json:"discord_id,omitempty"`
	TelegramId           string   `json:"telegram_id,omitempty"`
}
