package models

import (
	"time"

	"github.com/google/uuid"
)

type Storefront struct {
	Id          uuid.UUID `gorm:"type:uuid;primary_key;" json:"id"`
	Name        string    `json:"string"`
	Owner       string    `json:"owner"`
	Plan        string    `json:"plan"`
	Cost        int       `json:"cost"`
	Currency    string    `json:"currency"`
	Status      string    `json:"status"`
	Validity    time.Time `json:"validity"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	CreatedBy   string    `json:"createdBy"`
	UpdatedBy   string    `json:"updatedBy"`
	Image       string    `json:"image"`
	Headline    string    `json:"headline"`
	Description string    `json:"description"`
	Blockchain  string    `json:"blockchain"`
}
