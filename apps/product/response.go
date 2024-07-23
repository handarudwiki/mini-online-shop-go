package product

import "time"

type ProductListResponse struct {
	ID    int    `json:"id"`
	SKU   string `json:"sku"`
	Name  string `json:"name"`
	Stock int16  `json:"stock"`
	Price int    `json:"price"`
}

func NewProductListResponse(products []ProductEntity) []ProductListResponse {
	productList := []ProductListResponse{}

	for _, product := range products {
		productList = append(productList, ProductListResponse{
			ID:    product.ID,
			Name:  product.Name,
			SKU:   product.SKU,
			Stock: product.Stock,
			Price: product.Price,
		})
	}
	return productList
}

type ProductDetailResponse struct {
	ID        int       `json:"id"`
	SKU       string    `json:"sku"`
	Name      string    `json:"name"`
	Price     int       `json:"price"`
	Stock     int       `json:"stock"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
