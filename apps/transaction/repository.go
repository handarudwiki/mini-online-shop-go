package transaction

import (
	"context"
	"database/sql"
	"handarudwiki/mini-online-shop-go/infra/response"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func newRepository(db *sqlx.DB) repository {
	return repository{
		db: db,
	}
}

func (r repository) GetTransactionByUserPublicID(ctx context.Context, publcID string) (transactions []Transaction, err error) {
	query := `SELECT id, user_public_id, product_id, product_price, amount, subtotal, platform_fee, grand_total,
			status, product_snapshot , created_at, updated_at FROM transactions WHERE user_public_id = $1
		`
	err = r.db.SelectContext(ctx, &transactions, query, publcID)

	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}

	return
}

func (r repository) CreateTransactionWithTx(ctx context.Context, tx *sqlx.Tx, transaction Transaction) (err error) {
	query := `INSERT INTO transactions (
		user_public_id, product_id, product_price, amount, subtotal, platform_fee, grand_total,
		status, product_snapshot, created_at, updated_at
	) VALUES(
		:user_public_id, :product_id, :product_price, :amount, :subtotal, :platform_fee, :grand_total,
		:status, :product_snapshot, :created_at, :updated_at
	)
	`
	stmt, err := tx.PrepareNamedContext(ctx, query)
	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, transaction)

	return
}

func (r repository) Begin(ctx context.Context) (tx *sqlx.Tx, err error) {
	tx, err = r.db.BeginTxx(ctx, &sql.TxOptions{})
	if err != nil {
		return
	}
	return
}

func (r repository) Rollback(ctx context.Context, tx *sqlx.Tx) (err error) {
	return tx.Rollback()
}

func (r repository) Commit(ctx context.Context, tx *sqlx.Tx) (err error) {
	return tx.Commit()
}

func (r repository) GetProductBySKU(ctx context.Context, sku string) (product Product, err error) {

	query := ` SELECT id, sku, name, stock, price
			FROM products WHERE sku = $1
	`
	err = r.db.GetContext(ctx, &product, query, sku)

	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}

		return
	}
	return
}

func (r repository) UpdateProductWithTx(ctx context.Context, tx *sqlx.Tx, product Product) (err error) {
	query := `UPDATE products SET stock = :stock
			WHERE id= :id
	`
	stmt, err := r.db.PrepareNamedContext(ctx, query)

	if err != nil {
		return
	}

	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, product)

	if err != nil {
		return
	}

	return
}
