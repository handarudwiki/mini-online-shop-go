package product

import (
	"handarudwiki/mini-online-shop-go/infra/response"
	"time"

	"github.com/google/uuid"
)

type ProductEntity struct {
	ID         int       `db:"id"`
	SKU        string    `db:"sku"`
	Name       string    `db:"name"`
	Stock      int16     `db:"stock"`
	Price      int       `db:"price"`
	CreatedAt  time.Time `db:"created_at"`
	Updated_at time.Time `db:"updated_at"`
}

type ProductPagination struct {
	Cursor int `json:"cursor"`
	Size   int `json:"size"`
}

func NewProductPagination(req ListProductRequestPayload) ProductPagination {
	return ProductPagination{
		Cursor: req.Cursor,
		Size:   req.Size,
	}
}

func NewProductEntity(req CreateProductRequestPayload) ProductEntity {
	return ProductEntity{
		Name:       req.Name,
		Stock:      req.Stock,
		Price:      req.Price,
		SKU:        uuid.NewString(),
		CreatedAt:  time.Now(),
		Updated_at: time.Now(),
	}
}

func (p ProductEntity) Validate() (err error) {
	err = p.ValidateName()
	if err != nil {
		return
	}
	err = p.ValidateStcok()
	if err != nil {
		return
	}
	err = p.ValidatePrice()
	if err != nil {
		return
	}
	return

}

func (p ProductEntity) ValidateName() (err error) {
	if p.Name == "" {
		return response.ErrProductNameRequired
	}
	if len(p.Name) < 4 {
		return response.ErrProductNameInvalid
	}
	return
}

func (p ProductEntity) ValidateStcok() (err error) {
	if p.Stock <= 0 {
		return response.ErrSrockInvalid
	}
	return
}

func (p ProductEntity) ValidatePrice() (err error) {
	if p.Price <= 0 {
		return response.ErrPriceInvalid
	}
	return
}
