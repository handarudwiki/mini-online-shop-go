package product

import (
	"context"
	response "handarudwiki/mini-online-shop-go"
)

type Repository interface {
	CreateProduct(ctx context.Context, product ProductEntity) (err error)
	GetAllProducts(ctx context.Context, product ProductPagination) (products []ProductEntity, err error)
	GetProductBySKU(ctx context.Context, sku string) (product ProductEntity, err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}
func (s service) CreateProduct(ctx context.Context, req CreateProductRequestPayload) (err error) {
	productEntity := NewProductEntity(req)

	err = productEntity.Validate()

	if err != nil {
		return
	}

	err = s.repo.CreateProduct(ctx, productEntity)

	if err != nil {
		return
	}
	return
}

func (s service) ListProducts(ctx context.Context, req ListProductRequestPayload) (products []ProductEntity, err error) {
	productPagination := NewProductPagination(req)
	products, err = s.repo.GetAllProducts(ctx, productPagination)
	if err != nil {
		if err == response.ErrNotFound {
			return []ProductEntity{}, nil
		}
		return
	}

	if len(products) == 0 {
		return []ProductEntity{}, nil
	}
	return
}

func (s service) DetailProduct(ctx context.Context, sku string) (product ProductEntity, err error) {
	product, err = s.repo.GetProductBySKU(ctx, sku)

	if err != nil {
		return
	}

	return
}
