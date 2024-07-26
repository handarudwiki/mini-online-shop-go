package transaction

import "time"

type TransactionHistoryResponse struct {
	ID           int       `json:"id"`
	UserPublicID string    `json:"user_public_id"`
	ProductID    uint      `json:"product_id"`
	ProductPrice uint      `json:"product_price"`
	Amount       uint8     `json:"amount"`
	Subtotal     uint      `json:"subtotal"`
	PlatformFee  uint      `json:"platform_fee"`
	GrandTotal   uint      `json:"grand_total"`
	Status       string    `json:"status"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
	Product      Product   `json:"product"`
}
