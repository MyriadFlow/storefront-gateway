package storefront

import (
	"github.com/google/uuid"
)

type StorefrontRequest struct {
	Name        string `json:"name"`
	Owner       string `json:"owner"`
	Plan        string `json:"plan"`
	Cost        int    `json:"cost"`
	Currency    string `json:"currency"`
	CreatedBy   string `json:"createdBy"`
	UpdatedBy   string `json:"updatedBy"`
	Image       string `json:"image"`
	Headline    string `json:"headline"`
	Description string `json:"description"`
	Blockchain  string `json:"blockchain"`
}

type UpdateStorefrontRequest struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Status      string    `json:"status"`
	Validity    string    `json:"validity"`
	UpdatedBy   string    `json:"updatedBy"`
	Headline    string    `json:"headline"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
}
