package profile

type GetProfilePayload struct {
	Name              string `json:"name,omitempty"`
	WalletAddress     string `json:"walletAddress"`
	ProfilePictureUrl string `json:"profilePictureUrl,omitempty"`
	Country           string `json:"country,omitempty"`
	FacebookId           string   `json:"facebook_id,omitempty"`
	InstagramId           string   `json:"instagram_id,omitempty"`
}
