package transaction

import (
	"context"
	"handarudwiki/mini-online-shop-go/external/database"
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

func TestCreateTransaction(t *testing.T) {
	req := CreateTransactionRequestPayload{
		ProductSKU:   "a2e24441-b4ea-4615-be8a-b2a22d13ece5",
		Amount:       2,
		UserPublicID: "3fadb821-9531-4e85-ad40-95c0ac33c86a",
	}

	err := svc.CreateTransaction(context.Background(), req)
	require.Nil(t, err)
}

func TestGetTransactionHistory(t *testing.T) {
	req := CreateTransactionRequestPayload{
		ProductSKU:   "a2e24441-b4ea-4615-be8a-b2a22d13ece5",
		Amount:       2,
		UserPublicID: "3fadb821-9531-4e85-ad40-95c0ac33c86a",
	}

	err := svc.CreateTransaction(context.Background(), req)
	require.Nil(t, err)

	trasactions, err := svc.TransactionHistories(context.Background(), req.UserPublicID)
	require.Nil(t, err)
	require.NotEmpty(t, trasactions)
}
