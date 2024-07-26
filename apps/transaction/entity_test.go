package transaction

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestSubtot(t *testing.T) {
	transaction := Transaction{
		ProductPrice: 10000,
		Amount:       10,
	}
	// transaction.Subtotal = 1000
	transaction.SetSubtotal()
	require.Equal(t, uint(100_000), transaction.Subtotal)
}

func TestGrandTotal(t *testing.T) {
	t.Run("without set subtotal first", func(t *testing.T) {
		transaction := Transaction{
			ProductPrice: 10_000,
			Amount:       10,
		}

		transaction.SetSubtotal()
		require.Equal(t, uint(100_000), transaction.Subtotal)
	})
	t.Run("without platform fee", func(t *testing.T) {
		transaction := Transaction{
			ProductPrice: 10_000,
			Amount:       10,
		}
		transaction.SetSubtotal()
		transaction.SetGrandTotal()

		require.Equal(t, uint(100_000), transaction.GrandTotal)
	})

	t.Run("with platform fee", func(t *testing.T) {
		transaction := Transaction{
			ProductPrice: 10_000,
			Amount:       10,
		}

		transaction.SetSubtotal()
		transaction.SetPlatformFee(1000)
		transaction.SetGrandTotal()

		require.Equal(t, uint(101_000), transaction.GrandTotal)
	})
}

func TestProductJson(t *testing.T) {
	product := Product{
		ID:    1,
		SKU:   uuid.NewString(),
		Name:  "macboook m1",
		Price: 1000,
	}
	transaction := Transaction{}
	err := transaction.SetProductJson(product)
	require.Nil(t, err)
	require.NotNil(t, transaction.ProductJSON)

	productFromTrx, err := transaction.GetProduct()
	require.Nil(t, err)
	require.NotEmpty(t, productFromTrx)

	require.Equal(t, product, productFromTrx)
}

func TestTransactionStatus(t *testing.T) {
	type tabletest struct {
		title       string
		expected    string
		transaction Transaction
	}

	tableTests := []tabletest{
		{
			title:       "status created",
			transaction: Transaction{Status: TransactionStatusCreated},
			expected:    TRX_CREATED,
		},
		{
			title:       "status on progress",
			transaction: Transaction{Status: TransactionStatusProgres},
			expected:    TRX_ON_PROGRESS,
		},
		{
			title:       "status in delivery",
			transaction: Transaction{Status: TransactionStatusInDelivery},
			expected:    TRX_IN_DELIVERY,
		},
		{
			title:       "status completed",
			transaction: Transaction{Status: TransactionStatusCompleted},
			expected:    TRX_COMPLETED,
		},
		{
			title:       "status unknown",
			transaction: Transaction{Status: 0},
			expected:    TRX_UNKNOWN,
		},
	}

	for _, test := range tableTests {
		t.Run(test.title, func(t *testing.T) {
			require.Equal(t, test.expected, test.transaction.GetStatus())
		})
	}
}

// func TestGrandTotal(t *testing.T) {
// 	t.Run("without set subtotal first", func(t *testing.T) {
// 		transaction := Transaction{
// 			ProductPrice: 10_000,
// 			Amount:       10,
// 		}

// 		transaction.SetSubtotal()
// 		require.Equal(t, uint(100_000), transaction.GrandTotal)
// 	})
// 	t.Run("without platform fee", func(t *testing.T) {
// 		transaction := Transaction{
// 			ProductPrice: 10_000,
// 			Amount:       10,
// 		}
// 		transaction.SetSubtotal()
// 		transaction.SetGrandTotal()

// 		require.Equal(t, uint(100_000), transaction.GrandTotal)
// 	})

// 	t.Run("with platform fee", func(t *testing.T) {
// 		transaction := Transaction{
// 			ProductPrice: 10_000,
// 			Amount:       10,
// 		}

// 		transaction.SetSubtotal()
// 		transaction.SetPlatformFee(1000)
// 		transaction.SetGrandTotal()

// 		require.Equal(t, uint(100_100), transaction.GrandTotal)
// 	})
// }
