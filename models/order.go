package models

import (
	"time"

	"github.com/google/uuid"
)

type Order struct {
	OrderID     uint64     `json:"order_id"`
	UserID      uuid.UUID  `json:"user_id"`
	LineItems   []LineItem `json:"line_items"`
	CreateAt    *time.Time `json:"create_at"`
	ShipperdAt  *time.Time `json:"shipperd_at"`
	CompletedAt *time.Time `json:"completed_at"`
}
type LineItem struct {
	ItemID  uuid.UUID `json:"item_id"`
	Quntity uint64    `json:"quantity"`
	Price   uint64    `json:"price"`
}
