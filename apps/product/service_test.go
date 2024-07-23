package product

import (
	"context"
	"handarudwiki/mini-online-shop-go/external/database"
	"handarudwiki/mini-online-shop-go/infra/response"
	"handarudwiki/mini-online-shop-go/internal/config"
	"testing"

	"github.com/stretchr/testify/require"
)

var svc service

func init() {
	filename := "../../cmd/api/config.yaml"
	err := config.LoadConfig(filename)
	if err != nil {
		panic(err)
	}
	db, err := database.ConnectPostgres(config.Cfg.DB)
	if err != nil {
		panic(err)
	}
	repo := newRepository(db)
	svc = newService(repo)
}

func TestCreateProductSuccess(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "macbook m1",
		Stock: 15,
		Price: 1000,
	}

	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)
}

func TestCreateProductFail(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "",
		Stock: 10,
		Price: 100,
	}

	err := svc.CreateProduct(context.Background(), req)
	require.NotNil(t, err)
	require.Equal(t, response.ErrProductNameRequired, err)
}

func TestListProductSuccess(t *testing.T) {
	req := ListProductRequestPayload{
		Cursor: 0,
		Size:   5,
	}
	products, err := svc.ListProducts(context.Background(), req)
	require.Nil(t, err)
	require.NotNil(t, products)
	// require.Equal(t, 5, len(products))
}

func TestDetailProductSuccess(t *testing.T) {
	req := CreateProductRequestPayload{
		Name:  "macbookm1",
		Stock: 10,
		Price: 10000,
	}
	err := svc.CreateProduct(context.Background(), req)
	require.Nil(t, err)

	reqList := ListProductRequestPayload{
		Cursor: 0,
		Size:   5,
	}

	products, err := svc.ListProducts(context.Background(), reqList)
	require.Nil(t, err)
	require.NotNil(t, products)
	require.Greater(t, len(products), 0)

	product, err := svc.DetailProduct(context.Background(), products[0].SKU)
	require.Nil(t, err)
	require.NotNil(t, product)
	require.NotEmpty(t, product)
	require.Equal(t, products[0].ID, product.ID)
}
