package transaction

import (
	"context"
	"handarudwiki/mini-online-shop-go/infra/response"

	"github.com/jmoiron/sqlx"
)

type Repository interface {
	TransactionDBRepository
	TransactionRepository
	ProductRepository
}

type TransactionDBRepository interface {
	Begin(ctx context.Context) (tx *sqlx.Tx, err error)
	Rollback(ctx context.Context, tx *sqlx.Tx) (err error)
	Commit(ctx context.Context, tx *sqlx.Tx) (err error)
}
type TransactionRepository interface {
	CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, transaction Transaction) (err error)
	GetTransactionByUserPublicID(ctx context.Context, publcID string) (transactions []Transaction, err error)
}
type ProductRepository interface {
	GetProductBySKU(ctx context.Context, sku string) (product Product, err error)
	UpdateProductWithTx(ctx context.Context, tx *sqlx.Tx, product Product) (err error)
}

type service struct {
	repo Repository
}

func newService(repo Repository) service {
	return service{
		repo: repo,
	}
}

func (s service) CreateTransaction(ctx context.Context, req CreateTransactionRequestPayload) (err error) {
	myProduct, err := s.repo.GetProductBySKU(ctx, req.ProductSKU)
	if err != nil {
		return
	}

	if !myProduct.IsExists() {
		err = response.ErrNotFound
		return
	}

	transaction := NewTransactionFromReq(req)
	transaction.FromProduct(myProduct)
	transaction.SetPlatformFee(1000)
	transaction.SetGrandTotal()

	err = transaction.Validate(uint8(myProduct.Stock))

	if err != nil {
		return
	}

	//start transaction db
	tx, err := s.repo.Begin(ctx)

	if err != nil {
		return
	}

	//roleback if any error
	defer s.repo.Rollback(ctx, tx)

	//create transactiom
	err = s.repo.CreateTransactionWithTx(ctx, tx, transaction)

	if err != nil {
		return
	}

	//update current stock product
	err = myProduct.UpdateStockProduct(transaction.Amount)

	if err != nil {
		return
	}

	//update into database
	err = s.repo.UpdateProductWithTx(ctx, tx, myProduct)
	if err != nil {
		return
	}

	//commit
	err = s.repo.Commit(ctx, tx)

	if err != nil {
		return
	}

	return
}

func (s service) TransactionHistories(ctx context.Context, userPublicId string) (transactions []Transaction, err error) {
	transactions, err = s.repo.GetTransactionByUserPublicID(ctx, userPublicId)
	if err != nil {
		if err == response.ErrNotFound {
			return []Transaction{}, err
		}
		return
	}

	if len(transactions) == 0 {
		return []Transaction{}, err
	}
	return
}
