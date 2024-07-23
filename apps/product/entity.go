package product

import "time"

type ProductEntity struct {
	ID         int       `db:"id"`
	SKU        string    `db:"sku"`
	Name       string    `db:"string"`
	Stock      int16     `db:"stock"`
	Price      int       `db:"price"`
	CreatedAt  time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}
