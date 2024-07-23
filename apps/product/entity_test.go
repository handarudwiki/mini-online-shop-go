package product

import (
	"handarudwiki/mini-online-shop-go/infra/response"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateProduct(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		product := ProductEntity{
			Name:  "product",
			Stock: 10,
			Price: 1000,
		}
		err := product.Validate()
		require.Nil(t, err)
	})
	t.Run("Product name required", func(t *testing.T) {
		product := ProductEntity{
			Name:  "",
			Stock: 10,
			Price: 100,
		}
		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductNameRequired, err)
	})

	t.Run("Product name invalid", func(t *testing.T) {
		product := ProductEntity{
			Name:  "12",
			Stock: 10,
			Price: 100,
		}
		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrProductNameInvalid, err)
	})

	t.Run("Stock invalid", func(t *testing.T) {
		product := ProductEntity{
			Name:  "products",
			Stock: 0,
			Price: 100,
		}
		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrSrockInvalid, err)
	})
	t.Run("Price invalid", func(t *testing.T) {
		product := ProductEntity{
			Name:  "products",
			Stock: 10,
			Price: 0,
		}
		err := product.Validate()
		require.NotNil(t, err)
		require.Equal(t, response.ErrPriceInvalid, err)
	})
}
