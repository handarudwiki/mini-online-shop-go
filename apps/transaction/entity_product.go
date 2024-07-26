package transaction

import "handarudwiki/mini-online-shop-go/infra/response"

type Product struct {
	ID    int    `db:"id" json:"id"`
	SKU   string `db:"sku" json:"sku"`
	Name  string `db:"name" json:"name"`
	Stock int    `db:"stock" json:"-"`
	Price int    `db:"price" json:"price"`
}

func (p Product) IsExists() bool {
	return p.ID != 0
}

func (p *Product) UpdateStockProduct(amount uint8) (err error) {
	if p.Stock < int(amount) {
		return response.ErrAmountGreaterThanStock
	}
	p.Stock = p.Stock - int(amount)
	return
}
