package subscription

import (
	"time"

	"github.com/google/uuid"
)

type subscriptionRequest struct {
	Name      string `json:"string"`
	Owner     string `json:"owner"`
	Plan      string `json:"plan"`
	Cost      int    `json:"cost"`
	Currency  string `json:"currency"`
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
}

type UpdateSubscriptionRequest struct {
	Id        uuid.UUID `json:"id"`
	Status    string    `json:"status"`
	Validity  time.Time `json:"validity"`
	UpdatedBy string    `json:"updatedBy"`
}
