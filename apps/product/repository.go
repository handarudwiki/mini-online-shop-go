package product

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

func (r repository) CreateProduct(ctx context.Context, product ProductEntity) (err error) {
	query := `INSERT INTO products 
		(
			sku,name,price,stock,created_at,updated_at
		) VALUES(
			:sku, :name, :price, :stock, :created_at, :updated_at 
		)
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

func (r repository) GetAllProducts(ctx context.Context, product ProductPagination) (products []ProductEntity, err error) {
	qeury := `SELECT id, sku, name, price, stock, created_at, updated_at 
		FROM products WHERE id > $1
		ORDER BY id ASC LIMIT $2
	`
	err = r.db.SelectContext(ctx, &products, qeury, product.Cursor, product.Size)
	if err != nil {
		if err == sql.ErrNoRows {
			err = response.ErrNotFound
			return
		}
		return
	}
	return
}

func (r repository) GetProductBySKU(ctx context.Context, sku string) (product ProductEntity, err error) {
	query := `SELECT id, sku, name, price, stock, created_at, updated_at
		FROM products WHERE sku = $1
	`
	err = r.db.GetContext(ctx, &product, sku)

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
